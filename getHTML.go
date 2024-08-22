package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	res, err := http.Get(rawURL)

	if err != nil {
		return "", err
	}

	if res.StatusCode >= 400 {
		return "", fmt.Errorf("request failed with status code: %v", res.StatusCode)
	}

	if !strings.Contains(res.Header.Get("content-type"), "text/html") {
		return "", fmt.Errorf("invalid content type")
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		return "", err
	}
	
	return string(body), nil
}