// Реализовать бинарный поиск встроенными методами языка
package main

import "fmt"

// Бинарный поиск (возвращает true и индекс в случае удачного поиска и false -1 в случае неудачи)
func BinarySerch(slc []int, value int) (exists bool, index int) {
	left := 0
	right := len(slc) - 1

	// Пока левый индекс не сравнялся с правым
	for left <= right {
		center := (right + left) / 2
		if slc[center] == value {
			return true, center
		}
		// Меняем правую границу на центр
		if slc[center] > value {
			// Меняем правую границу на центр
			right = center - 1
		} else {
			// Меняем левую границу на центр
			left = center + 1
		}
	}

	return false, -1
}

func main() {
	slc := []int{-15, -8, -1, 0, 5, 10, 20, 20, 21, 100, 112}
	fmt.Println(BinarySerch(slc, 10))
}
