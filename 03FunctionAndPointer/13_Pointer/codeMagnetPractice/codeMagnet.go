// Author : Github Yoon Baek
// This is my solution of Head First Go 106pg
// We need to print answer 42

package main

import "fmt"

func main() {
	var myInt int
	var myIntPointer *int
	myInt = 42
	myIntPointer = &myInt

	fmt.Println(*myIntPointer)
}
