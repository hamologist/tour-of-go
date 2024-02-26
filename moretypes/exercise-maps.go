package main

import "strings"

func WordCount(s string) map[string]int {
	result := make(map[string]int)

	for _, word := range strings.Fields(s) {
		if _, ok := result[word]; ok {
			result[word] += 1
		} else {
			result[word] = 1
		}
	}

	return result
}
