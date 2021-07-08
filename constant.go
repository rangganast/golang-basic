package main

import "fmt"

func main() {
	// error example
	// const firstName string

	const firstName string = "Rangga"
	const lastName = "Aziz"
	const number = 1000

	const (
		name  string = "Rangga"
		myNum        = 100000
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(number)
	fmt.Println(name)
	fmt.Println(myNum)
}
