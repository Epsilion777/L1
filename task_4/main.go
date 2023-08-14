// Реализовать постоянную запись данных в канал (главный поток).
// Реализовать набор из N воркеров, которые читают произвольные данные из канала
// и выводят в stdout. Необходима возможность выбора количества воркеров при старте.
// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения
// работы всех воркеров.
package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"sync"
)

// Worker для обработки данных из канала
func Worker(ctx context.Context, wg *sync.WaitGroup, ch <-chan int, num int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Thread %d has ended\n", num)
			return
		case val := <-ch:
			fmt.Printf("Thread %d recived: %d\n", num, val)
		}
	}
}

func main() {
	maxNumCPU := runtime.NumCPU()
	N := 0
	ch := make(chan int)
	// Используем (maxNumCPU - 1), потому что 1 поток это main gourutine
	fmt.Printf("Enter a number of workers (max - %d):", maxNumCPU-1)
	fmt.Fscan(os.Stdin, &N)

	if N < 1 {
		log.Fatalln("The number of workers cannot be < 1")
	}
	// Если количество worker больше максимального, то мы устанавливаем значение равным (maxNumCPU - 1), чтобы выполнять параллельное выполнение на всех ядрах
	if N > maxNumCPU-1 {
		N = maxNumCPU - 1
	}

	// Создание контекста, который будет отменен при получении прерывания (ctrl+c)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	wg := sync.WaitGroup{}

	for i := 0; i < N; i++ {
		wg.Add(1)
		go Worker(ctx, &wg, ch, i+1)
	}
	loop := true
	for loop {
		select {
		case <-ctx.Done():
			fmt.Println("The program ended on interruption")
			loop = false
		case ch <- rand.Intn(100):
		}
	}
	// Ждем когда все рабочие закончат свою работу
	wg.Wait()
	// Закрываем канал после завершения всех worker
	close(ch)
}
