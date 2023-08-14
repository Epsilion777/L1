// Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Воркер который читает данные из канала
func Worker(ctx context.Context, wg *sync.WaitGroup, ch <-chan int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Thread has finished working ")
			return
		case val := <-ch:
			fmt.Println("Thread recived:", val)
		}
	}
}

func main() {
	N := 1

	// Создание контекста, который завершится через N секунд (WithTimeout)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(N))
	defer cancel()

	// // Создание контекста, который завершится через N секунд (WithDeadline)
	// ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*time.Duration(N)))
	// defer cancel()

	ch := make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go Worker(ctx, &wg, ch)

	loop := true
	for loop {
		select {
		case <-ctx.Done():
			fmt.Println("The program ended on time")
			loop = false
		case ch <- rand.Intn(100):
		}
	}

	// Ожидаем когда worker закончит свою работу
	wg.Wait()
	// Закрываем канал после завершения работы worker
	close(ch)
}
