package main

import "fmt"

func main() {
	var n int

	fmt.Println("Введите трехзначное число:")
	fmt.Scan(&n)

	sotki := n / 100
	desiatki := n % 100 / 10
	edinicy := n % 10

	fmt.Println("Сотен:", sotki, "\nДесятков:", desiatki, "\nЕдиниц:", edinicy)
}
