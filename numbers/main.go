package main

import "fmt"

func main() {
	var n int
	fmt.Println("Введите число до которого хотите найти все простые числа")
	fmt.Scanln(&n)
	if n <= 1 {
		fmt.Println("[]")
		return
	}

	allNumbersToN := make([]int, 0, n+1)

	// Запполняем слайс элементами
	for i := 0; i <= n; i++ {
		allNumbersToN = append(allNumbersToN, i)
	}

	allNumbersToN[1] = 0
	i := 2
	// Идем по слайсу все не простые числа заменяем 0
	for i < n {
		if allNumbersToN[i] != 0 {
			// Убираем все кратные i элементы
			j := i + i
			for j <= n {
				allNumbersToN[j] = 0
				j = j + i
			}
		}
		i = i + 1
	}

	// Заполняем новый слайс не нулевыми элементами
	simpleNumbers := make([]int, 0)
	for _, value := range allNumbersToN {
		if value != 0 {
			simpleNumbers = append(simpleNumbers, value)
		}
	}

	fmt.Println(simpleNumbers)

}
