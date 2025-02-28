package browser

import (
	"fmt"
	"io"
	"strings"

	html_utils "github.com/deifyed/introvert/pkg/html"
	http_utils "github.com/deifyed/introvert/pkg/http"
	"golang.org/x/net/html"
)

func loadURL(url string) (page, error) {
	raw, err := http_utils.OpenPage(url)
	if err != nil {
		return page{}, fmt.Errorf("opening page: %w", err)
	}

	parsedPage, err := parse(raw)
	if err != nil {
		return page{}, fmt.Errorf("parsing body: %w", err)
	}

	return parsedPage, nil
}

func parse(in io.Reader) (page, error) {
	doc, err := html.Parse(in)
	if err != nil {
		return page{}, fmt.Errorf("parsing: %w", err)
	}

	titleTags, err := html_utils.QuerySelect(doc, "h1")
	if err != nil {
		return page{}, fmt.Errorf("selecting title: %w", err)
	}

	anchorTags, err := html_utils.QuerySelect(doc, "a")
	if err != nil {
		return page{}, fmt.Errorf("selecting anchors: %w", err)
	}

	contentTags, err := html_utils.QuerySelect(doc, "p")
	if err != nil {
		return page{}, fmt.Errorf("selecting paragraphs: %w", err)
	}

	links := make([]link, len(anchorTags))
	for index, a := range anchorTags {
		data := a.FirstChild.Data

		links[index] = link{
			title: data,
		}
	}

	var content strings.Builder
	for _, p := range contentTags {
		content.WriteString(p.FirstChild.Data)
	}

	return page{
		Title:      titleTags[0].FirstChild.Data,
		Content:    content.String(),
		navigation: links,
	}, nil
}
