package main

import "fmt"

func main() {
	// variable declaration without data initialization
	var name string

	name = "Rangga"
	fmt.Println(name)

	name = "Rangga Aziz"
	fmt.Println(name)

	// error example
	// name = true

	// variable declaration with data initialization
	var friendName = "Dudung"
	fmt.Println(friendName)

	// variable declaration with type
	var myNumber int8 = 10
	fmt.Println(myNumber)

	// variable declaration with :=
	myOtherNumber := 20
	fmt.Println(myOtherNumber)

	// Multiple declarations
	var (
		firstName = "Rangga"
		lastName  = "Aziz"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
