package main

import "fmt"

// Sayer 接口
type Sayer interface {
	Say()
}

// Mover 接口
type Mover interface {
	Move()
}

type Dog struct {
	Name string
}

// Say 实现Sayer接口
func (d Dog) Say() {
	fmt.Printf("%s会叫汪汪汪\n", d.Name)
}

// Move 实现Mover接口
func (d Dog) Move() {
	fmt.Printf("%s会动\n", d.Name)
}

func main() {
	var d = Dog{Name: "旺财"}

	var s Sayer = d
	var m Mover = d
	//
	
	s.Say()  // 对Sayer类型调用Say方法
	m.Move() // 对Mov
}
