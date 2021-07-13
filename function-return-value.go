package main

import "fmt"

func sum(number1 int, number2 int) int {
	return number1 + number2
}

func main() {
	result := sum(1, 1)
	fmt.Println(result)

	fmt.Println(sum(3, 2))
}
