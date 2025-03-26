package main

import (
	"fmt"
	"sort"
)

func main() {
	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3

	fmt.Println(kidsWithCandies(candies, extraCandies))
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	result := make([]bool, len(candies))
	sortedCandies := make([]int, len(candies))
	copy(sortedCandies, candies)
	sort.Ints(sortedCandies)
	maxCountOfCandies := sortedCandies[len(sortedCandies)-1]
	for i, candiesByChild := range candies {
		if candiesByChild+extraCandies >= maxCountOfCandies {
			result[i] = true
		}
	}
	return result
}
