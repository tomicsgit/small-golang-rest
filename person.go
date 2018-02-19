package main

import (
	"strconv"
)

type Human struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Student struct {
	Human
	School string  `json:"school"`
	Loan   float32 `loan:"loan"`
}

type Employee struct {
	Human
	Company string  `json:"company"`
	Salary  float32 `json:"salary"`
}

// define interfaces
type ManInterface interface {
	SetId(id int, e error)
	GetId() int
	GetName() string
	SayHi() string
	Sing(lyrics string) string
}

func (h Human) SetId(id int, e error) {
	h.Id = id
}

func (h Human) GetId() int {
	return h.Id
}

func (h Human) GetName() string {
	return h.Name
}

func (h Human) SayHi() string {
	return "My name is " + h.Name + " and I'm " + strconv.Itoa(h.Age) + " years old!"
}

func (h Human) Sing(lyrics string) string {
	return "Jebem li gaaaaa " + lyrics
}

func (e Employee) SayHi() string {
	return "I'm a decent gye " + e.Name + " working at " + e.Company + "!"
}

func (e Employee) SpendSalary(amount float32) {
	e.Salary -= amount // More vodka please!!! Get me through the day!
}

func (s Student) BorrowMoney(amount float32) {
	s.Loan += amount // (again and again and...)
}
