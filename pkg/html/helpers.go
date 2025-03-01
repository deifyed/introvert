package html

import (
	"errors"
	"fmt"

	"golang.org/x/net/html"
)

// querySelect searches for all nodes with the given tag name in the HTML tree.
func querySelect(root *html.Node, query string) ([]*html.Node, error) {
	// Traverse the HTML tree recursively to find all matching tags
	var traverse func(*html.Node) []*html.Node

	traverse = func(n *html.Node) []*html.Node {
		var matches []*html.Node

		// If the node is a start tag and matches the query, add it to the results
		if n.Type == html.ElementNode && n.Data == query {
			matches = append(matches, n)
		}

		// Traverse through child nodes
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			matches = append(matches, traverse(c)...)
		}

		return matches
	}

	// Start traversing from the root node
	matchingNodes := traverse(root)
	if len(matchingNodes) == 0 {
		return nil, errors.New("no matching tags found")
	}

	return matchingNodes, nil
}

func extractNavigation(root *html.Node) ([]Link, error) {
	anchorTags, err := querySelect(root, "a")
	if err != nil {
		return nil, fmt.Errorf("selecting anchors: %w", err)
	}

	links := make([]Link, len(anchorTags))
	for index, a := range anchorTags {
		title := a.FirstChild.Data

		links[index] = Link{
			Title: title,
		}
	}

	return links, nil
}

func extractSections(root *html.Node) ([]Section, error) {
	sectionTags, err := querySelect(root, "section")
	if err != nil {
		return nil, fmt.Errorf("querying <section>'s: %w", err)
	}

	sections := make([]Section, len(sectionTags))
	for index, tag := range sectionTags {
		paragraphs, err := extractParagraphs(tag)
		if err != nil {
			return nil, fmt.Errorf("extracting paragraphs: %w", err)
		}

		sections[index] = Section{
			Header:     tag.FirstChild.Data,
			Paragraphs: paragraphs,
		}
	}

	return sections, nil
}

func extractPageHeader(root *html.Node) (string, error) {
	mains, err := querySelect(root, "main")
	if err != nil {
		return "", fmt.Errorf("selecting title: %w", err)
	}

	headers, err := querySelect(mains[0], "h1")
	if err != nil {
		return "", fmt.Errorf("selecting headers: %w", err)
	}

	return headers[0].FirstChild.Data, nil
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
