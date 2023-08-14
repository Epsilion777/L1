// Удалить i-ый элемент из слайса.
package main

import (
	"fmt"
	"log"
)

// Функция удаления элемента слайса с изменением исходного слайса
func DeleteFromSlice1(slc []int, index int) ([]int, error) {
	if index < 0 || index >= len(slc) {
		return nil, fmt.Errorf("index out of range. Index is %d with slice length %d", index, len(slc))
	}
	return append(slc[:index], slc[index+1:]...), nil
}

// Функция удаления элемента слайса без изменения исходного слайса
func DeleteFromSlice2(slc []int, index int) ([]int, error) {
	if index < 0 || index >= len(slc) {
		return nil, fmt.Errorf("index out of range. Index is %d with slice length %d", index, len(slc))
	}
	result := make([]int, 0, len(slc)-1)
	result = append(result, slc[:index]...)

	return append(result, slc[index+1:]...), nil
}

func main() {
	// Метод с изменением исходного слайса:
	fmt.Println("Method with changing of the original slice:")
	slc := []int{1, 2, 3, 4, 5, 6}
	newSLc, err := DeleteFromSlice1(slc, 1)
	if err != nil {
		log.Println("Error:", err)
		return
	}
	fmt.Println("Original slice:", slc)
	fmt.Println("New slice:", newSLc)

	// Метод без изменения исходного слайса:
	fmt.Println("Method without changing the original slice:")
	slc = []int{1, 2, 3, 4, 5, 6}
	newSLc, err = DeleteFromSlice2(slc, 3)
	if err != nil {
		log.Println("Error:", err)
		return
	}
	fmt.Println("Original slice:", slc)
	fmt.Println("New slice:", newSLc)

}
