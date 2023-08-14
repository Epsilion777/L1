// Написать программу, которая конкурентно рассчитает значение квадратов чисел
// взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

package main

import (
	"fmt"
	"strings"
	"sync"
)

func square(num int) int {
	return num * num
}

// Решение с WaitGroup
func FirstSolution(arr []int) {
	result := strings.Builder{}
	m := sync.Mutex{}
	wg := sync.WaitGroup{}

	// Добавляем 5 элементов в очередь
	wg.Add(len(arr))
	for _, v := range arr {
		// Прокидываем значение в функцию
		go func(v int) {
			defer wg.Done()
			// Блокируем доступ к разделяемому ресурсу
			m.Lock()
			result.WriteString(fmt.Sprintf("%d ", square(v)))
			m.Unlock()
		}(v)
	}
	wg.Wait()
	fmt.Println(result.String())
}

// Решение с Channel
func SecondSolution(arr []int) {
	result := strings.Builder{}

	ch := make(chan int, len(arr))
	for _, v := range arr {
		// Прокидываем значение в функцию
		go func(v int) {
			ch <- square(v)
		}(v)
	}
	// Считываем значения из канала и записываем в результат
	for i := 0; i < len(arr); i++ {
		result.WriteString(fmt.Sprintf("%d ", <-ch))
	}
	close(ch)
	fmt.Println(result.String())
}

func main() {
	arr := [...]int{2, 4, 6, 8, 10}
	fmt.Println("Solution with WaitGroup:")
	FirstSolution(arr[:])
	fmt.Println("Solution with Channel:")
	SecondSolution(arr[:])
}
