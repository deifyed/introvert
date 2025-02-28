package html

import (
	"errors"

	"golang.org/x/net/html"
)

// QuerySelect searches for all nodes with the given tag name in the HTML tree.
func QuerySelect(root *html.Node, query string) ([]*html.Node, error) {
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
