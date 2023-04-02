package main

func main() {
	s := "adaa"
	chars := "d"
	vals := []int{-1000}
	maximumCostSubstring(s, chars, vals)
}

func maximumCostSubstring(s string, chars string, vals []int) int {
	Map := make(map[string]int)
	for i := 0; i < 26; i++ {
		Map[string(int('a')+i)] = i + 1
	}

	for index, s := range chars {
		Map[string(s)] = vals[index]
	}

	return 1
}
