package main

import "fmt"

func main() {
	//list := [...]int{1, 3, 5, 7, 8}
	//sum := 0
	//for _, i2 := range list {
	//	sum += i2
	//}
	//fmt.Println(sum)
	//找出数组中和为指定值的两个元素的下标，比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)。
	list := [...]int{1, 3, 5, 7, 8}
	listMap := make(map[int]int)
	for index, value := range list {
		listMap[value] = index
	}

	for index, value := range list {

		val, ok := listMap[8-value]
		if ok {
			fmt.Println(index, val)
			delete(listMap, value)
		}

	}
	//for i := 0; i < len(list); i++ {
	//	for j := i + 1; j < len(list)-1; j++ {
	//		if list[i]+list[j] == 8 {
	//			println(i, j)
	//		}
	//	}
	//}

}
