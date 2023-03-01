package main

import "fmt"

/*
Метод помогает создавать разные продукты, обладающие одним  интерфейсом.
В данном случае производятся шоколадки, которые имеют метод eat.
*/

type creator interface {
	make(title string) product
}

type product interface {
	eat()
}

type concreteCreator struct{}

type snickers struct {
	slogan string
}

func (k *snickers) eat() {
	fmt.Println(k.slogan)
}

type kitkat struct {
	slogan string
}

func (k *kitkat) eat() {
	fmt.Println(k.slogan)
}

func (c *concreteCreator) make(title string) product {
	var prod product
	switch title {
	case "kitkat":
		prod = &kitkat{"have a break have a kit kat"}
	case "snickers":
		prod = &snickers{"you're not you when you're hungry"}
	}
	return prod
}

func main() {
	var create creator
	create = &concreteCreator{}
	create.make("kitkat").eat()
}
