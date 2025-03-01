package html

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestQuerySelectForItems(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		withHTML      string
		withQuery     string
		expectMatches int
	}{
		"should find title": {
			withHTML:      "<html><title></title></html>",
			withQuery:     "title",
			expectMatches: 1,
		},
		"should find all paragraphs": {
			withHTML:      "<html><body><section><p></p><p></p></section></body></html>",
			withQuery:     "p",
			expectMatches: 2,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			rawDocument := strings.NewReader(testCase.withHTML)

			document, err := html.Parse(rawDocument)
			if err != nil {
				t.Fatal("parsing document")
			}

			result, err := querySelect(document, testCase.withQuery)
			if err != nil {
				t.Fatalf("selecting: %s", err.Error())
			}

			if testCase.expectMatches != len(result) {
				t.Fatalf("expected matches %d did not match %d", testCase.expectMatches, len(result))
			}
		})
	}
}

func TestQuerySelectForData(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		withHTML   string
		withQuery  string
		expectData string
	}{
		"should find title": {
			withHTML:   "<html><title>Awesome!</title></html>",
			withQuery:  "title",
			expectData: "Awesome!",
		},
		"should find all paragraphs": {
			withHTML:   "<html><body><section><p>Hey all</p></section></body></html>",
			withQuery:  "p",
			expectData: "Hey all",
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			rawDocument := strings.NewReader(testCase.withHTML)

			document, err := html.Parse(rawDocument)
			if err != nil {
				t.Fatal("parsing document")
			}

			result, err := querySelect(document, testCase.withQuery)
			if err != nil {
				t.Fatalf("selecting: %s", err.Error())
			}

			gotData := result[0].FirstChild.Data

			if testCase.expectData != gotData {
				t.Fatalf("expected data %s did not match received data %s", testCase.expectData, gotData)
			}
		})
	}
}
