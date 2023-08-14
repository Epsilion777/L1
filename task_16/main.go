// Реализовать быструю сортировку массива (quicksort) встроенными методами
// языка.

package main

import "fmt"

// Функция быстрой сортировки
func QuickSort(slc []int, low, high int) {
	if low < high {
		pivot := slc[high] // Опорный элемент (последний элемент массива)
		i := low

		for j := low; j < high; j++ {
			if slc[j] < pivot {
				slc[i], slc[j] = slc[j], slc[i] // Меняем элементы местами (все элементы меньше опорного двигаем влево от опорного)
				i++
			}
		}

		// Меняем опорный элемент и элемент на позиции i местами
		slc[i], slc[high] = slc[high], slc[i]

		// Вызываем быструю сортировку для левого подмассива
		QuickSort(slc, low, i-1)
		// Вызываем быструю сортировку для правого подмассива
		QuickSort(slc, i+1, high)
	}
}

func main() {
	slc := []int{10, 7, 8, 9, 1, 5}
	fmt.Println("Original slice:", slc)

	QuickSort(slc, 0, len(slc)-1)
	fmt.Println("Sorted slice:", slc)
}
