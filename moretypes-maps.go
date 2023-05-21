package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	result := make(map[string]int)

	for _, word := range strings.Fields(s) {
		// if word is not in result, then value is the zero value of the element type, in this case `int`
		value, _ := result[word]
		result[word] = value + 1
	}

	return result
}

func main() {
	wc.Test(WordCount)
}

/*
Exercise: Maps

Implement WordCount. It should return a map of the counts of each “word” in the string s. The wc.Test function runs a test suite against the provided function and prints success or failure.

You might find [strings.Fields/https://pkg.go.dev/strings#Fields] helpful.
*/
