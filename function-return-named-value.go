package main

import "fmt"

func getFullName2() (firstName string, middleName string, lastName string) {
	firstName = "Rangga"
	middleName = "Aziz"
	lastName = "Nst"

	return
}

func main() {
	a, b, c := getFullName2()
	fmt.Println(a, b, c)
}
