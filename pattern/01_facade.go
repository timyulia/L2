package main

import "fmt"

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

type vacuumCleaner struct {
}

func (v *vacuumCleaner) startCleaning() {
	fmt.Println("cleaning started")
}

func (v *vacuumCleaner) stopCleaning() {
	fmt.Println("cleaning stopped")
}

type light struct {
}

func (l *light) turnOn() {
	fmt.Println("light is on")
}

func (l *light) turnOff() {
	fmt.Println("light is off")
}

type conditioner struct {
}

func (c *conditioner) setTemp(temp int) {
	fmt.Printf("setting temperature to %d celsius\n", temp)
}

func (c *conditioner) turnOn() {
	fmt.Println("conditioner is on")
}

func (c *conditioner) turnOff() {
	fmt.Println("conditioner is off")
}

type teapot struct {
}

func (t *teapot) boil() {
	fmt.Println("water is boiling")
}

type facadeSmartHome struct {
	tea    teapot
	cond   conditioner
	lig    light
	vacuum vacuumCleaner
}

func (f *facadeSmartHome) comingHome() {
	fmt.Println("____coming home____")
	f.cond.turnOn()
	f.cond.setTemp(23)
	f.vacuum.stopCleaning()
	f.tea.boil()
	f.lig.turnOn()
	fmt.Println()
}

func (f *facadeSmartHome) goingOut() {
	fmt.Println("____going out____")
	f.cond.turnOff()
	f.vacuum.startCleaning()
	f.lig.turnOff()
	fmt.Println()
}

func main() {
	f := facadeSmartHome{
		tea:    teapot{},
		cond:   conditioner{},
		lig:    light{},
		vacuum: vacuumCleaner{},
	}
	f.goingOut()
	f.comingHome()
}
