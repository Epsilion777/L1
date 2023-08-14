// Разработать программу нахождения расстояния между двумя точками, которые
// представлены в виде структуры Point с инкапсулированными параметрами x,y и
// конструктором.
package main

import (
	"fmt"
	"math"
)

// Структура точки
type Point struct {
	x float64
	y float64
}

// Конструктор
func NewPoint(x, y float64) *Point {
	return &Point{
		x,
		y,
	}
}

// Функция для нахождения расстояния между точками
func Distance(a, b *Point) float64 {
	dx := math.Abs(b.x - a.x)
	dy := math.Abs(b.y - a.y)
	return math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2))
}
func main() {
	pointA := NewPoint(5, 1)
	pointB := NewPoint(0, -3.5)
	distance := Distance(pointA, pointB)
	fmt.Printf("pointA: (%0.2f; %0.2f)\npointB: (%0.2f; %0.2f)\nDistance: %0.2f", pointA.x, pointA.y, pointB.x, pointB.y, distance)
}
