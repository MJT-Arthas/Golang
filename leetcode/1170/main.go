package main

import (
	"fmt"
	"sort"
)

func main() {
	queries, words := []string{"bbb", "cc"}, []string{"a", "aa", "aaa", "aaaa"}
	fmt.Println(numSmallerByFrequency(queries, words))
}

func numSmallerByFrequency(queries []string, words []string) []int {
	wordsArr := calc(words)
	sort.Ints(wordsArr)
	var ans []int
	for _, val := range calc(queries) {
		ans = append(ans, search(wordsArr, val))
	}
	return ans
}

func calc(arr []string) []int {
	var res []int
	for _, item := range arr {
		var tempArr []int
		tempMap := make(map[int]int)
		for _, s := range item {
			tempArr = append(tempArr, int(s))
			tempMap[int(s)]++
		}
		sort.Ints(tempArr)

		res = append(res, tempMap[tempArr[0]])
	}

	return res
}
func search(arr []int, target int) int {
	left, right := 0, len(arr)
	for left < right {
		mid := (left + right) >> 1
		if arr[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return len(arr) - left
}
