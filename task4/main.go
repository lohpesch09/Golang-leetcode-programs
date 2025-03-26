package main

import "fmt"

func main() {
	flowerbed := []int{1, 0, 0, 0, 1, 0, 0}
	n := 2

	fmt.Println(canPlaceFlowers(flowerbed, n))
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}

	for i, existingFlower := range flowerbed {
		if i != 0 && i != len(flowerbed)-1 {
			if existingFlower == 0 {
				if flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
					flowerbed[i] = 1
					n--
					continue
				}
			}
		} else if i == 0 {
			if existingFlower == 0 && len(flowerbed) > 1 {
				if flowerbed[i+1] == 0 {
					flowerbed[i] = 1
					n--
					continue
				}
			} else if existingFlower == 0 {
				flowerbed[i] = 1
				n--
				continue
			}
		} else if i == len(flowerbed)-1 {
			if existingFlower == 0 {
				if flowerbed[i-1] == 0 {
					flowerbed[i] = 1
					n--
					continue
				}
			}
		}
	}
	if n > 0 {
		return false
	}

	return true
}
