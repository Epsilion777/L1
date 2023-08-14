// Разработать программу, которая переворачивает подаваемую на ход строку
// (например: «главрыба — абырвалг»). Символы могут быть unicode.

package main

import "fmt"

func ReverseString(str string) string {
	runes := []rune(str)
	for i := 0; i < len(runes)/2; i++ {
		runes[i], runes[len(runes)-i-1] = runes[len(runes)-i-1], runes[i]
	}
	return string(runes)
}
func main() {
	str := "главрыба"
	fmt.Println("Original string:", str)
	reversedStr := ReverseString(str)
	fmt.Println("Reversed string:", reversedStr)
}
