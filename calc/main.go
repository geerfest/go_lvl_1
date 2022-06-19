package main

import (
	"fmt"
	"math"
	"os"
)

// Проверям деление на 0
func isDivisionByZero(op string, b float32) bool {
	return op == "/" && b == 0
}

// Проверяем является ли число натуральным
func isNumberNuture(number float32) bool {
	numberAsInt := int(number)

	return number == float32(numberAsInt)
}

func factorial(number float32) float32 {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)
}

func calcFactorial(number float32) {
	if !isNumberNuture(number) {
		fmt.Printf("Число: %f является не натуральным\n", number)
		os.Exit(1)
	}

	res := factorial(number)
	fmt.Printf("Результат выполнения операции: %2.f\n", res)
	os.Exit(1)
}

func main() {
	var a, b, res float32
	var op string

	fmt.Print("Введите первое число: ")
	fmt.Scanln(&a)
	fmt.Print("Введите арифметическую операцию (+, -, *, /, !, ^): ")
	fmt.Scanln(&op)
	if op == "!" {
		calcFactorial(a)
	}

	fmt.Print("Введите второе число: ")
	fmt.Scanln(&b)

	if isDivisionByZero(op, b) {
		fmt.Println("На ноль делить нельзя")
		os.Exit(1)
	}

	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res = a / b
	case "^":
		res = float32(math.Pow(float64(a), float64(b)))
	default:
		fmt.Println("Опреация выбрана неверно")
		os.Exit(1)
	}

	fmt.Printf("Результат выполнения операции: %.2f\n", res)
}
