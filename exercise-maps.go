package main

import (
	"fmt"
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	var words []string = strings.Fields(s)
	var wc = make(map[string]int)

	for i := 0; i < len(words); i++ {
		if _, ok := wc[words[i]]; ok {
			wc[words[i]]++
		} else {
			wc[words[i]] = 1
		}
	}
	fmt.Println(wc)
	return wc
}

func main() {
	wc.Test(WordCount)
}
