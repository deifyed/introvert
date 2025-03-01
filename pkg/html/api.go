package html

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

func Parse(in io.Reader) (Page, error) {
	root, err := html.Parse(in)
	if err != nil {
		return Page{}, fmt.Errorf("parsing: %w", err)
	}

	header, err := extractPageHeader(root)
	if err != nil {
		return Page{}, fmt.Errorf("making header: %w", err)
	}

	navigation, err := extractNavigation(root)
	if err != nil {
		return Page{}, fmt.Errorf("making navigation: %w", err)
	}

	sections, err := extractSections(root)
	if err != nil {
		return Page{}, fmt.Errorf("making sections: %w", err)
	}

	return Page{
		Title:      header,
		Sections:   sections,
		Navigation: navigation,
	}, nil
}
