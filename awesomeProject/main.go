package main

import "fmt"

type user struct {
	id   int
	name string
}

func main() {
	u := user{id: 0, name: "arthas"}
	fmt.Print(u)
}
