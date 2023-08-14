// Разработать программу, которая проверяет, что все символы в строке
// уникальные (true — если уникальные, false etc). Функция проверки должна быть
// регистронезависимой.
// Например:
// abcd — true
// abCdefAaf — false
// aabcd — false
package main

import (
	"fmt"
	"strings"
)

// Функция для проверки уникальности символов
func UniqueSymbols(s string) bool {
	runes := []rune(strings.ToLower(s))
	mapa := make(map[rune]struct{}, len(runes))
	for _, v := range runes {
		if _, ok := mapa[v]; !ok {
			mapa[v] = struct{}{}
		} else {
			return false
		}
	}
	return true

}
func main() {
	strs := []string{"abcd", "abCdefAaf", "aabcd", "1234abcd", "11asbasd", "", " "}

	for _, v := range strs {
		if UniqueSymbols(v) {
			fmt.Printf("%s - unique\n", v)
		} else {
			fmt.Printf("%s - not unique\n", v)
		}
	}
}
