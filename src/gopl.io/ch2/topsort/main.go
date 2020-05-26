package main

import (
	"fmt"
	"sort"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func topSort(prereqs map[string][]string) []string {
	seen := make(map[string]bool)
	var order []string

	var visitAll func(courses []string)
	visitAll = func(courses []string) {
		for _, course := range courses {
			if !seen[course] {
				visitAll(prereqs[course])
				seen[course] = true
				order = append(order, course)
			}
		}
	}

	var keys []string
	for key := range prereqs {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	visitAll(keys)
	return order
}

func main() {
	for i, course := range topSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}
