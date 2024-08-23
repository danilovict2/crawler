package main

import (
	"fmt"
	"os"
	"strconv"
)


func main() {
	if len(os.Args) < 4 {
		fmt.Println("not enough arguments provided")
		os.Exit(1)
	}

	if len(os.Args) > 4 {
		fmt.Println("too many arguments provided")
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

	cfg, err := configure(os.Args[1], maxConcurrency, maxPages)
	if err != nil {
		fmt.Printf("Error - configure: %v", err)
		return
	}

	cfg.wg.Add(1)
	go cfg.crawlPage(os.Args[1])
	cfg.wg.Wait()

	for key, val := range cfg.pages {
		fmt.Printf("Visited %s url %d times\n", key, val)
	}
}