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

// 获取用户输入的学员信息
func getInput() *student {
	var (
		id    int
		name  string
		class string
	)
	fmt.Println("请按要求输入学员信息")
	fmt.Print("请输入学员的学号：")
	fmt.Scanf("%d\n", &id)
	fmt.Print("请输入学员的姓名：")
	fmt.Scanf("%s\n", &name)
	fmt.Print("请输入学员的班级：")
	fmt.Scanf("%s\n", &class)
	// 就能拿到用户输入的学员的所有信息
	stu := newStudent(id, name, class) // 调用student的构造函数造一个学生
	return stu
}

func main() {
	sm := newStudentMgr()
	for {

		showMenu()
		var input int
		fmt.Scanf("%d\n", &input)
		switch input {
		case 1:
			// 添加学员
			stu := getInput()
			sm.addStudent(stu)
		case 2:
			// 编辑学员
			stu := getInput()
			sm.editStudent(stu)
		case 3:
			sm.showAllStudent()
		case 4:
			os.Exit(0)
		}
	}
}
