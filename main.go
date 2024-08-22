package main

import (
	"fmt"
	"net/url"
	"os"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 sync.Mutex
	concurrencyControl chan struct{}
	wg                 sync.WaitGroup
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	}

	if len(os.Args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	baseURL, err := url.Parse(os.Args[1])
	if err != nil {
		fmt.Printf("Error: couldn't parse URL '%s': %v\n", baseURL, err)
	}

	ch := make(chan struct{}, 4)
	cfg := &config {
		pages: map[string]int{},
		baseURL: baseURL,
		concurrencyControl: ch,
	}

	cfg.wg.Add(1)
	go cfg.crawlPage(baseURL.String())
	cfg.wg.Wait()

	for key, val := range cfg.pages {
		fmt.Printf("Visited %s url %d times\n", key, val)
	}
}