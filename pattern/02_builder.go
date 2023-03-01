package main

import "fmt"

/*
Суть паттерна заключается в том, чтобы разбить создание сложного объекта на подзадачи, которые должны выполниться в
определенной последовательности. Данный пример рассматривает написание письма, где по отдельности заполняются все его поля.
*/

type director struct {
	build builder
}

func (d *director) composeMessage() {
	d.build.makeFrom("director")
	d.build.makeTo("wife")
	d.build.makeSubject("dinner")
	d.build.makeContent("I'd like to have a pizza tonight")
}

type builder interface {
	makeSubject(str string)
	makeFrom(str string)
	makeTo(str string)
	makeContent(str string)
}

type mail struct {
	subject string
	from    string
	to      string
	content string
}

func (m *mail) show() {
	fmt.Println("from: ", m.from)
	fmt.Println("to: ", m.to)
	fmt.Println("subject: ", m.subject)
	fmt.Println("content: ", m.content)
}

type concreteBuilder struct {
	message *mail
}

func (b *concreteBuilder) makeSubject(str string) {
	b.message.subject = str
}
func (b *concreteBuilder) makeFrom(str string) {
	b.message.from = str
}
func (b *concreteBuilder) makeTo(str string) {
	b.message.to = str
}
func (b *concreteBuilder) makeContent(str string) {
	b.message.content = str
}

func main() {
	mailMessage := new(mail)
	dir := director{&concreteBuilder{mailMessage}}
	dir.composeMessage()
	mailMessage.show()
}
