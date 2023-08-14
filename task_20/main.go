// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow»

package main

import (
	"fmt"
	"strings"
)

func ReverseWordsString(str string) string {
	reversedStr := strings.Builder{}
	slcWords := strings.Split(str, " ")
	for i := len(slcWords) - 1; i >= 0; i-- {
		if i != 0 {
			reversedStr.WriteString(fmt.Sprintf("%s ", slcWords[i]))
		} else {
			// Если элемент последний, то добавлять " " не нужно
			reversedStr.WriteString(slcWords[i])
		}
	}
	return reversedStr.String()
}
func main() {
	str := "snow dog sun"
	fmt.Println("Original string:", str)
	reversedWordsStr := ReverseWordsString(str)
	fmt.Println("Reversed string:", reversedWordsStr)
}
