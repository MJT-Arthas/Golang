package main

import (
	"fmt"
)

func main() {
	var arr = []int{-2, 0, 10, -19, 4, 6, -8}

	fmt.Println(checkIfExist(arr))
}
func checkIfExist(arr []int) bool {
	var countMap = make(map[int]int)

	for _, item := range arr {
		countMap[item]++
	}

	for _, item := range arr {
		if item == 0 && countMap[0] > 1 {
			return true
		} else if item != 0 && countMap[2*item] > 0 {
			return true
		}
	}

	return false
}
