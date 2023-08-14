// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из
// массива, во второй — результат операции x*2, после чего данные из второго
// канала должны выводиться в stdout.
package main

import "fmt"

// Функция записывает данные в канал
func Writer(ch chan<- int, arr []int) {
	defer close(ch)

	// Запишись значения каждого среза в канал
	for _, v := range arr {
		ch <- v
	}
}

// Функция считывает данные из ch1 и записывает в ch2
func Worker(ch1 <-chan int, ch2 chan<- int) {
	defer close(ch2)
	// Считываем каждое значение ch1 и отправляем значение * 2 в ch2
	for val := range ch1 {
		ch2 <- val * 2
	}
}
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	arr := [...]int{1, 5, -1, 2, 0, 0, -15, 22, 47, 1024, 58, 11, 0, 9}

	go Writer(ch1, arr[:])
	go Worker(ch1, ch2)

	// Выводим результат
	for val := range ch2 {
		fmt.Println(val)
	}
}
