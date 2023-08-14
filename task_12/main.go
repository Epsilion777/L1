// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее
// собственное множество.

package main

import "fmt"

type Set struct {
	data map[string]struct{}
}

// Метод для создания множества
func NewSet(arrStr []string) *Set {
	mapa := make(map[string]struct{}, len(arrStr))
	for _, v := range arrStr {
		mapa[v] = struct{}{}
	}
	return &Set{
		data: mapa,
	}
}

// Метод для добавления элемента во множество
func (s *Set) Add(value string) {
	s.data[value] = struct{}{}
}

// Метод для проверки наличия элемента во множестве
func (s *Set) Contains(value string) bool {
	_, exists := s.data[value]
	return exists
}

// Метод для удаления элемента из множества
func (s *Set) Delete(value string) {
	delete(s.data, value)
}

// Метод для вывода множества
func (s *Set) Print() {
	for key := range s.data {
		fmt.Printf("%s ", key)
	}
}

func main() {
	arrStr := []string{"cat", "cat", "dog", "cat", "tree"}
	setStr := NewSet(arrStr)

	setStr.Print()
}
