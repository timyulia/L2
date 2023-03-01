package main

import "fmt"

/*
Запрос передается по цепочке объектов, которые могут его обработать.
В данном случае запрос на выполнение задачи передаются от сеньора к стажеру и в соответствии с уровнем задачи
обрабатывается кем-то из четырех людей.
*/

const (
	intern = iota
	junior
	middle
	senior
)

type handler interface {
	doTask(level int) string
}

type handlerSenior struct {
	next handler
}

func (h *handlerSenior) doTask(level int) string {
	var res string
	if level == senior {
		res = "a senior is doing this task"
		return res
	}
	if h.next != nil {
		res = h.next.doTask(level)
	}
	return res
}

type handlerMiddle struct {
	next handler
}

func (h *handlerMiddle) doTask(level int) string {
	var res string
	if level == middle {
		res = "a middle is doing this task"
		return res
	}
	if h.next != nil {
		res = h.next.doTask(level)
	}
	return res
}

type handlerJunior struct {
	next handler
}

func (h *handlerJunior) doTask(level int) string {
	var res string
	if level == junior {
		res = "a junior is doing this task"
		return res
	}
	if h.next != nil {
		res = h.next.doTask(level)
	}
	return res
}

type handlerIntern struct {
	next handler
}

func (h *handlerIntern) doTask(level int) string {
	var res string
	if level == intern {
		res = "an intern is doing this task"
		return res
	}
	if h.next != nil {
		res = h.next.doTask(level)
	}
	return res
}

func main() {
	handlersChain := &handlerSenior{&handlerMiddle{&handlerJunior{&handlerIntern{}}}}
	fmt.Println(handlersChain.doTask(1))
}
