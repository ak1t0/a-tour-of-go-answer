package main

import (
	"code.google.com/p/go-tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	f := strings.Fields(s)
	m := make(map[string]int)

	for _, v := range f {
		m[v] += 1
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
