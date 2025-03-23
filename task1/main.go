package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	word1 := "ab"
	word2 := "pqrs"
	fmt.Println(mergeArternately(word1, word2))
}

func mergeArternately(word1, word2 string) string {
	word1Length := utf8.RuneCountInString(word1)
	word2Length := utf8.RuneCountInString(word2)
	mergedWordLength := word1Length + word2Length

	word1Counter := 0
	word2Counter := 0

	mergedWordInBytes := make([]byte, mergedWordLength)

	for i := 0; i < mergedWordLength; i++ {
		if i%2 == 0 && word1Counter < word1Length {
			mergedWordInBytes[i] = word1[word1Counter]
			word1Counter++
		} else if i%2 != 0 && word2Counter < word2Length {
			mergedWordInBytes[i] = word2[word2Counter]
			word2Counter++
		} else if word1Counter < word1Length {
			mergedWordInBytes[i] = word1[word1Counter]
			word1Counter++
		} else if word2Counter < word2Length {
			mergedWordInBytes[i] = word2[word2Counter]
			word2Counter++
		}
	}
	mergedWord := string(mergedWordInBytes)

	return mergedWord
}
