package main

import "fmt"

func main() {
	// if with short statment
	var name = "Rangga"

	if length := len(name); length > 5 {
		fmt.Println("Name length is good")
	} else {
		fmt.Println("Name is too short")
	}
}
