package main

import "fmt"

type student struct {
	id    int
	name  string
	class string
}

func newStudent(id int, name, class string) *student {
	return &student{
		id:    id,
		name:  name,
		class: class,
	}
}

type studentMgr struct {
	allStudent []*student
}

func newStudentMgr() *studentMgr {
	return &studentMgr{allStudent: make([]*student, 0, 100)}
}

func (s *studentMgr) addStudent(newStu *student) {
	s.allStudent = append(s.allStudent, newStu)
}

func (s *studentMgr) editStudent(newStu *student) {
	for i, student := range s.allStudent {
		if student.id == newStu.id {
			s.allStudent[i] = newStu
		}
	}
}
func (s *studentMgr) showAllStudent() {
	for _, student := range s.allStudent {
		fmt.Println(student.id, student.name, student.class)
	}
}
