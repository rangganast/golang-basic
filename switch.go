package main

import "fmt"

func main() {
	switch number := 1; number {
	case 1:
		fmt.Println("This is one")
	case 2:
		fmt.Println("This is two")
	default:
		fmt.Println("This is a number")
	}

	// switch without condition
	var name = "Rangga"
	switch length := len(name); {
	case length > 10:
		fmt.Println("Name is too long")
	case length > 5:
		fmt.Println("Name quite long")
	default:
		fmt.Println("Name is good")
	}
}
