package main

import "fmt"

func main() {
	type FullName string
	type Age rune

	var myFullName FullName = "M Rangga A N"
	var myAge Age = 24

	fmt.Println("My name is", myFullName)
	fmt.Println("My age is", myAge)
}
