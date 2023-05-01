package main

import (
	"sort"
)

func main() {
	//staple, drinks, x := []int{10, 20, 5}, []int{5, 5, 2}, 15
	staple, drinks, x := []int{2, 1, 1}, []int{8, 9, 5, 1}, 9

	breakfastNumber(staple, drinks, x)
}
func breakfastNumber(staple []int, drinks []int, x int) int {
	ans := 0
	sort.Ints(drinks)
	for _, val := range staple {
		//二分
		left, right := 0, len(drinks)
		for left < right {
			mid := (left + right) >> 1
			if drinks[mid] <= x-val {
				left = mid + 1
			} else {
				right = mid
			}
		}

		ans += right
		ans %= 1000000007

	}
	return ans
}
