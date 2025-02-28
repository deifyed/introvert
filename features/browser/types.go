package browser

type link struct {
	title   string
	address string
}

type section struct {
	header     string
	paragraphs []string
}

type page struct {
	Title      string
	navigation []link
	Sections   []section
	Content    string
}
