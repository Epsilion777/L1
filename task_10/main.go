// Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0,
// 15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10
// градусов. Последовательность в подмножноствах не важна.
// Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.

package main

import (
	"fmt"
	"math"
)

func TemperatureRanges(temperatures []float64) map[int][]float64 {
	mapa := map[int][]float64{}

	for _, val := range temperatures {
		var group int

		if val > 0 {
			// Если температура положительна, то ее надо округлить к ближайшему меньшему
			group = int(math.Floor(val/10) * 10)

		} else {
			// Иначе надо округлить к ближайшему большему
			group = int(math.Ceil(val/10) * 10)
		}
		mapa[group] = append(mapa[group], val)
	}
	return mapa
}
func main() {
	tmp := []float64{24.5, -21.0, 32.5, 13.0, 19.0, 15.5, -25.4, -27.0, 0, -1.5, -124.8}
	mapa := TemperatureRanges(tmp)
	fmt.Println(mapa)
}
