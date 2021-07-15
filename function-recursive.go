package main

import "fmt"

func sumSeriesLoop(num int) int {
	result := 0
	for num != 0 {
		result += num
		num--
	}
	return result
}

func sumSeriesRecursive(num int) int {
	if num == 0 {
		return num
	} else {
		return num + sumSeriesRecursive(num-1)
	}
}

func main() {
	fmt.Println(sumSeriesLoop(5))
	fmt.Println(sumSeriesRecursive(5))
}
