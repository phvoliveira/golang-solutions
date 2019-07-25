package main

import (
	"fmt"
	"regexp"
	"strings"
)

func isPangram(s []string) string {
	output := []string{}
	for _, m := range s {
		output = append(output, hasAllAlphabeticLetters(m))
	}
	return strings.Join(output, "")
}

func hasAllAlphabeticLetters(s string) string {
	count := 0
	m := make(map[string]bool)
	r, _ := regexp.Compile("[a-z]+")

	for _, v := range s {
		l := string(v)
		if r.MatchString(l) {
			if !m[l] {
				m[l] = true
				count++
			}
		}
	}

	if count == 26 {
		return "1"
	} else {
		return "0"
	}
}

func main() {
	fmt.Println(isPangram([]string{"the quick brown fox jumps over the lazy dog"}))
}
