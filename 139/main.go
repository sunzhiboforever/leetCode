package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "catsandog"
	wordDict := []string{"cats", "dog", "sand", "and", "cat"}
	r := wordBreak(s, wordDict)
	fmt.Println(r)
}

func wordBreak(s string, wordDict []string) bool {

}

func wordBreakR(cache map[string]bool, s string, wordDict []string) bool {
	for _, str := range wordDict {
		if strings.Contains(s, str) == false {
			return false
		}
		yes := wordBreakR(cache, strings.Replace(s, str, "", 1), wordDict)
	}
}
