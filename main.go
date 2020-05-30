package main

import (
	"context"
	"encoding/xml"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/PuerkitoBio/goquery"
	"github.com/genuinetools/pkg/cli"
	"github.com/jessfraz/goodreads-dewey/version"
	"github.com/sirupsen/logrus"
)

var (
	goodreadsID string

	debug bool

	nsfaRegexp           = regexp.MustCompile(`\d+(?:\.\d*)?`)
	titleRegexp          = regexp.MustCompile(`(?P<title>[^:]+)(?: : )(?P<subtitle>.*)?`)
	titleTemplate        = `${title}_ ${subtitle}`
	authorRegexp         = regexp.MustCompile(`(?P<surname>[^,|]+), (?P<prename>[^(,\[|]+ ?[^ (,\[|]+?)`)
	authorTemplate       = `${prename} ${surname}`
	authorSeparator      = ", "
	titleAuthorSeparator = " - "
)

func main() {
	// Create a new cli program.
	p := cli.NewProgram()
	p.Name = "goodreads-dewey"
	p.Description = "A tool to get the dewey decimal number for books in your Goodreads account."
	// Set the GitCommit and Version.
	p.GitCommit = version.GITCOMMIT
	p.Version = version.VERSION

	// Setup the global flags.
	p.FlagSet = flag.NewFlagSet("goodreads-dewey", flag.ExitOnError)
	p.FlagSet.BoolVar(&debug, "d", false, "enable debug logging")
	p.FlagSet.BoolVar(&debug, "debug", false, "enable debug logging")

	p.FlagSet.StringVar(&goodreadsID, "id", os.Getenv("GOODREADS_ID"), "Goodreads user ID (or env var GOODREADS_ID)")

	// Set the before function.
	p.Before = func(ctx context.Context) error {
		// Set the log level.
		if debug {
			logrus.SetLevel(logrus.DebugLevel)
		}

		if len(goodreadsID) < 1 {
			return errors.New("the Goodreads user ID cannot be empty")
		}

		return nil
	}

	p.Action = func(ctx context.Context, args []string) error {
		fmt.Println("Paginating through your books, this may take a bit...")

		// Get the books.
		books, err := getBooks()
		if err != nil {
			return err
		}

		// Create the tab writer.
		w := tabwriter.NewWriter(os.Stdout, 20, 1, 3, ' ', 0)
		io.WriteString(w, "NUM\tTITLE\tAUTHOR\tISBN\tDEWEY\n")

		for i, book := range books {
			fmt.Fprintf(w, "%d\t%s\t%s %s\t%s\t%s\n",
				i,
				book.Title,
				book.AuthorFirstName,
				book.AuthorLastName,
				book.ISBN,
				book.DeweyDecimal)
		}

		w.Flush()

		return nil
	}

	// Run our program.
	p.Run()
}

type book struct {
	Title           string
	AuthorFirstName string
	AuthorLastName  string
	ISBN            string
	DeweyDecimal    string
}

func getBooks() ([]book, error) {
	books := []book{}
	page := 1

	ddc, err := parseDDC()
	if err != nil {
		return nil, fmt.Errorf("parsing ddc failed: %v", err)
	}

	// Iterate over the pages
	for page > 0 {
		url := fmt.Sprintf("https://www.goodreads.com/review/list/%s?page=%d&per_page=20", goodreadsID, page)
		resp, err := http.Get(url)
		if err != nil {
			return nil, fmt.Errorf("getting books from %s failed: %v", url, err)
		}
		defer resp.Body.Close()

		// Read the body.
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("creating document failed: %v", err)
		}

		nocontent := strings.TrimSpace(doc.Find(".greyText.nocontent.stacked").Text())
		if nocontent == "No matching items!" {
			// Reset the page to break the loop.
			page = -1
		}

		doc.Find("#booksBody tr").Each(func(j int, l *goquery.Selection) {
			title := strings.ReplaceAll(strings.TrimSuffix(strings.TrimSpace(l.Find(".title .value").Text()), "\n        *"), "\n        ", " ")

			author := strings.TrimSuffix(strings.TrimSpace(l.Find(".author .value").Text()), "\n        *")
			parts := strings.SplitN(author, ", ", 2)

			isbn := strings.TrimSpace(l.Find(".isbn .value").Text())
			if len(isbn) < 1 {
				// Try the ISBN 13.
				isbn = strings.TrimSpace(l.Find(".isbn13 .value").Text())

				if len(isbn) < 1 {
					// Try the asin.
					isbn = strings.TrimSpace(l.Find(".asin .value").Text())
				}
			}

			c, err := classifyBook("isbn", isbn, title, fmt.Sprintf("%s %s", parts[1], parts[0]))
			if err != nil {
				logrus.Warn(err)
			}
			dd := classifyString(c, ddc, 3)

			b := book{
				Title:           title,
				AuthorFirstName: parts[1],
				AuthorLastName:  parts[0],
				ISBN:            isbn,
				DeweyDecimal:    dd,
			}

			books = append(books, b)
		})

		// Iterate the page.
		page++
	}

	return books, nil
}

// DDC is a data type for the ddc.
type DDC struct {
	XMLName xml.Name `xml:"ddc"`
	Classes []*Class `xml:"class"`
}

// Class is the data type for a class.
type Class struct {
	XMLName     xml.Name    `xml:"class"`
	Number      string      `xml:"number,attr"`
	Description string      `xml:"description,attr"`
	Divisions   []*Division `xml:"division"`
}

// Division is the data type for a division.
type Division struct {
	XMLName     xml.Name   `xml:"division"`
	Number      string     `xml:"number,attr"`
	Description string     `xml:"description,attr"`
	Sections    []*Section `xml:"section"`
}

// Section is the data type for a section.
type Section struct {
	XMLName     xml.Name `xml:"section"`
	Number      string   `xml:"number,attr"`
	Description string   `xml:"description,attr"`
}

func parseDDC() (ddc DDC, err error) {
	err = xml.Unmarshal([]byte(ddcContents), &ddc)
	return ddc, err
}

// Classify is the data type for classifying a book.
type Classify struct {
	XMLName      xml.Name      `xml:"classify"`
	Response     Response      `xml:"response"`
	Input        Input         `xml:"input"`
	Works        []Work        `xml:"works>work"`
	Work         Work          `xml:"work"`
	Authors      []Author      `xml:"authors>author"`
	MostPopulars []MostPopular `xml:"recommendations>ddc>mostPopular"`
	MostPopular  MostPopular
}

// Response is the data type for a response.
type Response struct {
	Code string `xml:"code,attr"`
}

// Input is the data type for input.
type Input struct {
	Type  string `xml:"type,attr"`
	Value string `xml:",chardata"`
}

// Work is the data type for work.
type Work struct {
	Author  string `xml:"author,attr"`
	Owi     string `xml:"owi,attr"`
	Schemes string `xml:"schemes,attr"`
	Title   string `xml:"title,attr"`
	Wi      string `xml:"wi,attr"`
}

// Author is the data type for an author.
type Author struct {
	Value string `xml:",chardata"`
}

// MostPopular is the data type for most popular.
type MostPopular struct {
	Nsfa string `xml:"nsfa,attr"`
	Sfa  string `xml:"sfa,attr"`
}

func classifyBook(key, value, title, author string) (Classify, error) {
	url := fmt.Sprintf("http://classify.oclc.org/classify2/Classify?title=%s&author=%s&summary=true", title, author)
	if len(key) > 0 && len(value) > 0 {
		url = fmt.Sprintf("http://classify.oclc.org/classify2/Classify?%s=%s&summary=true", key, value)
	}

	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return Classify{}, fmt.Errorf("requesting %q=%q title=%q author=%q from the classify API failed: %v", key, value, title, author, err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Classify{}, fmt.Errorf("reading classify response body failed: %v", err)
	}

	c, err := parseClassify(key, value, title, author, body)
	if err != nil {
		return Classify{}, fmt.Errorf("classifying %q=%q title=%q author=%q failed: %v", key, value, title, author, err)
	}
	return c, nil
}

func parseClassify(key, value, title, author string, body []byte) (Classify, error) {
	var (
		classify Classify
		err      error
	)
	if err = xml.Unmarshal(body, &classify); err != nil {
		classify.Response.Code = "999"
	}
	switch classify.Response.Code {
	case "0": // Single-work summary
		fallthrough
	case "2": // Single-work detail
		for _, mostPopular := range classify.MostPopulars {
			if nsfaRegexp.MatchString(mostPopular.Nsfa) {
				classify.MostPopular = mostPopular
			}
		}
		return classify, err
	case "4": // Multi-work
		for _, work := range classify.Works {
			if strings.Contains(work.Schemes, "DDC") && classify.Input.Type != "wi" {
				c, err := classifyBook("wi", work.Wi, title, author)
				if err != nil {
					return c, err
				}
				if c.MostPopular.Nsfa != "" {
					return c, nil
				}
			}
		}
	case "100":
		classify, err = classifyBook("", "", title, author)
		if err != nil || (len(key) < 1 && len(value) < 1) {
			return classify, fmt.Errorf("error in 100 loop (no content): %v", err)
		}
	case "101":
		classify, err = classifyBook("", "", title, author)
		if err != nil || (len(key) < 1 && len(value) < 1) {
			return classify, fmt.Errorf("error in 100 loop (invalid input): %v", err)
		}
	case "102":
		classify, err = classifyBook("", "", title, author)
		if err != nil || (len(key) < 1 && len(value) < 1) {
			return classify, fmt.Errorf("error in 102 loop (not found): %v", err)
		}
	case "200":
		classify, err = classifyBook("", "", title, author)
		if err != nil || (len(key) < 1 && len(value) < 1) {
			return classify, fmt.Errorf("error in 200 loop (unexpected error): %v", err)
		}
	default:
		c, e := classifyBook("", "", title, author)
		if e != nil || (len(key) < 1 && len(value) < 1) {
			return c, fmt.Errorf("error in unknown loop (unknown error): %v %v", err, e)
		}
	}
	return classify, nil
}

func classifyString(c Classify, ddc DDC, depth int) string {
	return fmt.Sprintf("%s : %s/%s%s%s", c.MostPopular.Nsfa, classify(c, ddc, depth),
		titleString(c.Work.Title), titleAuthorSeparator, authorString(c))
}

func classify(classify Classify, ddc DDC, depth int) string {
	if classify.MostPopular.Nsfa == "" {
	} else if i, err := strconv.Atoi(classify.MostPopular.Nsfa[:3]); err != nil {
		log.Println(err)
	} else if class := ddc.Classes[i/100]; depth == 1 {
		return filepath.Join(class.Description)
	} else if division := class.Divisions[(i/10)%10]; depth == 2 {
		return filepath.Join(class.Description, division.Description)
	} else {
		section := division.Sections[i%10]
		return filepath.Join(class.Description, division.Description, section.Description)
	}
	return "Unassigned"
}

func titleString(title string) string {
	return strings.Title(capture(title, titleRegexp, titleTemplate))
}

func capture(s string, re *regexp.Regexp, t string) string {
	if re.MatchString(s) {
		b := []byte{}
		for _, m := range re.FindAllStringSubmatchIndex(s, -1) {
			b = re.ExpandString(b, t, s, m)
		}
		return string(b)
	}
	return s
}

func authorString(c Classify) string {
	n := len(c.Authors)
	if n == 0 {
		return capture(c.Work.Author, authorRegexp, authorTemplate)
	}

	b := strings.Builder{}
	for i, author := range c.Authors {
		b.WriteString(capture(author.Value, authorRegexp, authorTemplate))
		if i < n-1 {
			b.WriteString(authorSeparator)
		}
	}
	return b.String()
}
