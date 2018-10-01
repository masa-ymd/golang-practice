package main

import (
	"fmt"
)

func main() {
	fmt.Println(max(1, 2, 3, 4, 1, 3))
	fmt.Println(max())
	fmt.Println(min(1, 2, 3, 4, 1, 3))
	fmt.Println(min())
}

func max(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("%s\n", "no argument")
	}
	var max int

	for _, v := range vals {
		if v > max {
			max = v
		}
	}

	return max, nil
}

func min(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("%s\n", "no argument")
	}
	var min int

	for i, v := range vals {
		if i == 0 {
			min = v
			continue
		}
		if v < min {
			min = v
		}
	}

	return min, nil
}
