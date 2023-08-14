// Реализовать паттерн «адаптер» на любом примере.

package main

import "fmt"

/*
Представим ситуацию, что мы были фанатами телефонов Samsung и всю жизнь пользовались только ими.
В какой-то прекрасный день нам захотелось преобрести iPhone. Купив его, мы обнаруживаем, что провод Ligtning оказался бракованным.
Отдав кругленькую сумму за iPhone не очень хочется идти и докупать новый оригинальный Ligtning, который стоит немалых денег.
Хорошо бы иметь переходник с Type-C на Ligtning, ведь проводов типа Type-C у нас скопилось огромное количество.
*/

// Провод Lightning
type Lightning struct {
}

// Метод, который вставляет Lightning в iPhone
func (l Lightning) InsertLightning() {
	fmt.Println("Lightning is inserted and works correctly")
}

// Провод TypeC
type TypeC struct {
}

// Человек, который будет заряжать iPhone
type Person struct {
}

// Метод, который позволяет человеку начать заряжать iPhone
func (p Person) ChargeiPhone(u Lightninger) {
	u.InsertLightning()
	fmt.Println("The iPhone is charging")
}

// Интерфейс, который определяет, что может делать Lightning
type Lightninger interface {
	InsertLightning()
}

// Адаптер, который используется как переходник с Type-C на Lightning
type AdapterTypecToLightning struct {
	Adapter *TypeC
}

// Реализуем у адаптера метод, чтобы он имплементировал интерфейс UseLightning
func (a AdapterTypecToLightning) InsertLightning() {
	fmt.Println("TypeC with Adapter to Lightning is inserted and works correctly")
}

func main() {
	p := Person{}
	// Заряжаем iPhone с помощью оригинального Lightning
	originalLightning := Lightning{}
	p.ChargeiPhone(originalLightning)
	fmt.Println("---------------------------------------------------------------")
	// Заряжаем iPhone с помощью Type-C и переходника (адаптера) под Lightning
	originalTypeC := TypeC{}
	adapter := AdapterTypecToLightning{&originalTypeC}
	p.ChargeiPhone(adapter)
}
