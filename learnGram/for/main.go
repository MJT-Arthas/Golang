package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		for j := i; j < 10; j++ {
			fmt.Print(i, "X", j, "=", i*j)
			fmt.Print("  ")
		}
		fmt.Println()
	}

}
