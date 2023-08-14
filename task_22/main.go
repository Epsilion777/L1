// Разработать программу, которая перемножает, делит, складывает, вычитает две
// числовых переменных a,b, значение которых > 2^20.

package main

import (
	"fmt"
	"math/big"
)

func main() {
	firstNumber, secondNumber, result := new(big.Int), new(big.Int), new(big.Int)
	firstNumber.SetString("55555555555555555555555555555555555555555555555555555555555555", 10)
	secondNumber.SetString("44444444444444444444444444444444444444444444444444444444444444", 10)
	fmt.Printf("First number: %s\nSecond number: %s\n", firstNumber, secondNumber)

	// Деление
	result.Div(firstNumber, secondNumber)
	fmt.Println("Result of division:", result)

	// Очистка результата
	result.And(result, big.NewInt(0))

	// Умножение
	result.Mul(firstNumber, secondNumber)
	fmt.Println("Result of multiplication:", result)

	// Очистка результата
	result.And(result, big.NewInt(0))

	// Сумма
	result.Add(firstNumber, secondNumber)
	fmt.Println("Result of sum:", result)

	// Очистка результата
	result.And(result, big.NewInt(0))

	// Разность
	result.Sub(firstNumber, secondNumber)
	fmt.Println("Result of subtraction:", result)

}
