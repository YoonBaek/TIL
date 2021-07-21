// Author Github Yoon Baek
package main

import (
	. "fmt"
	"math"
)

// You can declare multiple returns by putting types in second ()
func manyReturns() (int, bool, string) {
	return 1, true, "hello"
}

// You can also assign names of each return.
// This is mainly for documentation for other developers
func sepFloats(num float64) (integer int, restFloat float64) {
	wholeNumber := math.Floor(num)
	return int(wholeNumber), num - wholeNumber
}

func main() {
	myInt, myBool, myString := manyReturns()
	Println(myInt, myBool, myString)

	dollars, cents := sepFloats(1.25)
	Println(dollars, "and", cents)
}
