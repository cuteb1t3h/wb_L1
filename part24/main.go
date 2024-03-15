package main

import (
	"fmt"
	"math"
)

/* Приватные (достуные в рамках одного пакета) параметры объявляются с маленькой буквы.
Создается два новых объекта структуры Point, с помощью конструктора структуры.
Рассчитывается расстояние между ними по формуле d = sqrt(x2−x1)^2 + (y2−y1)^2
*/

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{
		x,
		y,
	}
}

func (p1 *Point) distance(p2 *Point) float64 {
	dx := p2.x - p1.x
	dy := p2.y - p1.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	point1 := NewPoint(2, 4)
	point2 := NewPoint(8, 12)

	dist := point2.distance(point1)

	fmt.Println(dist)
}
