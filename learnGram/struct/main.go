package main

import "fmt"

type student struct {
	name string
	age  int
}

func main() {
	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}
	for _, stu := range stus {

		m[stu.name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
}

////Person 结构体
//type Person struct {
//	name string
//	age  int8
//}
//
////NewPerson 构造函数
//func NewPerson(name string, age int8) *Person {
//	return &Person{
//		name: name,
//		age:  age,
//	}
//}
//
////Dream Person做梦的方法
//func (p Person) Dream() {
//	fmt.Printf("%s的梦想是学好Go语言！\n", p.name)
//}
//
//func main() {
//	p1 := NewPerson("小王子", 25)
//	p1.Dream()
//}
