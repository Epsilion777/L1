// Разработать программу, которая в рантайме способна определить тип
// переменной: int, string, bool, channel из переменной типа interface{}.

package main

import (
	"fmt"
	"math"
	"reflect"
)

func PrintType(i interface{}) {
	// Используем type assertion для определения типа переменной
	switch v := i.(type) {
	case int:
		fmt.Printf("Type: int, value: %d\n", v)
	case string:
		fmt.Printf("Type: string, value: %s\n", v)
	case bool:
		fmt.Printf("Type: bool, value: %v\n", v)
	case chan int:
		fmt.Printf("Type: chan, value: %v\n", v)
	default:
		fmt.Printf("Unexpected type for value: %v\n", v)
	}

}

func main() {
	elems := []interface{}{10, "10", false, make(chan int), math.Pi}

	// Метод с использованием type assertion
	for _, v := range elems {
		PrintType(v)
	}

	// Метод с использованием reflect.TypeOf()
	fmt.Println("Using reflect")
	for _, v := range elems {
		fmt.Printf("Type: %v, value: %v\n", reflect.TypeOf(v), v)
	}

}
