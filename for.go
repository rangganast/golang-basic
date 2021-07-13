package main

import "fmt"

func main() {
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("The", counter, "loop")
	}

	slice := []string{"Rangga", "Aziz", "Nst"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// for range
	for _, value := range slice {
		// fmt.Println("index", i, "=", value)
		fmt.Println(value)
	}

	person := make(map[string]string)
	person["name"] = "Rangga"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println("key", key, "=", value)
	}
}
