package main

import (
	"fmt"
)

var prereqs = map[string][]string{
	"algorithms":     {"data structures"},
	"calculus":       {"linear algebra"},
	"linear algebra": {"calculus"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programing"},
	"formal languages":      {"discreate math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	visiting := make(map[string]bool)
	var visitAll func(items []string) error

	visitAll = func(items []string) error {
		for _, item := range items {
			if visiting[item] {
				return fmt.Errorf("%s", item)
			}
			if !seen[item] {
				seen[item] = true
				visiting[item] = true
				if err := visitAll(m[item]); err != nil {
					fmt.Printf("dependency cycle %s\n", err.Error())
				}
				visiting[item] = false
				order = append(order, item) // 子がないものが一番最初に追加される
			}
		}
		return nil
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	if err := visitAll(keys); err != nil {
		fmt.Printf("dependency cycle %s\n", err.Error())
	}
	return order
}
