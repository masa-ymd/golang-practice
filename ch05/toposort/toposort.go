package main

import (
	"fmt"
	"sort"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calcules":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":      {"discrete math"},
	"databases":            {"linear algebra"},
	"discrete math":        {"intro to programing"},
	"formal languages":     {"discrete math"},
	"networks":             {"operating systems"},
	"operating systems":    {"data structures", "computer organization"},
	"programing languages": {"data structures", "computer organization"},
}

func main() {
	for i, course := range topSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	// 無名関数を再帰的に利用する場合は、事前に宣言が必要
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}
