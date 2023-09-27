package main

import (
	"fmt"
	"math"
)

/*
Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
*/

// AB = √((xb - xa)2 + (yb - ya)2).
func main() {
	a := NewPoint(2, 2)
	b := NewPoint(5, 5)
	result := distanse(a, b)
	fmt.Println(result)
}
func distanse(a, b *Point) float64 {
	first := b.x - a.x
	second := b.y - a.y
	ab := math.Sqrt((math.Pow(float64(first), 2)) + (math.Pow(float64(second), 2)))
	return ab
}

type Point struct {
	x int
	y int
}

func NewPoint(x, y int) *Point {
	return &Point{
		x: x,
		y: y,
	}
}
