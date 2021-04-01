// Author : Github Yoon Baek
// You can convert types using internal function.
// But you should assign float value to float variables.
// for example,
// var four int = 4
// four = float64(four)
// is not permitted.

package main

import (
	"fmt"
	"reflect"
)

func main() {
	length := 1.2 // float
	length = 2
	fmt.Println(reflect.TypeOf(length)) // automatically compile as float
	width := 2                          // int

	// If you want to see the error, remove comments
	// fmt.Println("Area is", length*width)           // operation between different types is invalid.
	// fmt.Println("length > width?", length > width) // too.

	fmt.Println("Area is", length*float64(width)) // valid.

	/*
		And also,
		float32 and float64 : invalid.
	*/

	// Becareful!
	// If you convert float to int, you lose the value under the dot.
	var near4 = 4 - 1e-10
	fmt.Println("Near 4 should be 4, but result is", int(near4))

	// typePractice
	// refer typePractice.go
	typePractice()
}
