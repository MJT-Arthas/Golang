package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{3, 5, 2, 6}
	nums2 := []int{3, 1, 7}
	fmt.Println(minNumber(nums1, nums2))
}
func minNumber(nums1 []int, nums2 []int) int {
	a := nums1
	sort.Ints(a)
	b := nums2
	sort.Ints(b)
	for _, value := range a {
		if Contains(b, value) {
			return value
		}
	}
	for _, value := range b {
		if Contains(a, value) {
			return value
		}
	}
	if a[0] > b[0] {
		return b[0]*10 + a[0]
	} else if a[0] < b[0] {
		return b[0] + a[0]*10

	} else {
		return b[0]
	}
}
func Contains(slice []int, element int) bool {
	for _, e := range slice {
		if e == element {
			return true
		}
	}

	return false
}
