package main

import "fmt"

/*
Суть заключается в том, чтобы обойти набор эдементов с разными интерфейсами.
В данном случае посетитель-директор школы обходит учителя и ученика, чтобы заставить их работать на уроке.
*/
//элемент обхода
type teacher struct {
}

func (t *teacher) accept(v visitor) {
	v.makeTeach(t)
}

//элемент обхода
type student struct {
}

func (s *student) accept(v visitor) {
	v.makeStudy(s)
}

type visitor interface {
	makeTeach(p *teacher)
	makeStudy(p *student)
}

type classroom interface {
	accept(v visitor)
}

// конкретный посетитель, удовлетворяет интерфейсу visitor
type schoolDirector struct {
}

func (s *schoolDirector) makeTeach(t *teacher) {
	t.teach()
}

func (s *schoolDirector) makeStudy(st *student) {
	st.study()
}

func (s *student) study() {
	fmt.Println("now I'm studying")
}

func (t *teacher) teach() {
	fmt.Println("now I'm teaching")
}

func main() {
	realTeacher := new(teacher)
	realStudent := new(student)
	needToCheck := []classroom{realStudent, realTeacher} //структура для обхода
	realDirector := new(schoolDirector)
	for _, person := range needToCheck {
		person.accept(realDirector)
	}
}
