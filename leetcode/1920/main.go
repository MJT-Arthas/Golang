package main

func main() {
	nums := []int{0, 2, 1, 5, 3, 4}
	buildArray(nums)
}

func buildArray(nums []int) []int {
	var ans []int
	for _, item := range nums {
		ans = append(ans, nums[item])
	}
	return ans
}
