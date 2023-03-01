package main

import "fmt"

/*
Паттерн заключается в том, что у объекта может быть несколько состояний, в зависимости от которых выполняются разные
действия.
В данном случае есть человек, у которого состояние-потребность в еде или сне.
*/

type need interface {
	act() string
}

type person struct {
	state need
}

func (a *person) act() string {
	return a.state.act()
}

func (a *person) setState(state need) {
	a.state = state
}

func newPerson() *person {
	return &person{state: &eatNeed{}}
}

type eatNeed struct {
}

func (a *eatNeed) act() string {
	return "eating..."
}

type sleepNeed struct {
}

func (a *sleepNeed) act() string {
	return "sleeping..."
}

func main() {
	per := newPerson()
	fmt.Println(per.act())
	per.setState(&sleepNeed{})
	fmt.Println(per.act())
}
