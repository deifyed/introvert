package html

type Link struct {
	Title   string
	Address string
}

type Section struct {
	Header     string
	Paragraphs []string
}

type Page struct {
	Title      string
	Navigation []Link
	Sections   []Section
}
