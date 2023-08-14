// Реализовать конкурентную запись данных в map.
package main

import (
	"fmt"
	"sync"
)

type MySyncMap struct {
	mapa map[int]int
	sync.RWMutex
}

// Метод с использованием собственной MySyncMap
func MethodWithMutex() {
	m := MySyncMap{
		mapa: make(map[int]int),
	}
	wg := sync.WaitGroup{}
	wg.Add(50)
	for i := 0; i < 50; i++ {
		go func(i int) {
			defer wg.Done()
			// Блокирование разделяемого ресурса на чтение и запись
			m.Lock()
			m.mapa[i] = i * i
			m.Unlock()

			// Блокирование разделяемого ресурса для параллельного чтения
			m.RLock()
			fmt.Printf("map[%d]=%d\n", i, m.mapa[i])
			m.RUnlock()
		}(i)
	}
	wg.Wait()
}

// Метод с использованием sync.Map
func MethodWithSyncMap() {
	m := sync.Map{}
	wg := sync.WaitGroup{}
	wg.Add(50)
	for i := 0; i < 50; i++ {
		go func(i int) {
			defer wg.Done()
			m.Store(i, i*i)
			val, ok := m.Load(i)
			if ok {
				fmt.Printf("map[%d]=%d\n", i, val)
			}
		}(i)
	}
	wg.Wait()
}

func main() {
	fmt.Println("Solution with Mutex")
	MethodWithMutex()
	fmt.Println("Solution with SyncMap")
	MethodWithSyncMap()
}
