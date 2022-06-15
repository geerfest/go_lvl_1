package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Введите по-очередно 10 чисел")
	countElements := 10

	a := []int64{}
	scanner := bufio.NewScanner(os.Stdin)
	for n := countElements; n >= 0; n-- {
		if scanner.Scan() {
			num, err := strconv.ParseInt(scanner.Text(), 10, 64)
			if err != nil {
				panic(err)
			}
			a = append(a, num)
		}
	}
	fmt.Println("Изначальная последовательность:\n", a)

	for i := 0; i < len(a)-1; i++ {
		j := i + 1
		if a[j] > a[i] {
			continue
		}
		a[j], a[i] = a[i], a[j]
		for i := i; i > 0; i-- {
			bi := i - 1
			if a[i] < a[bi] {
				a[i], a[bi] = a[bi], a[i]
			}
		}
	}

	fmt.Println("Отсортированная последовательность:\n", a)
}
