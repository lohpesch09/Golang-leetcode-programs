package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "leetcode"
	fmt.Printf("%s", reverseVowels(s))
}

func reverseVowels(s string) string {
	vowels := "aeiou"
	resultS := ""
	vowelsInS := make([]rune, 0)
	for _, letter := range s {
		for _, vowel := range vowels {
			if strings.ToLower(string(letter)) == string(vowel) {
				vowelsInS = append(vowelsInS, letter)
				break
			}
		}
	}

	counter := len(vowelsInS) - 1

	for _, letter := range s {
		for i, vowel := range vowels {
			if strings.ToLower(string(letter)) == string(vowel) {
				resultS += string(vowelsInS[counter])
				counter--
				break
			} else {
				if i == len(vowels)-1 {
					resultS += string(letter)
				}
			}
		}
	}

	return resultS
}
