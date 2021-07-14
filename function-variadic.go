package main

import "fmt"

func sumAll(firstNumber int, secondNumber int, restNumbers ...int) int {
	total := 0
	total += firstNumber
	total += secondNumber

	for _, value := range restNumbers {
		total += value
	}

	return total
}

func main() {
	fmt.Println(sumAll(10, 20, 10, 10, 10, 10, 10))

	// varargs not inserted
	fmt.Println(sumAll(10, 20))

	// use slice as varargs
	numbers := []int{10, 10, 10, 10, 10}
	fmt.Println(sumAll(10, 40, numbers...))
}
