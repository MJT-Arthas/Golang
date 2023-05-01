package main

import (
	"fmt"
	"sort"
)

func main() {
	names := []string{"Mary", "John", "Emma"}
	heights := []int{180, 165, 170}
	fmt.Println(sortPeople(names, heights))
}

//
//func sortPeople(names []string, heights []int) []string {
//	mp := make(map[int]string)
//	for i, height := range heights {
//		mp[height] = names[i]
//	}
//	heights2 := heights[:]
//	sort.Sort(sort.Reverse(sort.IntSlice(heights2)))
//	res := make([]string, 0)
//	for _, h := range heights2 {
//		res = append(res, mp[h])
//	}
//	return res
//}

func sortPeople(names []string, heights []int) []string {
	n := len(names)
	indices := make([]int, n)
	for i := 0; i < n; i++ {
		indices[i] = i
	}
	sort.Slice(indices, func(i, j int) bool {
		return heights[indices[j]] < heights[indices[i]]
	})
	res := make([]string, n)
	for i := 0; i < n; i++ {
		res[i] = names[indices[i]]
	}
	return res
}
