package main

import (
	"fmt"
	"sort"
	"strconv"
)

func printReport(pages map[string]int, baseUrl string) {
	fmt.Println("=============================")
	fmt.Println(" REPORT for " + baseUrl)
	fmt.Println("=============================")

	keys := sortKeys(pages)

	for _, key := range keys {
		visitTimes := strconv.Itoa(pages[key])
		fmt.Println("Found " + visitTimes + " internal links to " + key)
	}
}

func sortKeys(mp map[string]int) []string{
	keys := make([]string, 0, len(mp))

	for k := range mp {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return mp[keys[i]] > mp[keys[j]]
	})

	return keys
}
