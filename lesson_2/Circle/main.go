package main

import (
	"fmt"
	"math"
)

func main() {
	var s float64

	fmt.Println("Введите площадь круга:")
	fmt.Scanln(&s)

	d := math.Sqrt(s * 4 / math.Pi)
	fmt.Println("Диаметр окружности:", d)

	p := math.Pi * d
	fmt.Println("Длина окружности:", p)
}
