package main

import (
	"fmt"
	"reflect"
)

//func main() {
//	nums1 := []int{3, 5, 2, 6}
//	nums2 := []int{3, 1, 7}
//	fmt.Println(minNumber(nums1, nums2))
//}
//func minNumber(nums1 []int, nums2 []int) int {
//	a := nums1
//	sort.Ints(a)
//	b := nums2
//	sort.Ints(b)
//	for _, value := range a {
//		if Contains(b, value) {
//			return value
//		}
//	}
//	for _, value := range b {
//		if Contains(a, value) {
//			return value
//		}
//	}
//	if a[0] > b[0] {
//		return b[0]*10 + a[0]
//	} else if a[0] < b[0] {
//		return b[0] + a[0]*10
//
//	} else {
//		return b[0]
//	}
//}
//func Contains(slice []int, element int) bool {
//	for _, e := range slice {
//		if e == element {
//			return true
//		}
//	}
//
//	return false
//}

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	//b := []string{"aa", "bb", "c"}
	fmt.Println(reflect.TypeOf(a))
	//a = append(a, 1)
	//fmt.Println(a[:3])
	//fmt.Println(a[3:])
	//fmt.Println(a[0:1])
	//fmt.Println(append(a[:3], a[3:]...))

	//ForEach(a, test)
	//ForEach(b, Add)
	//fmt.Println(_map(a, func(index int, item int) int {
	//	return item * 2
	//}))

}

func ForEach[T any](arr []T, fun func(index int, item T)) {
	for i, i2 := range arr {
		fun(i, i2)
	}
}
func _map[T any](arr []T, fun func(index int, item T) T) []T {
	res := make([]T, 0)
	for i, i2 := range arr {
		res = append(res, fun(i, i2))

	}
	return res
}

func test(index int, item int) {
	fmt.Println(item)

}

func Add(index int, item string) {
	fmt.Println(index, item)

}
