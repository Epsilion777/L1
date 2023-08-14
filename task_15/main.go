/*
К каким негативным последствиям может привести данный фрагмент кода, и как
это исправить? Приведите корректный пример реализации.
var justString string
func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}
func main() {
	someFunc()
}
*/

package main

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
)

// 1. v[:100] вернет 100 первых байт строки, а не 100 первых символов
// 2. Если по какой-то причине v будет меньше 100 байт, то вызовется panic: out of range [:100]
// 3. Использование глобальной переменной здесь излишне, лучше завести локальную переменную
func createHugeString(size int) string {
	sb := strings.Builder{}
	for i := 0; i < size; i++ {
		_, err := sb.WriteString(fmt.Sprintf("%v", rand.Intn(10)))
		if err != nil {
			log.Printf("Error: %v", err)
		}
	}
	return sb.String()
}

func someFunc() {
	var justString string
	v := createHugeString(1 << 10)
	runes := []rune(v)
	if len(runes) >= 100 {
		justString = string(runes[:100])
	} else {
		justString = v
	}
	fmt.Println(justString)
}
func main() {
	someFunc()
}
