package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64

	fmt.Println("Введите сторону 1-ую сторону прямоугольника:")
	fmt.Scanln(&a)

	fmt.Println("Введите сторону 2-ую сторону прямоугольника:")
	fmt.Scanln(&b)

	fmt.Println("Введите сторону 3-ую сторону прямоугольника:")
	fmt.Scanln(&c)

	p := (a + b + c) / 2 // полупериметр
	areaOfTriangle := math.Sqrt(p * (p - a) * (p - b) * (p - c))

	fmt.Println("Area of treangle equale:", areaOfTriangle)
}
