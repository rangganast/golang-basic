package main

import "fmt"

func main() {
	myNumber := map[int]string{
		1: "1",
		2: "2",
		// "3": "3", ERROR EXAMPLE
	}

	fmt.Println(myNumber)
	fmt.Println(myNumber[1])
	fmt.Println(myNumber[2])

	// add new map element
	myNumber[3] = "3"
	fmt.Println(myNumber)

	var vehicleWheels map[string]int = make(map[string]int)
	vehicleWheels["bicycle"] = 2
	vehicleWheels["car"] = 4
	vehicleWheels["motorbike"] = 3

	fmt.Println(vehicleWheels)
	fmt.Println(len(vehicleWheels))

	delete(vehicleWheels, "motorbike")

	fmt.Println(vehicleWheels)
	fmt.Println(len(vehicleWheels))
}
