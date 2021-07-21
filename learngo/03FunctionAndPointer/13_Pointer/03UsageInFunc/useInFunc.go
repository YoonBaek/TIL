// Author : Github YoonBaek
package main

import "fmt"

// Function can return pointer.
func createPointer() *float64 {
	myFloat := 98.5
	return &myFloat
}

// Function also can get pointer as a parameter.
func printPointer(myBoolPointer *bool) {
	fmt.Println(*myBoolPointer)
}

func main() {
	//Go Programming can return local variable's pointer.
	var myFloatPointer *float64 = createPointer()
	fmt.Println(*myFloatPointer)

	myBool := true
	printPointer(&myBool)
}
