// Реализовать структуру-счетчик, которая будет инкрементироваться в
// конкурентной среде. По завершению программа должна выводить итоговое
// значение счетчика.
package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	count int
	sync.Mutex
}

func (c *Counter) Inc(wg *sync.WaitGroup) {
	defer wg.Done()
	c.Lock()
	c.count++
	c.Unlock()
}

func main() {
	wg := sync.WaitGroup{}
	counter := Counter{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go counter.Inc(&wg)
	}
	wg.Wait()
	fmt.Println(counter.count)
}
