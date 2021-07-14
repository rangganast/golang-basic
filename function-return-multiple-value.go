package main

import "fmt"

func getFullName() (string, string, string) {
	return "Rangga", "Aziz", "Nst"
}

func main() {
	firstName, _, lastName := getFullName()
	fmt.Println(firstName, lastName)
}
