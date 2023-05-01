package main

func main() {

}
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])

	for _, item := range matrix {
		if search(item, target, n) {
			return true
		}
	}
	return false
}

func search(arr []int, target, length int) bool {
	left, right := 0, length
	for left < right {
		mid := (left + right) >> 1
		if arr[mid] > target {
			right = mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			return true
		}
	}
	return false
}
