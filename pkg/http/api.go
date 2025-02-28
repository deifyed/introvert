package http

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func OpenPage(url string) (io.Reader, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Add("User-Agent", "Introvert/0.0.1")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("doing request: %w", err)
	}
	defer resp.Body.Close()

	var buff bytes.Buffer

	_, err = io.Copy(&buff, resp.Body)
	if err != nil {
		return nil, fmt.Errorf("buffering: %w", err)
	}

	return &buff, nil
}
