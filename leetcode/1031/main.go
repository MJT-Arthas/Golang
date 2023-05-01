package main

import "fmt"

func main() {
	nums := []int{0, 6, 5, 2, 2, 5, 1, 9, 4}

	fmt.Println(maxSumTwoNoOverlap(nums, 1, 2))
}
func help(nums []int, firstLen int, secondLen int) int {
	suml := 0
	for i := 0; i < firstLen; i++ {
		suml += nums[i]
	}
	maxSumL := suml
	sumr := 0
	for i := firstLen; i < firstLen+secondLen; i++ {
		sumr += nums[i]
	}
	res := suml + sumr
	for i, j := firstLen+secondLen, firstLen; i < len(nums); i, j = i+1, j+1 {
		suml += nums[j] - nums[j-firstLen]
		maxSumL = max(maxSumL, suml)
		sumr += nums[i] - nums[i-secondLen]
		res = max(res, maxSumL+sumr)
	}

	return res
}

func maxSumTwoNoOverlap(nums []int, firstLen int, secondLen int) int {
	return max(help(nums, firstLen, secondLen), help(nums, secondLen, firstLen))
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
