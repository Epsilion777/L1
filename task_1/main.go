// Дана структура Human (с произвольным набором полей и методов). Реализовать встраивание
// методов в структуре Action от родительской структуры Human (аналог наследования).
package main

import (
	"fmt"
	"log"
)

// Описание структуры Human
type Human struct {
	name string
	age  int
}

// Уникальный метода для структуры Human
func (h *Human) Talk(message string) {
	fmt.Printf("%s: %s\n", h.name, message)
}

// Не уникальный метода для структуры Human
func (h *Human) Describe() {
	fmt.Printf("My name is %s, i`m %d\n", h.name, h.age)
}

// Описание структуры Action
type Action struct {
	Human
	ActionName string
}

// Уникальный метода для структуры Action
func (a *Action) StartAction() {
	log.Printf("%s started %s\n", a.name, a.ActionName)
}

// Переопределение родительского метода Describe структуры Human
func (a *Action) Describe() {
	fmt.Printf("My name is %s, i`m %d, i usually %s\n", a.name, a.age, a.ActionName)
}

func main() {
	human := Human{
		"Dmitriy",
		22,
	}
	// Вызов собственного метода
	human.Talk("Hello guys!")

	// Вызов собственного метода
	human.Describe()

	action := Action{
		human,
		"dig",
	}
	// Вызов собственного метода
	action.StartAction()

	// Вызов родительского метода
	action.Talk("Hello world!")

	// Вызов переопредленного родительского метода
	action.Describe()
}
