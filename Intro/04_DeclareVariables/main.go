// Author : Github YoonBaek
// Golang is a static language.
// Unlike Python, You have to tell Go the type of the varibles
package main

import "fmt"

func main() {
	// Declare variables
	var quantity int
	var length, width float64 // you can declare multiple variables
	var customerName string

	/*
		default value
		int, float = 0
		string = ""
		bool = false
	*/

	// Assign values
	quantity = 4
	length, width = 1.2, 2.4
	customerName = "Yoon Baek"

	/*
		or you can declare and assign at the sane time like this
		var quantity int = 4
		var length, width float64 = 1.2, 2.4
		var customerName string = "Yoon Baek"

		or type of variable can be automatically declared
		var quantity = 4
		var length, width = 1.2, 2.4
		var customerName = "Yoon Baek"
	*/

	// Use Variables
	fmt.Println(customerName, "has ordered",
		quantity, "sheets each with an area of",
		length*width, "square meters.")

	// code magnet quiz
	codeMagnet(10, 4)

	// Declare variables in short way
	goDeveloper := "Gopher"
	year := 2021
	fmt.Println("Be A", goDeveloper, "in", year)

}

// there are two ways of declaring variables,
// but both ways are used frequently.
// Get used to it!
