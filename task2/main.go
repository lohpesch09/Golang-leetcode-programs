package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str1 := "AA"
	str2 := "A"
	fmt.Println(gcdOfStrings(str1, str2))
}

func gcdOfStrings(str1, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}

	str1Length := utf8.RuneCountInString(str1)
	str2Length := utf8.RuneCountInString(str2)

	prefixBorder := gdc(str1Length, str2Length)

	return str1[:prefixBorder]

}

func gdc(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
