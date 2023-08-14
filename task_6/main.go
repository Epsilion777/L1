// Реализовать все возможные способы остановки выполнения горутины.
package main

import (
	"context"
	"fmt"
	"time"
)

// Метод с использованием context
func CancelWithContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("WorkerContext has finished working")
			return
		default:
			fmt.Println("WorkerContext is working")
		}
	}
}

// Метод с использованием channel
func CancelWithChannel(exit chan int) {
	for {
		select {
		case <-exit:
			fmt.Println("WorkerChannel has finished working")
			return
		default:
			fmt.Println("WorkerChannel is working")
		}
	}
}

// Метод с использованием закрытия главной goroutine
func CancelWithCancelMainRoutine() {
	for {
		fmt.Println("The worker works while main goroutine is alive")
	}
}

func main() {

	// Остановка с помощью cotnext
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	// Запуск программы, которая завершит работу через 2 секунды
	go CancelWithContext(ctx)

	time.Sleep(time.Second * 3)

	// Остановка с закрытым каналом
	exit := make(chan int)

	// Запуск goroutine, которая завершит работу при закрытии канала
	go CancelWithChannel(exit)

	time.Sleep(time.Second * 3)
	close(exit)
	time.Sleep(time.Second * 1)

	// Запуск goroutine, которая завершит работу после завершения главной goroutine
	go CancelWithCancelMainRoutine()
	time.Sleep(time.Second * 1)
}
