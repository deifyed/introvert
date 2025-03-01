package browser

import (
	"fmt"

	html_utils "github.com/deifyed/introvert/pkg/html"
	http_utils "github.com/deifyed/introvert/pkg/http"
)

func loadURL(url string) (page, error) {
	raw, err := http_utils.OpenPage(url)
	if err != nil {
		return page{}, fmt.Errorf("opening page: %w", err)
	}

	parsedPage, err := html_utils.Parse(raw)
	if err != nil {
		return page{}, fmt.Errorf("parsing body: %w", err)
	}

	return asPage(parsedPage), nil
}

func asLinks(links []html_utils.Link) []link {
	result := make([]link, len(links))

	for index, item := range links {
		result[index] = link{
			title:   item.Title,
			address: item.Address,
		}
	}

	return result
}

func asSections(sections []html_utils.Section) []section {
	result := make([]section, len(sections))

	for index, item := range sections {
		result[index] = section{
			header:     item.Header,
			paragraphs: item.Paragraphs,
		}
	}

	return result
}

func asPage(p html_utils.Page) page {
	return page{
		Title:      p.Title,
		navigation: asLinks(p.Navigation),
		Sections:   asSections(p.Sections),
	}
}
