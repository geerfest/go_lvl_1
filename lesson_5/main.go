package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func fibonachi(myMap map[int]int, num int) int {
	if num == 0 {
		return 0
	}
	if num == 1 {
		return 1
	}

	value, numExists := myMap[num]
	if numExists {
		return value
	}

	myMap[num] = fibonachi(myMap, num-1) + fibonachi(myMap, num-2)

	return myMap[num]
}

func main() {
	myMap := map[int]int{}
	for {
		fmt.Println("Введите n-ое число для расчета Фибоначи")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			num, err := strconv.ParseInt(scanner.Text(), 10, 64)
			if err != nil {
				panic(err)
			}

			reslut := fibonachi(myMap, int(num))
			fmt.Println("Результат: ", reslut)
		}
	}
}
