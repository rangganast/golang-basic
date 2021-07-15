package main

import "fmt"

func endApp() {
	fmt.Println("Aplikasi selesai")
	message := recover()
	if message != nil {
		fmt.Println("Error:", message)
	}
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("APP IS ERROR")
	}
	fmt.Println("App is running")
}

func main() {
	runApp(true)
	fmt.Println("Rangga")
}
