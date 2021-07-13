package main

import "fmt"

func main() {

	var names [3]string

	names[0] = "Muhammad"
	names[1] = "Rangga"
	names[2] = "Aziz"
	// error example
	// names[3] = "Nasution"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	// fmt.Println(names[3])

	// direct initialization
	var numbers = []int{
		90,
		85,
		80,
	}

	fmt.Println(numbers)
	fmt.Println(numbers[0])
	fmt.Println(numbers[1])
	fmt.Println(numbers[2])

	fmt.Println(len(names))
	fmt.Println(len(numbers))

	var newArray [10]string

	fmt.Println(len(newArray))
}
