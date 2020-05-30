package main

const ddcContents string = `<ddc>
  <class number="000" description="Computer science, information &amp; general works">
    <division number="000" description="Computer science, knowledge &amp; systems">
      <section number="000" description="Computer science, information &amp; general works"></section>
      <section number="001" description="Knowledge"></section>
      <section number="002" description="The book"></section>
      <section number="003" description="Systems"></section>
      <section number="004" description="Data processing &amp; computer science"></section>
      <section number="005" description="Computer programming, programs &amp; data"></section>
      <section number="006" description="Special computer methods"></section>
      <section number="007" description="Unassigned"></section>
      <section number="008" description="Unassigned"></section>
      <section number="009" description="Unassigned"></section>
    </division>
    <division number="010" description="Bibliographies">
      <section number="010" description="Bibliography"></section>
      <section number="011" description="Bibliographies"></section>
      <section number="012" description="Bibliographies of individuals"></section>
      <section number="013" description="Unassigned"></section>
      <section number="014" description="Of anonymous &amp; pseudonymous works"></section>
      <section number="015" description="Bibliographies of works from specific places"></section>
      <section number="016" description="Bibliographies of works on specific subjects"></section>
      <section number="017" description="General subject catalogs"></section>
      <section number="018" description="Catalogs arranged by author, date, etc."></section>
      <section number="019" description="Dictionary catalogs"></section>
    </division>
    <division number="020" description="Library &amp; information sciences">
      <section number="020" description="Library &amp; information sciences"></section>
      <section number="021" description="Library relationships"></section>
      <section number="022" description="Administration of physical plant"></section>
      <section number="023" description="Personnel management"></section>
      <section number="024" description="Unassigned"></section>
      <section number="025" description="Library operations"></section>
      <section number="026" description="Libraries for specific subjects"></section>
      <section number="027" description="General libraries"></section>
      <section number="028" description="Reading &amp; use of other information media"></section>
      <section number="029" description="Unassigned"></section>
    </division>
    <division number="030" description="Encyclopedias &amp; books of facts">
      <section number="030" description="General encyclopedic works"></section>
      <section number="031" description="Encyclopedias in American English"></section>
      <section number="032" description="Encyclopedias in English"></section>
      <section number="033" description="In other Germanic languages"></section>
      <section number="034" description="Encyclopedias in French, Occitan &amp; Catalan"></section>
      <section number="035" description="In Italian, Romanian &amp; related languages"></section>
      <section number="036" description="Encyclopedias in Spanish &amp; Portuguese"></section>
      <section number="037" description="Encyclopedias in Slavic languages"></section>
      <section number="038" description="Encyclopedias in Scandinavian languages"></section>
      <section number="039" description="Encyclopedias in other languages"></section>
    </division>
    <division number="040" description="Unassigned">
      <section number="040" description="Unassigned"></section>
      <section number="041" description="Unassigned"></section>
      <section number="042" description="Unassigned"></section>
      <section number="043" description="Unassigned"></section>
      <section number="044" description="Unassigned"></section>
      <section number="045" description="Unassigned"></section>
      <section number="046" description="Unassigned"></section>
      <section number="047" description="Unassigned"></section>
      <section number="048" description="Unassigned"></section>
      <section number="049" description="Unassigned"></section>
    </division>
    <division number="050" description="Magazines, journals &amp; serials">
      <section number="050" description="General serial publications"></section>
      <section number="051" description="Serials in American English"></section>
      <section number="052" description="Serials in English"></section>
      <section number="053" description="Serials in other Germanic languages"></section>
      <section number="054" description="Serials in French, Occitan &amp; Catalan"></section>
      <section number="055" description="In Italian, Romanian &amp; related languages"></section>
      <section number="056" description="Serials in Spanish &amp; Portuguese"></section>
      <section number="057" description="Serials in Slavic languages"></section>
      <section number="058" description="Serials in Scandinavian languages"></section>
      <section number="059" description="Serials in other languages"></section>
    </division>
    <division number="060" description="Associations, organizations &amp; museums">
      <section number="060" description="General organizations &amp; museum science"></section>
      <section number="061" description="Organizations in North America"></section>
      <section number="062" description="Organizations in British Isles; in England"></section>
      <section number="063" description="Organizations in central Europe; in Germany"></section>
      <section number="064" description="Organizations in France &amp; Monaco"></section>
      <section number="065" description="Organizations in Italy &amp; adjacent islands"></section>
      <section number="066" description="In Iberian peninsula &amp; adjacent islands"></section>
      <section number="067" description="Organizations in eastern Europe; in Russia"></section>
      <section number="068" description="Organizations in other geographic areas"></section>
      <section number="069" description="Museum science"></section>
    </division>
    <division number="070" description="News media, journalism &amp; publishing">
      <section number="070" description="News media, journalism &amp; publishing"></section>
      <section number="071" description="Newspapers in North America"></section>
      <section number="072" description="Newspapers in British Isles; in England"></section>
      <section number="073" description="Newspapers in central Europe; in Germany"></section>
      <section number="074" description="Newspapers in France &amp; Monaco"></section>
      <section number="075" description="Newspapers in Italy &amp; adjacent islands"></section>
      <section number="076" description="In Iberian peninsula &amp; adjacent islands"></section>
      <section number="077" description="Newspapers in eastern Europe; in Russia"></section>
      <section number="078" description="Newspapers in Scandinavia"></section>
      <section number="079" description="Newspapers in other geographic areas"></section>
    </division>
    <division number="080" description="Quotations">
      <section number="080" description="General collections"></section>
      <section number="081" description="Collections in American English"></section>
      <section number="082" description="Collections in English"></section>
      <section number="083" description="Collections in other Germanic languages"></section>
      <section number="084" description="Collections in French, Occitan &amp; Catalan"></section>
      <section number="085" description="In Italian, Romanian &amp; related languages"></section>
      <section number="086" description="Collections in Spanish &amp; Portuguese"></section>
      <section number="087" description="Collections in Slavic languages"></section>
      <section number="088" description="Collections in Scandinavian languages"></section>
      <section number="089" description="Collections in other languages"></section>
    </division>
    <division number="090" description="Manuscripts &amp; rare books">
      <section number="090" description="Manuscripts &amp; rare books"></section>
      <section number="091" description="Manuscripts"></section>
      <section number="092" description="Block books"></section>
      <section number="093" description="Incunabula"></section>
      <section number="094" description="Printed books"></section>
      <section number="095" description="Books notable for bindings"></section>
      <section number="096" description="Books notable for illustrations"></section>
      <section number="097" description="Books notable for ownership or origin"></section>
      <section number="098" description="Prohibited works, forgeries &amp; hoaxes"></section>
      <section number="099" description="Books notable for format"></section>
    </division>
  </class>
  <class number="100" description="Philosophy &amp; psychology">
    <division number="100" description="Philosophy">
      <section number="100" description="Philosophy &amp; psychology"></section>
      <section number="101" description="Theory of philosophy"></section>
      <section number="102" description="Miscellany"></section>
      <section number="103" description="Dictionaries &amp; encyclopedias"></section>
      <section number="104" description="Unassigned"></section>
      <section number="105" description="Serial publications"></section>
      <section number="106" description="Organizations &amp; management"></section>
      <section number="107" description="Education, research &amp; related topics"></section>
      <section number="108" description="Kinds of persons treatment"></section>
      <section number="109" description="Historical &amp; collected persons treatment"></section>
    </division>
    <division number="110" description="Metaphysics">
      <section number="110" description="Metaphysics"></section>
      <section number="111" description="Ontology"></section>
      <section number="112" description="Unassigned"></section>
      <section number="113" description="Cosmology"></section>
      <section number="114" description="Space"></section>
      <section number="115" description="Time"></section>
      <section number="116" description="Change"></section>
      <section number="117" description="Structure"></section>
      <section number="118" description="Force &amp; energy"></section>
      <section number="119" description="Number &amp; quantity"></section>
    </division>
    <division number="120" description="Epistemology">
      <section number="120" description="Epistemology, causation &amp; humankind"></section>
      <section number="121" description="Epistemology"></section>
      <section number="122" description="Causation"></section>
      <section number="123" description="Determinism &amp; indeterminism"></section>
      <section number="124" description="Teleology"></section>
      <section number="125" description="Unassigned"></section>
      <section number="126" description="The self"></section>
      <section number="127" description="The unconscious &amp; the subconscious"></section>
      <section number="128" description="Humankind"></section>
      <section number="129" description="Origin &amp; destiny of individual souls"></section>
    </division>
    <division number="130" description="Parapsychology &amp; occultism">
      <section number="130" description="Parapsychology &amp; occultism"></section>
      <section number="131" description="Parapsychological &amp; occult methods"></section>
      <section number="132" description="Unassigned"></section>
      <section number="133" description="Specific topics in parapsychology &amp; occultism"></section>
      <section number="134" description="Unassigned"></section>
      <section number="135" description="Dreams &amp; mysteries"></section>
      <section number="136" description="Unassigned"></section>
      <section number="137" description="Divinatory graphology"></section>
      <section number="138" description="Physiognomy"></section>
      <section number="139" description="Phrenology"></section>
    </division>
    <division number="140" description="Philosophical schools of thought">
      <section number="140" description="Specific philosophical schools"></section>
      <section number="141" description="Idealism &amp; related systems"></section>
      <section number="142" description="Critical philosophy"></section>
      <section number="143" description="Bergsonism &amp; intuitionism"></section>
      <section number="144" description="Humanism &amp; related systems"></section>
      <section number="145" description="Sensationalism"></section>
      <section number="146" description="Naturalism &amp; related systems"></section>
      <section number="147" description="Pantheism &amp; related systems"></section>
      <section number="148" description="Eclecticism, liberalism &amp; traditionalism"></section>
      <section number="149" description="Other philosophical systems"></section>
    </division>
    <division number="150" description="Psychology">
      <section number="150" description="Psychology"></section>
      <section number="151" description="Unassigned"></section>
      <section number="152" description="Perception, movement, emotions &amp; drives"></section>
      <section number="153" description="Mental processes &amp; intelligence"></section>
      <section number="154" description="Subconscious &amp; altered states"></section>
      <section number="155" description="Differential &amp; developmental psychology"></section>
      <section number="156" description="Comparative psychology"></section>
      <section number="157" description="Unassigned"></section>
      <section number="158" description="Applied psychology"></section>
      <section number="159" description="Unassigned"></section>
    </division>
    <division number="160" description="Logic">
      <section number="160" description="Logic"></section>
      <section number="161" description="Induction"></section>
      <section number="162" description="Deduction"></section>
      <section number="163" description="Unassigned"></section>
      <section number="164" description="Unassigned"></section>
      <section number="165" description="Fallacies &amp; sources of error"></section>
      <section number="166" description="Syllogisms"></section>
      <section number="167" description="Hypotheses"></section>
      <section number="168" description="Argument &amp; persuasion"></section>
      <section number="169" description="Analogy"></section>
    </division>
    <division number="170" description="Ethics">
      <section number="170" description="Ethics"></section>
      <section number="171" description="Ethical systems"></section>
      <section number="172" description="Political ethics"></section>
      <section number="173" description="Ethics of family relationships"></section>
      <section number="174" description="Occupational ethics"></section>
      <section number="175" description="Ethics of recreation &amp; leisure"></section>
      <section number="176" description="Ethics of sex &amp; reproduction"></section>
      <section number="177" description="Ethics of social relations"></section>
      <section number="178" description="Ethics of consumption"></section>
      <section number="179" description="Other ethical norms"></section>
    </division>
    <division number="180" description="Ancient, medieval &amp; eastern philosophy">
      <section number="180" description="Ancient, medieval &amp; eastern philosophy"></section>
      <section number="181" description="Eastern philosophy"></section>
      <section number="182" description="Pre-Socratic Greek philosophies"></section>
      <section number="183" description="Socratic &amp; related philosophies"></section>
      <section number="184" description="Platonic philosophy"></section>
      <section number="185" description="Aristotelian philosophy"></section>
      <section number="186" description="Skeptic &amp; Neoplatonic philosophies"></section>
      <section number="187" description="Epicurean philosophy"></section>
      <section number="188" description="Stoic philosophy"></section>
      <section number="189" description="Medieval western philosophy"></section>
    </division>
    <division number="190" description="Modern western philosophy">
      <section number="190" description="Modern western philosophy"></section>
      <section number="191" description="Philosophy of United States &amp; Canada"></section>
      <section number="192" description="Philosophy of British Isles"></section>
      <section number="193" description="Philosophy of Germany &amp; Austria"></section>
      <section number="194" description="Philosophy of France"></section>
      <section number="195" description="Philosophy of Italy"></section>
      <section number="196" description="Philosophy of Spain &amp; Portugal"></section>
      <section number="197" description="Philosophy of former Soviet Union"></section>
      <section number="198" description="Philosophy of Scandinavia"></section>
      <section number="199" description="Philosophy in other geographic areas"></section>
    </division>
  </class>
  <class number="200" description="Religion">
    <division number="200" description="Religion">
      <section number="200" description="Religion"></section>
      <section number="201" description="Religious mythology &amp; social theology"></section>
      <section number="202" description="Doctrines"></section>
      <section number="203" description="Public worship &amp; other practices"></section>
      <section number="204" description="Religious experience, life &amp; practice"></section>
      <section number="205" description="Religious ethics"></section>
      <section number="206" description="Leaders &amp; organization"></section>
      <section number="207" description="Missions &amp; religious education"></section>
      <section number="208" description="Sources"></section>
      <section number="209" description="Sects &amp; reform movements"></section>
    </division>
    <division number="210" description="Philosophy &amp; theory of religion">
      <section number="210" description="Philosophy &amp; theory of religion"></section>
      <section number="211" description="Concepts of God"></section>
      <section number="212" description="Existence, knowability &amp; attributes of God"></section>
      <section number="213" description="Creation"></section>
      <section number="214" description="Theodicy"></section>
      <section number="215" description="Science &amp; religion"></section>
      <section number="216" description="Unassigned"></section>
      <section number="217" description="Unassigned"></section>
      <section number="218" description="Humankind"></section>
      <section number="219" description="Unassigned"></section>
    </division>
    <division number="220" description="The Bible">
      <section number="220" description="Bible"></section>
      <section number="221" description="Old Testament (Tanakh)"></section>
      <section number="222" description="Historical books of Old Testament"></section>
      <section number="223" description="Poetic books of Old Testament"></section>
      <section number="224" description="Prophetic books of Old Testament"></section>
      <section number="225" description="New Testament"></section>
      <section number="226" description="Gospels &amp; Acts"></section>
      <section number="227" description="Epistles"></section>
      <section number="228" description="Revelation (Apocalypse)"></section>
      <section number="229" description="Apocrypha &amp; pseudepigrapha"></section>
    </division>
    <division number="230" description="Christianity &amp; Christian theology">
      <section number="230" description="Christianity &amp; Christian theology"></section>
      <section number="231" description="God"></section>
      <section number="232" description="Jesus Christ &amp; his family"></section>
      <section number="233" description="Humankind"></section>
      <section number="234" description="Salvation &amp; grace"></section>
      <section number="235" description="Spiritual beings"></section>
      <section number="236" description="Eschatology"></section>
      <section number="237" description="Unassigned"></section>
      <section number="238" description="Creeds &amp; catechisms"></section>
      <section number="239" description="Apologetics &amp; polemics"></section>
    </division>
    <division number="240" description="Christian practice &amp; observance">
      <section number="240" description="Christian moral &amp; devotional theology"></section>
      <section number="241" description="Christian ethics"></section>
      <section number="242" description="Devotional literature"></section>
      <section number="243" description="Evangelistic writings for individuals"></section>
      <section number="244" description="Unassigned"></section>
      <section number="245" description="Unassigned"></section>
      <section number="246" description="Use of art in Christianity"></section>
      <section number="247" description="Church furnishings &amp; articles"></section>
      <section number="248" description="Christian experience, practice &amp; life"></section>
      <section number="249" description="Christian observances in family life"></section>
    </division>
    <division number="250" description="Christian pastoral practice &amp; religious orders">
      <section number="250" description="Christian orders &amp; local church"></section>
      <section number="251" description="Preaching"></section>
      <section number="252" description="Texts of sermons"></section>
      <section number="253" description="Pastoral office &amp; work"></section>
      <section number="254" description="Parish administration"></section>
      <section number="255" description="Religious congregations &amp; orders"></section>
      <section number="256" description="Unassigned"></section>
      <section number="257" description="Unassigned"></section>
      <section number="258" description="Unassigned"></section>
      <section number="259" description="Pastoral care of families &amp; kinds of persons"></section>
    </division>
    <division number="260" description="Christian organization, social work &amp; worship">
      <section number="260" description="Social &amp; ecclesiastical theology"></section>
      <section number="261" description="Social theology"></section>
      <section number="262" description="Ecclesiology"></section>
      <section number="263" description="Days, times &amp; places of observance"></section>
      <section number="264" description="Public worship"></section>
      <section number="265" description="Sacraments, other rites &amp; acts"></section>
      <section number="266" description="Missions"></section>
      <section number="267" description="Associations for religious work"></section>
      <section number="268" description="Religious education"></section>
      <section number="269" description="Spiritual renewal"></section>
    </division>
    <division number="270" description="History of Christianity">
      <section number="270" description="History of Christianity &amp; Christian church"></section>
      <section number="271" description="Religious orders in church history"></section>
      <section number="272" description="Persecutions in church history"></section>
      <section number="273" description="Doctrinal controversies &amp; heresies"></section>
      <section number="274" description="History of Christianity in Europe"></section>
      <section number="275" description="History of Christianity in Asia"></section>
      <section number="276" description="History of Christianity in Africa"></section>
      <section number="277" description="History of Christianity in North America"></section>
      <section number="278" description="History of Christianity in South America"></section>
      <section number="279" description="History of Christianity in other areas"></section>
    </division>
    <division number="280" description="Christian denominations">
      <section number="280" description="Christian denominations &amp; sects"></section>
      <section number="281" description="Early church &amp; Eastern churches"></section>
      <section number="282" description="Roman Catholic Church"></section>
      <section number="283" description="Anglican churches"></section>
      <section number="284" description="Protestants of Continental origin"></section>
      <section number="285" description="Presbyterian, Reformed &amp; Congregational"></section>
      <section number="286" description="Baptist, Disciples of Christ &amp; Adventist"></section>
      <section number="287" description="Methodist &amp; related churches"></section>
      <section number="288" description="Unassigned"></section>
      <section number="289" description="Other denominations &amp; sects"></section>
    </division>
    <division number="290" description="Other religions">
      <section number="290" description="Other religions"></section>
      <section number="291" description="Unassigned"></section>
      <section number="292" description="Greek &amp; Roman religion"></section>
      <section number="293" description="Germanic religion"></section>
      <section number="294" description="Religions of Indic origin"></section>
      <section number="295" description="Zoroastrianism"></section>
      <section number="296" description="Judaism"></section>
      <section number="297" description="Islam, Bábism &amp; Bahá&#39;í Faith"></section>
      <section number="298" description="Optional"></section>
      <section number="299" description="Religions not provided for elsewhere"></section>
    </division>
  </class>
  <class number="300" description="Social sciences">
    <division number="300" description="Social sciences, sociology &amp; anthropology">
      <section number="300" description="Social sciences"></section>
      <section number="301" description="Sociology &amp; anthropology"></section>
      <section number="302" description="Social interaction"></section>
      <section number="303" description="Social processes"></section>
      <section number="304" description="Factors affecting social behavior"></section>
      <section number="305" description="Social groups"></section>
      <section number="306" description="Culture &amp; institutions"></section>
      <section number="307" description="Communities"></section>
      <section number="308" description="Unassigned"></section>
      <section number="309" description="Unassigned"></section>
    </division>
    <division number="310" description="Statistics">
      <section number="310" description="Collections of general statistics"></section>
      <section number="311" description="Unassigned"></section>
      <section number="312" description="Unassigned"></section>
      <section number="313" description="Unassigned"></section>
      <section number="314" description="General statistics of Europe"></section>
      <section number="315" description="General statistics of Asia"></section>
      <section number="316" description="General statistics of Africa"></section>
      <section number="317" description="General statistics of North America"></section>
      <section number="318" description="General statistics of South America"></section>
      <section number="319" description="General statistics of other areas"></section>
    </division>
    <division number="320" description="Political science">
      <section number="320" description="Political science"></section>
      <section number="321" description="Systems of governments &amp; states"></section>
      <section number="322" description="Relation of state to organized groups"></section>
      <section number="323" description="Civil &amp; political rights"></section>
      <section number="324" description="The political process"></section>
      <section number="325" description="International migration &amp; colonization"></section>
      <section number="326" description="Slavery &amp; emancipation"></section>
      <section number="327" description="International relations"></section>
      <section number="328" description="The legislative process"></section>
      <section number="329" description="Unassigned"></section>
    </division>
    <division number="330" description="Economics">
      <section number="330" description="Economics"></section>
      <section number="331" description="Labor economics"></section>
      <section number="332" description="Financial economics"></section>
      <section number="333" description="Economics of land &amp; energy"></section>
      <section number="334" description="Cooperatives"></section>
      <section number="335" description="Socialism &amp; related systems"></section>
      <section number="336" description="Public finance"></section>
      <section number="337" description="International economics"></section>
      <section number="338" description="Production"></section>
      <section number="339" description="Macroeconomics &amp; related topics"></section>
    </division>
    <division number="340" description="Law">
      <section number="340" description="Law"></section>
      <section number="341" description="Law of nations"></section>
      <section number="342" description="Constitutional &amp; administrative law"></section>
      <section number="343" description="Military, tax, trade &amp; industrial law"></section>
      <section number="344" description="Labor, social, education &amp; cultural law"></section>
      <section number="345" description="Criminal law"></section>
      <section number="346" description="Private law"></section>
      <section number="347" description="Civil procedure &amp; courts"></section>
      <section number="348" description="Laws, regulations &amp; cases"></section>
      <section number="349" description="Law of specific jurisdictions &amp; areas"></section>
    </division>
    <division number="350" description="Public administration &amp; military science">
      <section number="350" description="Public administration &amp; military science"></section>
      <section number="351" description="Public administration"></section>
      <section number="352" description="General considerations of public administration"></section>
      <section number="353" description="Specific fields of public administration"></section>
      <section number="354" description="Administration of economy &amp; environment"></section>
      <section number="355" description="Military science"></section>
      <section number="356" description="Infantry forces &amp; warfare"></section>
      <section number="357" description="Mounted forces &amp; warfare"></section>
      <section number="358" description="Air &amp; other specialized forces"></section>
      <section number="359" description="Sea forces &amp; warfare"></section>
    </division>
    <division number="360" description="Social problems &amp; social services">
      <section number="360" description="Social problems &amp; services; associations"></section>
      <section number="361" description="Social problems &amp; social welfare in general"></section>
      <section number="362" description="Social welfare problems &amp; services"></section>
      <section number="363" description="Other social problems &amp; services"></section>
      <section number="364" description="Criminology"></section>
      <section number="365" description="Penal &amp; related institutions"></section>
      <section number="366" description="Associations"></section>
      <section number="367" description="General clubs"></section>
      <section number="368" description="Insurance"></section>
      <section number="369" description="Miscellaneous kinds of associations"></section>
    </division>
    <division number="370" description="Education">
      <section number="370" description="Education"></section>
      <section number="371" description="Schools &amp; their activities; special education"></section>
      <section number="372" description="Elementary education"></section>
      <section number="373" description="Secondary education"></section>
      <section number="374" description="Adult education"></section>
      <section number="375" description="Curricula"></section>
      <section number="376" description="Unassigned"></section>
      <section number="377" description="Unassigned"></section>
      <section number="378" description="Higher education"></section>
      <section number="379" description="Public policy issues in education"></section>
    </division>
    <division number="380" description="Commerce, communications &amp; transportation">
      <section number="380" description="Commerce, communications &amp; transportation"></section>
      <section number="381" description="Commerce"></section>
      <section number="382" description="International commerce"></section>
      <section number="383" description="Postal communication"></section>
      <section number="384" description="Communications; telecommunication"></section>
      <section number="385" description="Railroad transportation"></section>
      <section number="386" description="Inland waterway &amp; ferry transportation"></section>
      <section number="387" description="Water, air &amp; space transportation"></section>
      <section number="388" description="Transportation; ground transportation"></section>
      <section number="389" description="Metrology &amp; standardization"></section>
    </division>
    <division number="390" description="Customs, etiquette &amp; folklore">
      <section number="390" description="Customs, etiquette &amp; folklore"></section>
      <section number="391" description="Costume &amp; personal appearance"></section>
      <section number="392" description="Customs of life cycle &amp; domestic life"></section>
      <section number="393" description="Death customs"></section>
      <section number="394" description="General customs"></section>
      <section number="395" description="Etiquette (Manners)"></section>
      <section number="396" description="Unassigned"></section>
      <section number="397" description="Unassigned"></section>
      <section number="398" description="Folklore"></section>
      <section number="399" description="Customs of war &amp; diplomacy"></section>
    </division>
  </class>
  <class number="400" description="Language">
    <division number="400" description="Language">
      <section number="400" description="Language"></section>
      <section number="401" description="Philosophy &amp; theory"></section>
      <section number="402" description="Miscellany"></section>
      <section number="403" description="Dictionaries &amp; encyclopedias"></section>
      <section number="404" description="Special topics"></section>
      <section number="405" description="Serial publications"></section>
      <section number="406" description="Organizations &amp; management"></section>
      <section number="407" description="Education, research &amp; related topics"></section>
      <section number="408" description="Kinds of persons treatment"></section>
      <section number="409" description="Geographic &amp; persons treatment"></section>
    </division>
    <division number="410" description="Linguistics">
      <section number="410" description="Linguistics"></section>
      <section number="411" description="Writing systems"></section>
      <section number="412" description="Etymology"></section>
      <section number="413" description="Dictionaries"></section>
      <section number="414" description="Phonology &amp; phonetics"></section>
      <section number="415" description="Grammar"></section>
      <section number="416" description="Unassigned"></section>
      <section number="417" description="Dialectology &amp; historical linguistics"></section>
      <section number="418" description="Standard usage &amp; applied linguistics"></section>
      <section number="419" description="Sign languages"></section>
    </division>
    <division number="420" description="English &amp; Old English languages">
      <section number="420" description="English &amp; Old English"></section>
      <section number="421" description="English writing system &amp; phonology"></section>
      <section number="422" description="English etymology"></section>
      <section number="423" description="English dictionaries"></section>
      <section number="424" description="Unassigned"></section>
      <section number="425" description="English grammar"></section>
      <section number="426" description="Unassigned"></section>
      <section number="427" description="English language variations"></section>
      <section number="428" description="Standard English usage"></section>
      <section number="429" description="Old English (Anglo-Saxon)"></section>
    </division>
    <division number="430" description="German &amp; related languages">
      <section number="430" description="Germanic languages; German"></section>
      <section number="431" description="German writing systems &amp; phonology"></section>
      <section number="432" description="German etymology"></section>
      <section number="433" description="German dictionaries"></section>
      <section number="434" description="Unassigned"></section>
      <section number="435" description="German grammar"></section>
      <section number="436" description="Unassigned"></section>
      <section number="437" description="German language variations"></section>
      <section number="438" description="Standard German usage"></section>
      <section number="439" description="Other Germanic languages"></section>
    </division>
    <division number="440" description="French &amp; related languages">
      <section number="440" description="Romance languages; French"></section>
      <section number="441" description="French writing systems &amp; phonology"></section>
      <section number="442" description="French etymology"></section>
      <section number="443" description="French dictionaries"></section>
      <section number="444" description="Unassigned"></section>
      <section number="445" description="French grammar"></section>
      <section number="446" description="Unassigned"></section>
      <section number="447" description="French language variations"></section>
      <section number="448" description="Standard French usage"></section>
      <section number="449" description="Occitan &amp; Catalan"></section>
    </division>
    <division number="450" description="Italian, Romanian &amp; related languages">
      <section number="450" description="Italian, Romanian &amp; related languages"></section>
      <section number="451" description="Italian writing systems &amp; phonology"></section>
      <section number="452" description="Italian etymology"></section>
      <section number="453" description="Italian dictionaries"></section>
      <section number="454" description="Unassigned"></section>
      <section number="455" description="Italian grammar"></section>
      <section number="456" description="Unassigned"></section>
      <section number="457" description="Italian language variations"></section>
      <section number="458" description="Standard Italian usage"></section>
      <section number="459" description="Romanian &amp; related languages"></section>
    </division>
    <division number="460" description="Spanish &amp; Portuguese languages">
      <section number="460" description="Spanish &amp; Portuguese languages"></section>
      <section number="461" description="Spanish writing systems &amp; phonology"></section>
      <section number="462" description="Spanish etymology"></section>
      <section number="463" description="Spanish dictionaries"></section>
      <section number="464" description="Unassigned"></section>
      <section number="465" description="Spanish grammar"></section>
      <section number="466" description="Unassigned"></section>
      <section number="467" description="Spanish language variations"></section>
      <section number="468" description="Standard Spanish usage"></section>
      <section number="469" description="Portuguese"></section>
    </division>
    <division number="470" description="Latin &amp; Italic languages">
      <section number="470" description="Italic languages; Latin"></section>
      <section number="471" description="Classical Latin writing &amp; phonology"></section>
      <section number="472" description="Classical Latin etymology"></section>
      <section number="473" description="Classical Latin dictionaries"></section>
      <section number="474" description="Unassigned"></section>
      <section number="475" description="Classical Latin grammar"></section>
      <section number="476" description="Unassigned"></section>
      <section number="477" description="Old, postclassical &amp; Vulgar Latin"></section>
      <section number="478" description="Classical Latin usage"></section>
      <section number="479" description="Other Italic languages"></section>
    </division>
    <division number="480" description="Classical &amp; modern Greek languages">
      <section number="480" description="Hellenic languages; classical Greek"></section>
      <section number="481" description="Classical Greek writing &amp; phonology"></section>
      <section number="482" description="Classical Greek etymology"></section>
      <section number="483" description="Classical Greek dictionaries"></section>
      <section number="484" description="Unassigned"></section>
      <section number="485" description="Classical Greek grammar"></section>
      <section number="486" description="Unassigned"></section>
      <section number="487" description="Preclassical &amp; postclassical Greek"></section>
      <section number="488" description="Classical Greek usage"></section>
      <section number="489" description="Other Hellenic languages"></section>
    </division>
    <division number="490" description="Other languages">
      <section number="490" description="Other languages"></section>
      <section number="491" description="East Indo-European &amp; Celtic languages"></section>
      <section number="492" description="Afro-Asiatic languages; Semitic languages"></section>
      <section number="493" description="Non-Semitic Afro-Asiatic languages"></section>
      <section number="494" description="Altic, Uralic, Hyperborean &amp; Dravidian"></section>
      <section number="495" description="Languages of East &amp; Southeast Asia"></section>
      <section number="496" description="African languages"></section>
      <section number="497" description="North American native languages"></section>
      <section number="498" description="South American native languages"></section>
      <section number="499" description="Austronesian &amp; other languages"></section>
    </division>
  </class>
  <class number="500" description="Science">
    <division number="500" description="Science">
      <section number="500" description="Natural sciences &amp; mathematics"></section>
      <section number="501" description="Philosophy &amp; theory"></section>
      <section number="502" description="Miscellany"></section>
      <section number="503" description="Dictionaries &amp; encyclopedias"></section>
      <section number="504" description="Unassigned"></section>
      <section number="505" description="Serial publications"></section>
      <section number="506" description="Organizations &amp; management"></section>
      <section number="507" description="Education, research &amp; related topics"></section>
      <section number="508" description="Natural history"></section>
      <section number="509" description="Historical, geographic &amp; persons treatment"></section>
    </division>
    <division number="510" description="Mathematics">
      <section number="510" description="Mathematics"></section>
      <section number="511" description="General principles of mathematics"></section>
      <section number="512" description="Algebra"></section>
      <section number="513" description="Arithmetic"></section>
      <section number="514" description="Topology"></section>
      <section number="515" description="Analysis"></section>
      <section number="516" description="Geometry"></section>
      <section number="517" description="Unassigned"></section>
      <section number="518" description="Numerical analysis"></section>
      <section number="519" description="Probabilities &amp; applied mathematics"></section>
    </division>
    <division number="520" description="Astronomy">
      <section number="520" description="Astronomy &amp; allied sciences"></section>
      <section number="521" description="Celestial mechanics"></section>
      <section number="522" description="Techniques, equipment &amp; materials"></section>
      <section number="523" description="Specific celestial bodies &amp; phenomena"></section>
      <section number="524" description="Unassigned"></section>
      <section number="525" description="Earth (Astronomical geography)"></section>
      <section number="526" description="Mathematical geography"></section>
      <section number="527" description="Celestial navigation"></section>
      <section number="528" description="Ephemerides"></section>
      <section number="529" description="Chronology"></section>
    </division>
    <division number="530" description="Physics">
      <section number="530" description="Physics"></section>
      <section number="531" description="Classical mechanics; solid mechanics"></section>
      <section number="532" description="Fluid mechanics; liquid mechanics"></section>
      <section number="533" description="Gas mechanics"></section>
      <section number="534" description="Sound &amp; related vibrations"></section>
      <section number="535" description="Light &amp; infrared &amp; ultraviolet phenomena"></section>
      <section number="536" description="Heat"></section>
      <section number="537" description="Electricity &amp; electronics"></section>
      <section number="538" description="Magnetism"></section>
      <section number="539" description="Modern physics"></section>
    </division>
    <division number="540" description="Chemistry">
      <section number="540" description="Chemistry &amp; allied sciences"></section>
      <section number="541" description="Physical chemistry"></section>
      <section number="542" description="Techniques, equipment &amp; materials"></section>
      <section number="543" description="Analytical chemistry"></section>
      <section number="544" description="Unassigned"></section>
      <section number="545" description="Unassigned"></section>
      <section number="546" description="Inorganic chemistry"></section>
      <section number="547" description="Organic chemistry"></section>
      <section number="548" description="Crystallography"></section>
      <section number="549" description="Mineralogy"></section>
    </division>
    <division number="550" description="Earth sciences &amp; geology">
      <section number="550" description="Earth sciences"></section>
      <section number="551" description="Geology, hydrology &amp; meteorology"></section>
      <section number="552" description="Petrology"></section>
      <section number="553" description="Economic geology"></section>
      <section number="554" description="Earth sciences of Europe"></section>
      <section number="555" description="Earth sciences of Asia"></section>
      <section number="556" description="Earth sciences of Africa"></section>
      <section number="557" description="Earth sciences of North America"></section>
      <section number="558" description="Earth sciences of South America"></section>
      <section number="559" description="Earth sciences of other areas"></section>
    </division>
    <division number="560" description="Fossils &amp; prehistoric life">
      <section number="560" description="Paleontology; paleozoology"></section>
      <section number="561" description="Paleobotany; fossil microorganisms"></section>
      <section number="562" description="Fossil invertebrates"></section>
      <section number="563" description="Fossil marine &amp; seashore invertebrates"></section>
      <section number="564" description="Fossil mollusks &amp; molluscoids"></section>
      <section number="565" description="Fossil arthropods"></section>
      <section number="566" description="Fossil chordates"></section>
      <section number="567" description="Fossil cold-blooded vertebrates; fossil fishes"></section>
      <section number="568" description="Fossil birds"></section>
      <section number="569" description="Fossil mammals"></section>
    </division>
    <division number="570" description="Life sciences; biology">
      <section number="570" description="Life sciences; biology"></section>
      <section number="571" description="Physiology &amp; related subjects"></section>
      <section number="572" description="Biochemistry"></section>
      <section number="573" description="Specific physiological systems in animals"></section>
      <section number="574" description="Unassigned"></section>
      <section number="575" description="Specific parts of &amp; systems in plants"></section>
      <section number="576" description="Genetics and evolution"></section>
      <section number="577" description="Ecology"></section>
      <section number="578" description="Natural history of organisms"></section>
      <section number="579" description="Microorganisms, fungi &amp; algae"></section>
    </division>
    <division number="580" description="Plants (Botany)">
      <section number="580" description="Plants (Botany)"></section>
      <section number="581" description="Specific topics in natural history"></section>
      <section number="582" description="Plants noted for characteristics &amp; flowers"></section>
      <section number="583" description="Dicotyledones"></section>
      <section number="584" description="Monocotyledones"></section>
      <section number="585" description="Gymnosperms; conifers"></section>
      <section number="586" description="Seedless plants"></section>
      <section number="587" description="Vascular seedless plants"></section>
      <section number="588" description="Bryophyta"></section>
      <section number="589" description="Unassigned"></section>
    </division>
    <division number="590" description="Animals (Zoology)">
      <section number="590" description="Animals (Zoology)"></section>
      <section number="591" description="Specific topics in natural history"></section>
      <section number="592" description="Invertebrates"></section>
      <section number="593" description="Marine &amp; seashore invertebrates"></section>
      <section number="594" description="Mollusks &amp; Molluscoids"></section>
      <section number="595" description="Arthropods"></section>
      <section number="596" description="Chordates"></section>
      <section number="597" description="Cold-blooded vertebrates; fishes"></section>
      <section number="598" description="Birds"></section>
      <section number="599" description="Mammals"></section>
    </division>
  </class>
  <class number="600" description="Technology">
    <division number="600" description="Technology">
      <section number="600" description="Technology"></section>
      <section number="601" description="Philosophy &amp; theory"></section>
      <section number="602" description="Miscellany"></section>
      <section number="603" description="Dictionaries &amp; encyclopedias"></section>
      <section number="604" description="Special Topics"></section>
      <section number="605" description="Serial publications"></section>
      <section number="606" description="Organizations"></section>
      <section number="607" description="Education, research &amp; related topics"></section>
      <section number="608" description="Inventions &amp; patents"></section>
      <section number="609" description="Historical, geographic &amp; persons treatment"></section>
    </division>
    <division number="610" description="Medicine &amp; health">
      <section number="610" description="Medicine &amp; health"></section>
      <section number="611" description="Human anatomy, cytology &amp; histology"></section>
      <section number="612" description="Human physiology"></section>
      <section number="613" description="Personal health &amp; safety"></section>
      <section number="614" description="Incidence &amp; prevention of disease"></section>
      <section number="615" description="Pharmacology &amp; therapeutics"></section>
      <section number="616" description="Diseases"></section>
      <section number="617" description="Surgery &amp; related medical specialties"></section>
      <section number="618" description="Gynecology, obstetrics, pediatrics &amp; geriatrics"></section>
      <section number="619" description="Unassigned"></section>
    </division>
    <division number="620" description="Engineering">
      <section number="620" description="Engineering &amp; allied operations"></section>
      <section number="621" description="Applied physics"></section>
      <section number="622" description="Mining &amp; related operations"></section>
      <section number="623" description="Military &amp; nautical engineering"></section>
      <section number="624" description="Civil engineering"></section>
      <section number="625" description="Engineering of railroads &amp; roads"></section>
      <section number="626" description="Unassigned"></section>
      <section number="627" description="Hydraulic engineering"></section>
      <section number="628" description="Sanitary &amp; municipal engineering"></section>
      <section number="629" description="Other branches of engineering"></section>
    </division>
    <division number="630" description="Agriculture">
      <section number="630" description="Agriculture &amp; related technologies"></section>
      <section number="631" description="Techniques, equipment &amp; materials"></section>
      <section number="632" description="Plant injuries, diseases &amp; pests"></section>
      <section number="633" description="Field &amp; plantation crops"></section>
      <section number="634" description="Orchards, fruits, forestry"></section>
      <section number="635" description="Garden crops (Horticulture)"></section>
      <section number="636" description="Animal husbandry"></section>
      <section number="637" description="Processing dairy &amp; related products"></section>
      <section number="638" description="Insect culture"></section>
      <section number="639" description="Hunting, fishing &amp; conservation"></section>
    </division>
    <division number="640" description="Home &amp; family management">
      <section number="640" description="Home &amp; family management"></section>
      <section number="641" description="Food &amp; drink"></section>
      <section number="642" description="Meals &amp; table service"></section>
      <section number="643" description="Housing &amp; household equipment"></section>
      <section number="644" description="Household utilities"></section>
      <section number="645" description="Household furnishings"></section>
      <section number="646" description="Sewing, clothing &amp; personal living"></section>
      <section number="647" description="Management of public households"></section>
      <section number="648" description="Housekeeping"></section>
      <section number="649" description="Child rearing &amp; home care of persons"></section>
    </division>
    <division number="650" description="Management &amp; public relations">
      <section number="650" description="Management &amp; auxiliary services"></section>
      <section number="651" description="Office services"></section>
      <section number="652" description="Processes of written communication"></section>
      <section number="653" description="Shorthand"></section>
      <section number="654" description="Unassigned"></section>
      <section number="655" description="Unassigned"></section>
      <section number="656" description="Unassigned"></section>
      <section number="657" description="Accounting"></section>
      <section number="658" description="General management"></section>
      <section number="659" description="Advertising &amp; public relations"></section>
    </division>
    <division number="660" description="Chemical engineering">
      <section number="660" description="Chemical engineering"></section>
      <section number="661" description="Industrial chemicals"></section>
      <section number="662" description="Explosives, fuels &amp; related products"></section>
      <section number="663" description="Beverage technology"></section>
      <section number="664" description="Food technology"></section>
      <section number="665" description="Industrial oils, fats, waxes &amp; gases"></section>
      <section number="666" description="Ceramic &amp; allied technologies"></section>
      <section number="667" description="Cleaning, color &amp; coating technologies"></section>
      <section number="668" description="Technology of other organic products"></section>
      <section number="669" description="Metallurgy"></section>
    </division>
    <division number="670" description="Manufacturing">
      <section number="670" description="Manufacturing"></section>
      <section number="671" description="Metalworking &amp; primary metal products"></section>
      <section number="672" description="Iron, steel &amp; other iron alloys"></section>
      <section number="673" description="Nonferrous metals"></section>
      <section number="674" description="Lumber processing, wood products &amp; cork"></section>
      <section number="675" description="Leather &amp; fur processing"></section>
      <section number="676" description="Pulp &amp; paper technology"></section>
      <section number="677" description="Textiles"></section>
      <section number="678" description="Elastomers &amp; elastomer products"></section>
      <section number="679" description="Other products of specific materials"></section>
    </division>
    <division number="680" description="Manufacture for specific uses">
      <section number="680" description="Manufacture for specific uses"></section>
      <section number="681" description="Precision instruments &amp; other devices"></section>
      <section number="682" description="Small forge work (Blacksmithing)"></section>
      <section number="683" description="Hardware &amp; household appliances"></section>
      <section number="684" description="Furnishings &amp; home workshops"></section>
      <section number="685" description="Leather, fur goods &amp; related products"></section>
      <section number="686" description="Printing &amp; related activities"></section>
      <section number="687" description="Clothing &amp; accessories"></section>
      <section number="688" description="Other final products &amp; packaging"></section>
      <section number="689" description="Unassigned"></section>
    </division>
    <division number="690" description="Building &amp; construction">
      <section number="690" description="Buildings"></section>
      <section number="691" description="Building materials"></section>
      <section number="692" description="Auxiliary construction practices"></section>
      <section number="693" description="Specific materials &amp; purposes"></section>
      <section number="694" description="Wood construction &amp; carpentry"></section>
      <section number="695" description="Roof covering"></section>
      <section number="696" description="Utilities"></section>
      <section number="697" description="Heating, ventilating &amp; air-conditioning"></section>
      <section number="698" description="Detail finishing"></section>
      <section number="699" description="Unassigned"></section>
    </division>
  </class>
  <class number="700" description="Arts &amp; recreation">
    <division number="700" description="Arts">
      <section number="700" description="The Arts; fine &amp; decorative arts"></section>
      <section number="701" description="Philosophy of fine &amp; decorative arts"></section>
      <section number="702" description="Miscellany of fine &amp; decorative arts"></section>
      <section number="703" description="Dictionaries of fine &amp; decorative arts"></section>
      <section number="704" description="Special topics in fine &amp; decorative arts"></section>
      <section number="705" description="Serial publications of fine &amp; decorative arts"></section>
      <section number="706" description="Organizations &amp; management"></section>
      <section number="707" description="Education, research &amp; related topics"></section>
      <section number="708" description="Galleries, museums &amp; private collections"></section>
      <section number="709" description="Historical, geographic &amp; persons treatment"></section>
    </division>
    <division number="710" description="Landscape &amp; area planning">
      <section number="710" description="Civic &amp; landscape art"></section>
      <section number="711" description="Area planning"></section>
      <section number="712" description="Landscape architecture"></section>
      <section number="713" description="Landscape architecture of trafficways"></section>
      <section number="714" description="Water features"></section>
      <section number="715" description="Woody plants"></section>
      <section number="716" description="Herbaceous plants"></section>
      <section number="717" description="Structures in landscape architecture"></section>
      <section number="718" description="Landscape design of cemeteries"></section>
      <section number="719" description="Natural landscapes"></section>
    </division>
    <division number="720" description="Architecture">
      <section number="720" description="Architecture"></section>
      <section number="721" description="Architectural structure"></section>
      <section number="722" description="Architecture to ca. 300"></section>
      <section number="723" description="Architecture from ca. 300 to 1399"></section>
      <section number="724" description="Architecture from 1400"></section>
      <section number="725" description="Public structures"></section>
      <section number="726" description="Buildings for religious purposes"></section>
      <section number="727" description="Buildings for education &amp; research"></section>
      <section number="728" description="Residential &amp; related buildings"></section>
      <section number="729" description="Design &amp; decoration"></section>
    </division>
    <division number="730" description="Sculpture, ceramics &amp; metalwork">
      <section number="730" description="Plastic arts; sculpture"></section>
      <section number="731" description="Processes, forms &amp; subjects of sculpture"></section>
      <section number="732" description="Sculpture to ca. 500"></section>
      <section number="733" description="Greek, Etruscan &amp; Roman sculpture"></section>
      <section number="734" description="Sculpture from ca. 500 to 1399"></section>
      <section number="735" description="Sculpture from 1400"></section>
      <section number="736" description="Carving &amp; carvings"></section>
      <section number="737" description="Numismatics &amp; sigillography"></section>
      <section number="738" description="Ceramic arts"></section>
      <section number="739" description="Art metalwork"></section>
    </division>
    <division number="740" description="Drawing &amp; decorative arts">
      <section number="740" description="Drawing &amp; decorative arts"></section>
      <section number="741" description="Drawing &amp; drawings"></section>
      <section number="742" description="Perspective"></section>
      <section number="743" description="Drawing &amp; drawings by subject"></section>
      <section number="744" description="Unassigned"></section>
      <section number="745" description="Decorative arts"></section>
      <section number="746" description="Textile arts"></section>
      <section number="747" description="Interior decoration"></section>
      <section number="748" description="Glass"></section>
      <section number="749" description="Furniture &amp; accessories"></section>
    </division>
    <division number="750" description="Painting">
      <section number="750" description="Painting &amp; paintings"></section>
      <section number="751" description="Techniques, equipment, materials &amp; forms"></section>
      <section number="752" description="Color"></section>
      <section number="753" description="Symbolism, allegory, mythology &amp; legend"></section>
      <section number="754" description="Genre paintings"></section>
      <section number="755" description="Religion"></section>
      <section number="756" description="Unassigned"></section>
      <section number="757" description="Human figures"></section>
      <section number="758" description="Other subjects"></section>
      <section number="759" description="Historical, geographic &amp; persons treatment"></section>
    </division>
    <division number="760" description="Graphic arts">
      <section number="760" description="Graphic arts; printmaking &amp; prints"></section>
      <section number="761" description="Relief processes (Block printing)"></section>
      <section number="762" description="Unassigned"></section>
      <section number="763" description="Lithographic processes"></section>
      <section number="764" description="Chromolithography &amp; serigraphy"></section>
      <section number="765" description="Metal engraving"></section>
      <section number="766" description="Mezzotinting, aquatinting &amp; related processes"></section>
      <section number="767" description="Etching &amp; drypoint"></section>
      <section number="768" description="Unassigned"></section>
      <section number="769" description="Prints"></section>
    </division>
    <division number="770" description="Photography &amp; computer art">
      <section number="770" description="Photography, photographs &amp; computer art"></section>
      <section number="771" description="Techniques, equipment &amp; materials"></section>
      <section number="772" description="Metallic salt processes"></section>
      <section number="773" description="Pigment processes of printing"></section>
      <section number="774" description="Holography"></section>
      <section number="775" description="Digital photography"></section>
      <section number="776" description="Computer art (Digital art)"></section>
      <section number="777" description="Unassigned"></section>
      <section number="778" description="Fields &amp; kinds of photography"></section>
      <section number="779" description="Photographs"></section>
    </division>
    <division number="780" description="Music">
      <section number="780" description="Music"></section>
      <section number="781" description="General principles &amp; musical forms"></section>
      <section number="782" description="Vocal music"></section>
      <section number="783" description="Music for single voices; the voice"></section>
      <section number="784" description="Instruments &amp; instrumental ensembles"></section>
      <section number="785" description="Ensembles with one instrument per part"></section>
      <section number="786" description="Keyboard &amp; other instruments"></section>
      <section number="787" description="Stringed instruments"></section>
      <section number="788" description="Wind instruments"></section>
      <section number="789" description="Optional"></section>
    </division>
    <division number="790" description="Sports, games &amp; entertainment">
      <section number="790" description="Recreational &amp; performing arts"></section>
      <section number="791" description="Public performances"></section>
      <section number="792" description="Stage presentations"></section>
      <section number="793" description="Indoor games &amp; amusements"></section>
      <section number="794" description="Indoor games of skill"></section>
      <section number="795" description="Games of chance"></section>
      <section number="796" description="Athletic &amp; outdoor sports &amp; games"></section>
      <section number="797" description="Aquatic &amp; air sports"></section>
      <section number="798" description="Equestrian sports &amp; animal racing"></section>
      <section number="799" description="Fishing, hunting &amp; shooting"></section>
    </division>
  </class>
  <class number="800" description="Literature">
    <division number="800" description="Literature, rhetoric &amp; criticism">
      <section number="800" description="Literature &amp; rhetoric"></section>
      <section number="801" description="Philosophy &amp; theory"></section>
      <section number="802" description="Miscellany"></section>
      <section number="803" description="Dictionaries &amp; encyclopedias"></section>
      <section number="804" description="Unassigned"></section>
      <section number="805" description="Serial publications"></section>
      <section number="806" description="Organizations &amp; management"></section>
      <section number="807" description="Education, research &amp; related topics"></section>
      <section number="808" description="Rhetoric &amp; collections of literature"></section>
      <section number="809" description="History, description &amp; criticism"></section>
    </division>
    <division number="810" description="American literature in English">
      <section number="810" description="American literature in English"></section>
      <section number="811" description="American poetry in English"></section>
      <section number="812" description="American drama in English"></section>
      <section number="813" description="American fiction in English"></section>
      <section number="814" description="American essays in English"></section>
      <section number="815" description="American speeches in English"></section>
      <section number="816" description="American letters in English"></section>
      <section number="817" description="American humor &amp; satire in English"></section>
      <section number="818" description="American miscellaneous writings"></section>
      <section number="819" description="Optional"></section>
    </division>
    <division number="820" description="English &amp; Old English literatures">
      <section number="820" description="English &amp; Old English literatures"></section>
      <section number="821" description="English poetry"></section>
      <section number="822" description="English drama"></section>
      <section number="823" description="English fiction"></section>
      <section number="824" description="English essays"></section>
      <section number="825" description="English speeches"></section>
      <section number="826" description="English letters"></section>
      <section number="827" description="English humor &amp; satire"></section>
      <section number="828" description="English miscellaneous writings"></section>
      <section number="829" description="Old English (Anglo-Saxon)"></section>
    </division>
    <division number="830" description="German &amp; related literatures">
      <section number="830" description="Literatures of Germanic languages"></section>
      <section number="831" description="German poetry"></section>
      <section number="832" description="German drama"></section>
      <section number="833" description="German fiction"></section>
      <section number="834" description="German essays"></section>
      <section number="835" description="German speeches"></section>
      <section number="836" description="German letters"></section>
      <section number="837" description="German humor &amp; satire"></section>
      <section number="838" description="German miscellaneous writings"></section>
      <section number="839" description="Other Germanic literatures"></section>
    </division>
    <division number="840" description="French &amp; related literatures">
      <section number="840" description="Literatures of Romance languages"></section>
      <section number="841" description="French poetry"></section>
      <section number="842" description="French drama"></section>
      <section number="843" description="French fiction"></section>
      <section number="844" description="French essays"></section>
      <section number="845" description="French speeches"></section>
      <section number="846" description="French letters"></section>
      <section number="847" description="French humor &amp; satire"></section>
      <section number="848" description="French miscellaneous writings"></section>
      <section number="849" description="Occitan &amp; Catalan literatures"></section>
    </division>
    <division number="850" description="Italian, Romanian &amp; related literatures">
      <section number="850" description="Italian, Romanian &amp; related literatures"></section>
      <section number="851" description="Italian poetry"></section>
      <section number="852" description="Italian drama"></section>
      <section number="853" description="Italian fiction"></section>
      <section number="854" description="Italian essays"></section>
      <section number="855" description="Italian speeches"></section>
      <section number="856" description="Italian letters"></section>
      <section number="857" description="Italian humor &amp; satire"></section>
      <section number="858" description="Italian miscellaneous writings"></section>
      <section number="859" description="Romanian &amp; related literatures"></section>
    </division>
    <division number="860" description="Spanish &amp; Portuguese literatures">
      <section number="860" description="Spanish &amp; Portuguese literatures"></section>
      <section number="861" description="Spanish poetry"></section>
      <section number="862" description="Spanish drama"></section>
      <section number="863" description="Spanish fiction"></section>
      <section number="864" description="Spanish essays"></section>
      <section number="865" description="Spanish speeches"></section>
      <section number="866" description="Spanish letters"></section>
      <section number="867" description="Spanish humor &amp; satire"></section>
      <section number="868" description="Spanish miscellaneous writings"></section>
      <section number="869" description="Portuguese literature"></section>
    </division>
    <division number="870" description="Latin &amp; Italic literatures">
      <section number="870" description="Italic literature; Latin literature"></section>
      <section number="871" description="Latin poetry"></section>
      <section number="872" description="Latin dramatic poetry &amp; drama"></section>
      <section number="873" description="Latin epic poetry &amp; fiction"></section>
      <section number="874" description="Latin lyric poetry"></section>
      <section number="875" description="Latin speeches"></section>
      <section number="876" description="Latin letters"></section>
      <section number="877" description="Latin humor &amp; satire"></section>
      <section number="878" description="Latin miscellaneous writings"></section>
      <section number="879" description="Literatures of other Italic languages"></section>
    </division>
    <division number="880" description="Classical &amp; modern Greek literatures">
      <section number="880" description="Hellenic literatures; classical Greek"></section>
      <section number="881" description="Classical Greek poetry"></section>
      <section number="882" description="Classical Greek dramatic poetry &amp; drama"></section>
      <section number="883" description="Classical Greek epic poetry &amp; fiction"></section>
      <section number="884" description="Classical Greek lyric poetry"></section>
      <section number="885" description="Classical Greek speeches"></section>
      <section number="886" description="Classical Greek letters"></section>
      <section number="887" description="Classical Greek humor &amp; satire"></section>
      <section number="888" description="Classical Greek miscellaneous writings"></section>
      <section number="889" description="Modern Greek literature"></section>
    </division>
    <division number="890" description="Other literatures">
      <section number="890" description="Literatures of other languages"></section>
      <section number="891" description="East Indo-European &amp; Celtic literatures"></section>
      <section number="892" description="Afro-Asiatic literatures; Semitic literatures"></section>
      <section number="893" description="Non-Semitic Afro-Asiatic literatures"></section>
      <section number="894" description="Altaic, Uralic, Hyperborean &amp; Dravidian"></section>
      <section number="895" description="Literatures of East &amp; Southeast Asia"></section>
      <section number="896" description="African literatures"></section>
      <section number="897" description="North American native literatures"></section>
      <section number="898" description="South American native literatures"></section>
      <section number="899" description="Austronesian &amp; other literatures"></section>
    </division>
  </class>
  <class number="900" description="History &amp; geography">
    <division number="900" description="History">
      <section number="900" description="History &amp; geography"></section>
      <section number="901" description="Philosophy &amp; theory"></section>
      <section number="902" description="Miscellany"></section>
      <section number="903" description="Dictionaries &amp; encyclopedias"></section>
      <section number="904" description="Collected accounts of events"></section>
      <section number="905" description="Serial publications"></section>
      <section number="906" description="Organizations &amp; management"></section>
      <section number="907" description="Education, research &amp; related topics"></section>
      <section number="908" description="Kinds of persons treatment"></section>
      <section number="909" description="World history"></section>
    </division>
    <division number="910" description="Geography &amp; travel">
      <section number="910" description="Geography &amp; travel"></section>
      <section number="911" description="Historical geography"></section>
      <section number="912" description="Atlases, maps, charts &amp; plans"></section>
      <section number="913" description="Geography of &amp; travel in ancient world"></section>
      <section number="914" description="Geography of &amp; travel in Europe"></section>
      <section number="915" description="Geography of &amp; travel in Asia"></section>
      <section number="916" description="Geography of &amp; travel in Africa"></section>
      <section number="917" description="Geography of &amp; travel in North America"></section>
      <section number="918" description="Geography of &amp; travel in South America"></section>
      <section number="919" description="Geography of &amp; travel in other areas"></section>
    </division>
    <division number="920" description="Biography &amp; genealogy">
      <section number="920" description="Biography, genealogy &amp; insignia"></section>
      <section number="921" description="Optional"></section>
      <section number="922" description="Optional"></section>
      <section number="923" description="Optional"></section>
      <section number="924" description="Optional"></section>
      <section number="925" description="Optional"></section>
      <section number="926" description="Optional"></section>
      <section number="927" description="Optional"></section>
      <section number="928" description="Optional"></section>
      <section number="929" description="Genealogy, names &amp; insignia"></section>
    </division>
    <division number="930" description="History of ancient world (to ca. 499)">
      <section number="930" description="History of ancient world to ca. 499"></section>
      <section number="931" description="China to 420"></section>
      <section number="932" description="Egypt to 640"></section>
      <section number="933" description="Palestine to 70"></section>
      <section number="934" description="South Asia to 647"></section>
      <section number="935" description="Mesopotamia &amp; Iranian Plateau to 637"></section>
      <section number="936" description="Europe north &amp; west of Italy to ca. 499"></section>
      <section number="937" description="Italy &amp; adjacent territories to 476"></section>
      <section number="938" description="Greece to 323"></section>
      <section number="939" description="Other parts of ancient world to ca. 640"></section>
    </division>
    <division number="940" description="History of Europe">
      <section number="940" description="History of Europe"></section>
      <section number="941" description="British Isles"></section>
      <section number="942" description="England &amp; Wales"></section>
      <section number="943" description="Central Europe; Germany"></section>
      <section number="944" description="France &amp; Monaco"></section>
      <section number="945" description="Italian Peninsula &amp; adjacent islands"></section>
      <section number="946" description="Iberian Peninsula &amp; adjacent islands"></section>
      <section number="947" description="Eastern Europe; Russia"></section>
      <section number="948" description="Scandinavia"></section>
      <section number="949" description="Other parts of Europe"></section>
    </division>
    <division number="950" description="History of Asia">
      <section number="950" description="History of Asia; Far East"></section>
      <section number="951" description="China &amp; adjacent areas"></section>
      <section number="952" description="Japan"></section>
      <section number="953" description="Arabian Peninsula &amp; adjacent areas"></section>
      <section number="954" description="South Asia; India"></section>
      <section number="955" description="Iran"></section>
      <section number="956" description="Middle East (Near East)"></section>
      <section number="957" description="Siberia (Asiatic Russia)"></section>
      <section number="958" description="Central Asia"></section>
      <section number="959" description="Southeast Asia"></section>
    </division>
    <division number="960" description="History of Africa">
      <section number="960" description="History of Africa"></section>
      <section number="961" description="Tunisia &amp; Libya"></section>
      <section number="962" description="Egypt &amp; Sudan"></section>
      <section number="963" description="Ethiopia &amp; Eritrea"></section>
      <section number="964" description="Northwest African coast &amp; offshore islands"></section>
      <section number="965" description="Algeria"></section>
      <section number="966" description="West Africa &amp; offshore islands"></section>
      <section number="967" description="Central Africa &amp; offshore islands"></section>
      <section number="968" description="Southern Africa; Republic of South Africa"></section>
      <section number="969" description="South Indian Ocean islands"></section>
    </division>
    <division number="970" description="History of North America">
      <section number="970" description="History of North America"></section>
      <section number="971" description="Canada"></section>
      <section number="972" description="Middle America; Mexico"></section>
      <section number="973" description="United States"></section>
      <section number="974" description="Northeastern United States"></section>
      <section number="975" description="Southeastern United States"></section>
      <section number="976" description="South central United States"></section>
      <section number="977" description="North central United States"></section>
      <section number="978" description="Western United States"></section>
      <section number="979" description="Great Basin &amp; Pacific Slope region"></section>
    </division>
    <division number="980" description="History of South America">
      <section number="980" description="History of South America"></section>
      <section number="981" description="Brazil"></section>
      <section number="982" description="Argentina"></section>
      <section number="983" description="Chile"></section>
      <section number="984" description="Bolivia"></section>
      <section number="985" description="Peru"></section>
      <section number="986" description="Colombia &amp; Ecuador"></section>
      <section number="987" description="Venezuela"></section>
      <section number="988" description="Guiana"></section>
      <section number="989" description="Paraguay &amp; Uruguay"></section>
    </division>
    <division number="990" description="History of other areas">
      <section number="990" description="History of other areas"></section>
      <section number="991" description="Unassigned"></section>
      <section number="992" description="Unassigned"></section>
      <section number="993" description="New Zealand"></section>
      <section number="994" description="Australia"></section>
      <section number="995" description="Melanesia; New Guinea"></section>
      <section number="996" description="Other parts of Pacific; Polynesia"></section>
      <section number="997" description="Atlantic Ocean islands"></section>
      <section number="998" description="Arctic islands &amp; Antarctica"></section>
      <section number="999" description="Extraterrestrial worlds"></section>
    </division>
  </class>
</ddc>`
