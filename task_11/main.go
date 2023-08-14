// Реализовать пересечение двух неупорядоченных множеств.

package main

import "fmt"

// Функция для нахождения пересечения двух неупорядоченных множеств
func IntersectionOfSets(arr1, arr2 []int) []int {
	mapa := make(map[int]struct{}, len(arr1))

	for _, v := range arr1 {
		mapa[v] = struct{}{}
	}
	// Результирующий слайс
	result := make([]int, 0, len(mapa))

	for _, v := range arr2 {
		if _, ok := mapa[v]; ok {
			result = append(result, v)
		}
	}
	return result
}
func main() {
	arr1 := []int{-15, 0, -3, 12, 6, 5, -100, 1000, 18, 115, 300, 14, -1, 8}
	arr2 := []int{10, 0, -8, 7, 13, 15, 144, -200, -1, 18, 114, 14, 6, -11, 128}
	fmt.Println(IntersectionOfSets(arr1, arr2))
}
