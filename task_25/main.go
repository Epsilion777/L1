// Реализовать собственную функцию sleep.
package main

import (
	"fmt"
	"time"
)

// Sleep with time.Since
func Sleep1(d time.Duration) {
	timeStart := time.Now()
	for {
		if time.Since(timeStart) >= d {
			return
		}
	}
}

// Sleep with channel
func Sleep2(d time.Duration) {
	<-time.After(d)
}

func main() {
	fmt.Println("Start Sleep1")
	timeNow := time.Now()
	Sleep1(500 * time.Millisecond)
	fmt.Println("Time has passed:", time.Since(timeNow))

	fmt.Println("Start Sleep2")
	timeNow = time.Now()
	Sleep2(2 * time.Second)
	fmt.Println("Time has passed:", time.Since(timeNow))
}
