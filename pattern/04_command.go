package main

import "fmt"

/*
Представление запроса в виде объекта. Есть интерфейс command и тип, удовлетворяющей ему incCommand.
В виде исполнителя выступает другой тип-increment. Invoker содержит очередь запросов.
*/

type command interface {
	execute()
}

type incCommand struct {
	inc *increment
}

func (incCom *incCommand) execute() {
	incCom.inc.increase()
}

// receiver
type increment struct {
	num int
}

func (inc *increment) increase() {
	inc.num++
}

type invoker struct {
	coms []command
}

func (i *invoker) add(com command) {
	i.coms = append(i.coms, com)
}

func (i *invoker) cancel() {
	if len(i.coms) > 0 {
		i.coms = i.coms[:len(i.coms)-1]
	}
}

func (i *invoker) do() {
	for _, com := range i.coms {
		com.execute()
	}
}

func main() {
	receiver := new(increment)
	concreteCommand := incCommand{receiver}
	invo := invoker{}
	for i := 0; i < 5; i++ {
		invo.add(&concreteCommand)
	}
	invo.cancel()
	invo.do()
	fmt.Println(receiver.num)
}
