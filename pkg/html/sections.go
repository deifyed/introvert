package html

import (
	"errors"
	"fmt"

	"golang.org/x/net/html"
)

func extractSections(root *html.Node) ([]Section, error) {
	sectionTags, err := querySelect(root, "section")
	if err != nil {
		return nil, fmt.Errorf("querying <section>'s: %w", err)
	}

	sections := make([]Section, len(sectionTags))
	for index, tag := range sectionTags {
		title, err := extractHeader(tag)
		if err != nil {
			return nil, fmt.Errorf("extracting section header: %w", err)
		}

		paragraphs, err := extractParagraphs(tag)
		if err != nil {
			return nil, fmt.Errorf("extracting paragraphs: %w", err)
		}

		sections[index] = Section{
			Header:     title,
			Paragraphs: paragraphs,
		}
	}

	return sections, nil
}

func extractHeader(root *html.Node) (string, error) {
	for _, tag := range []string{"h1", "h2", "h3", "h4", "h5"} {
		elements, _ := querySelect(root, tag)

		if len(elements) == 0 {
			continue
		}

		return elements[0].FirstChild.Data, nil
	}

	return "", errors.New("found no headers")
}

func extractParagraphs(parent *html.Node) ([]string, error) {
	pTags, err := querySelect(parent, "p")
	if err != nil {
		return nil, fmt.Errorf("querying <p>'s: %w", err)
	}

	paragraphs := make([]string, len(pTags))

	for index, node := range pTags {
		paragraphs[index] = node.FirstChild.Data
	}

	return paragraphs, nil
}
