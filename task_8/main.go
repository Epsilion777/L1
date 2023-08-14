// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
package main

import "fmt"

func SetBit(num *int64, index int) {
	mask := 1 << index
	*num = *num | int64(mask)
}

func ClearBit(num *int64, index int) {
	mask := 1 << index
	// Инвертируем маску
	mask = ^mask
	*num = *num & int64(mask)
}

func main() {
	var num int64 = 128
	var index int = 4

	fmt.Printf("Binary: %b, Decimal: %d\n", num, num)
	SetBit(&num, index)

	fmt.Printf("After SetBit %d\n", index)
	fmt.Printf("Binary: %b, Decimal: %d\n", num, num)
	ClearBit(&num, 4)
	fmt.Printf("After ClearBit %d\n", index)
	fmt.Printf("Binary: %b, Decimal: %d\n", num, num)
}
