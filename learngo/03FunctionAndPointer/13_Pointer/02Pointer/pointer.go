// Author : Github YoonBaek

package main

import (
	"fmt"
	"reflect"
)

func main() {
	amount := 6
	fmt.Println(amount)  // Value of the variable
	fmt.Println(&amount) // Address of the variable

	// In Go, You can load every type of variable's address
	var (
		myInt   int
		myFloat float64
		myBool  bool
	)
	fmt.Println(&myInt, &myFloat, &myBool)
	// &Variable indicates the address value of the variable.
	// And this is Pointer

	// Type of Pointer
	// Pointer also has its own type
	fmt.Println(
		reflect.TypeOf(&myInt),
		reflect.TypeOf(&myFloat),
		reflect.TypeOf(&myBool),
	)

	// You can also declare pointer as a variable
	var myIntPointer *int
	myIntPointer = &myInt
	fmt.Println(myIntPointer)

	// Short Declaration is also possible
	myBoolPointer := &myBool
	fmt.Println(myBoolPointer)

	// See the value of the adress' variable.
	// We call this value at.
	fmt.Println(*myIntPointer, *myBoolPointer)

	// Change the value of the variable.
	*myIntPointer = 9
	fmt.Println(myInt) // myInt changed to 9
}
