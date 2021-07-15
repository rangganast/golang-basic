package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) string {
	nameFiltered := filter(name)
	return "Hello " + nameFiltered
}

func spamFilter(name string) string {
	if name == "fuck" {
		return "*****"
	} else {
		return name
	}
}

func main() {
	filter := spamFilter

	fmt.Println(sayHelloWithFilter("Rangga", spamFilter))
	fmt.Println(sayHelloWithFilter("fuck", filter))
}
