package browser

import (
	"fmt"
	"io"

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

func makeHeader(root *html.Node) (string, error) {
	titleTags, err := html_utils.QuerySelect(root, "h1")
	if err != nil {
		return "", fmt.Errorf("selecting title: %w", err)
	}

	return titleTags[0].FirstChild.Data, nil
}

func makeNavigation(root *html.Node) ([]link, error) {
	anchorTags, err := html_utils.QuerySelect(root, "a")
	if err != nil {
		return nil, fmt.Errorf("selecting anchors: %w", err)
	}

	links := make([]link, len(anchorTags))
	for index, a := range anchorTags {
		title := a.FirstChild.Data

		links[index] = link{
			title: title,
		}
	}

	return links, nil
}

func extractParagraphs(parent *html.Node) ([]string, error) {
	pTags, err := html_utils.QuerySelect(parent, "p")
	if err != nil {
		return nil, fmt.Errorf("querying <p>'s: %w", err)
	}

	paragraphs := make([]string, len(pTags))

	for index, node := range pTags {
		paragraphs[index] = node.FirstChild.Data
	}

	return paragraphs, nil
}

func makeSections(root *html.Node) ([]section, error) {
	sectionTags, err := html_utils.QuerySelect(root, "section")
	if err != nil {
		return nil, fmt.Errorf("querying <section>'s: %w", err)
	}

	sections := make([]section, len(sectionTags))
	for index, tag := range sectionTags {
		paragraphs, err := extractParagraphs(tag)
		if err != nil {
			return nil, fmt.Errorf("extracting paragraphs: %w", err)
		}

		sections[index] = section{
			header:     tag.FirstChild.Data,
			paragraphs: paragraphs,
		}
	}

	return sections, nil
}

func parse(in io.Reader) (page, error) {
	root, err := html.Parse(in)
	if err != nil {
		return page{}, fmt.Errorf("parsing: %w", err)
	}

	header, err := makeHeader(root)
	if err != nil {
		return page{}, fmt.Errorf("making header: %w", err)
	}

	navigation, err := makeNavigation(root)
	if err != nil {
		return page{}, fmt.Errorf("making navigation: %w", err)
	}

	sections, err := makeSections(root)
	if err != nil {
		return page{}, fmt.Errorf("making sections: %w", err)
	}

	return page{
		Title:      header,
		Sections:   sections,
		navigation: navigation,
	}, nil
}
