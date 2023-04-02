package main

import (
	"fmt"
	"os"
)

func showMenu() {
	fmt.Println("欢迎访问")

	fmt.Println("1.add")
	fmt.Println("2.edit")
	fmt.Println("3.show")
	fmt.Println("4.exit")

}

func main() {
	showMenu()
	var input int
	fmt.Scanf("%d\n", &input)
	switch input {
	case 1:
	case 2:
	case 3:
	case 4:
		os.Exit(0)
	}
}
