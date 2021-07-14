package main

import "fmt"

func introduceSelf(name string) string {
	return "Hello, my name is " + name
}

func main() {
	introduce := introduceSelf
	fmt.Println(introduce("Rangga"))
}
