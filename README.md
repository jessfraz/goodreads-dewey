# goodreads-dewey

[![make all](https://github.com/jessfraz/goodreads-dewey/workflows/make%20all/badge.svg)](https://github.com/jessfraz/goodreads-dewey/actions?query=workflow%3A%22make+all%22)
[![make image](https://github.com/oxidecomputer/jessfraz/goodreads-dewey/make%20image/badge.svg)](https://github.com/jessfraz/goodreads-dewey/actions?query=workflow%3A%22make+image%22)
[![GoDoc](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=for-the-badge)](https://godoc.org/github.com/jessfraz/goodreads-dewey)
[![Github All Releases](https://img.shields.io/github/downloads/jessfraz/goodreads-dewey/total.svg?style=for-the-badge)](https://github.com/jessfraz/goodreads-dewey/releases)

A tool to get the dewey decimal number for books in your Goodreads account.

This tool scrapes your public Goodreads profile for your books versus using the
Goodreads API because their API is terrible.

**Table of Contents**

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [Installation](#installation)
    - [Binaries](#binaries)
    - [Via Go](#via-go)
- [Usage](#usage)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Installation

#### Binaries

For installation instructions from binaries please visit the [Releases Page](https://github.com/jessfraz/goodreads-dewey/releases).

#### Via Go

```console
$ go get github.com/jessfraz/goodreads-dewey
```

## Usage

```console
$ goodreads-dewey -h
goodreads-dewey -  A tool to get the dewey decimal number for books in your Goodreads account.

Usage: goodreads-dewey <command>

Flags:

  -d, --debug  enable debug logging (default: false)
  --id         Goodreads user ID (or env var GOODREADS_ID)

Commands:

  version  Show the version information.
```
