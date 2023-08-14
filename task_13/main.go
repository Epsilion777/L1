// Поменять местами два числа без создания временной переменной.

package main

import "fmt"

func main() {
	a := 5
	b := 10
	fmt.Printf("a = %d,\t b = %d\n", a, b) // 5, 10

	// Метод swap
	a, b = b, a

	fmt.Printf("a = %d,\t b = %d\t\n", a, b) // 10, 5

	// Метод XOR (Не работает с типом float)
	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Printf("a = %d,\t b = %d\n", a, b) // 5, 10

	// Метод + и -
	a = a + b
	b = a - b
	a = a - b

	fmt.Printf("a = %d,\t b = %d\n", a, b) // 10, 5

	// Метод * и / ( Не сработает, если возникнет переполнение переменной )
	a = a * b
	b = a / b
	a = a / b

	fmt.Printf("a = %d,\t b = %d\n", a, b) // 5, 10
}
