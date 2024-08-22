package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 sync.Mutex
	concurrencyControl chan struct{}
	wg                 sync.WaitGroup
	maxPages int
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("not enough arguments provided")
		os.Exit(1)
	}

	if len(os.Args) > 4 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	baseURL, err := url.Parse(os.Args[1])
	if err != nil {
		fmt.Printf("Error: couldn't parse URL '%s': %v\n", baseURL, err)
		os.Exit(1)
	}

	maxConcurrency, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	maxPages, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	cfg := &config {
		pages: map[string]int{},
		baseURL: baseURL,
		concurrencyControl: make(chan struct{}, maxConcurrency),
		maxPages: maxPages,
	}

	cfg.wg.Add(1)
	go cfg.crawlPage(baseURL.String())
	cfg.wg.Wait()

	for key, val := range cfg.pages {
		fmt.Printf("Visited %s url %d times\n", key, val)
	}
}