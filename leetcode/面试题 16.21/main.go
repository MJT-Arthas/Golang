package main

import (
	"fmt"
	"sort"
)

func main() {
	array1, array2 := []int{4, 1, 2, 1, 1, 2}, []int{3, 6, 3, 3}
	fmt.Println(findSwapValues(array1, array2))
}

func findSwapValues(array1 []int, array2 []int) []int {
	sum1, sum2 := 0, 0
	for _, val := range array1 {
		sum1 += val
	}
	for _, val := range array2 {
		sum2 += val
	}
	diff := sum1 - sum2
	sort.Ints(array2)
	for _, val := range array1 {
		left, right := 0, len(array2)-1
		for left < right {
			mid := (left + right) >> 1
			if 2*array2[mid] > 2*val-diff {
				right = mid
			} else if (2 * array2[mid]) < 2*val-diff {
				left = mid + 1
			} else {
				return []int{val, array2[mid]}
			}
		}
	}
	return []int{}
}
