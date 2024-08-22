package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/url"
	"strings"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse base URL: %v", err)
	}

	head, err := html.Parse(strings.NewReader(htmlBody))

	if err != nil {
		return nil, fmt.Errorf("couldn't parse HTML: %v", err)
	}

	var URLs []string
	traverseParseTree(head, &URLs, *baseURL)

	return URLs, nil
}

func traverseParseTree(node *html.Node, URLs *[]string, baseURL url.URL) {
	if node == nil {
		return
	}

	if node.Data == "a" {
		for _, attr := range node.Attr {
			if attr.Key == "href" {
				val, err := url.Parse(attr.Val)
				if err != nil {
					fmt.Printf("couldn't parse href '%v': %v\n", attr.Val, err)
					continue
				}

				resolvedUrl := baseURL.ResolveReference(val)
				*URLs = append(*URLs, resolvedUrl.String())
			}
		}
	}

	traverseParseTree(node.FirstChild, URLs, baseURL)
	traverseParseTree(node.NextSibling, URLs, baseURL)
}
