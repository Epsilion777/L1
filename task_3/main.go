// Дана последовательность чисел: 2,4,6,8,10.
// Найти сумму их квадратов(2^2+3^2+4^2...) с использованием конкурентных вычислений.
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func square(num int) int {
	return num * num
}

// Решение с WaitGroup и Mutex
func FirstSolution(arr []int) {
	wg := sync.WaitGroup{}
	m := sync.Mutex{}
	sum := 0

	// Добавляем 5 элементов в очередь
	wg.Add(len(arr))
	for _, v := range arr {
		// Прокидываем значение в функцию
		go func(v int) {
			defer wg.Done()
			// Блокируем доступ к разделяемому ресурсу
			m.Lock()
			sum += square(v)
			m.Unlock()
		}(v)
	}

	wg.Wait()
	fmt.Println(sum)
}

// Решение с channel
func SecondSolution(arr []int) {
	ch := make(chan int, len(arr))
	sum := 0
	// Прокидываем значение в функцию
	for _, v := range arr {
		go func(v int) {
			ch <- square(v)
		}(v)
	}
	// Считываем значения из канала и суммируем их
	for i := 0; i < len(arr); i++ {
		sum += <-ch
	}
	close(ch)
	fmt.Println(sum)
}

// Решение с WaitGroup и Atomic
func ThirdSolution(arr []int) {
	wg := sync.WaitGroup{}
	var sum int64
	pSum := &sum

	// Добавляем 5 элементов в очередь
	wg.Add(len(arr))
	for _, v := range arr {
		// Прокидываем значение в функцию
		go func(v int) {
			defer wg.Done()
			// Атомарно накапливам сумму с помощью указателя pSum
			atomic.AddInt64(pSum, int64(square(v)))
		}(v)
	}

	wg.Wait()
	fmt.Println(sum)
}

func main() {
	arr := [...]int{2, 4, 6, 8, 10}
	fmt.Println("Solution with WaitGroup and Mutex:")
	FirstSolution(arr[:])
	fmt.Println("Solution with Channel:")
	SecondSolution(arr[:])
	fmt.Println("Solution with WaitGroup and Atomic:")
	ThirdSolution(arr[:])
}
