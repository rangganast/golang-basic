package main

import "fmt"

func main() {
	name := "Rangga"

	greet := func() {
		name = "Budi"
		fmt.Println("Hello", name)
		// myNumber := 123
	}

	// cannot be accessed
	// fmt.Println(myNumber)

	greet()

	fmt.Println(name)
}
