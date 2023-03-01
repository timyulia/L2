package main

import (
	"fmt"
	"sort"
)

/*
Схожие алгоритмы объединяются в один интерфейс, который является полем структуры context. Таким образом алгоритмы могут
заменять друг друга.
*/

type strategyAverage interface {
	find(nums []int) int
}

type median struct {
}

func (m *median) find(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	return nums[n/2]
}

type avg struct {
}

func (m *avg) find(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum / len(nums)
}

type contextAverage struct {
	way strategyAverage
}

func (c *contextAverage) set(way strategyAverage) {
	c.way = way
}

func (c *contextAverage) execute(nums []int) int {
	return c.way.find(nums)
}

func main() {
	nums := []int{1, 5, 8, 2, 4}
	m := avg{}
	context := contextAverage{}
	context.set(&m)
	fmt.Println(context.execute(nums))
}
