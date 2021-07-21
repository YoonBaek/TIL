// Author Github Yoon Baek
// Amount like height, width, area can't be negative
// So we're going to figure out how to handle these exceptions.
package main

import (
	"errors"
	. "fmt"
)

func main() {
	err := errors.New("height can't be negative")
	Println(err.Error()) // Println(err) is also possible due to builtin function.

	// Error with value
	errf := Errorf("A height of %.2f is invalid.", -2.0)
	// this returns error, not print error
	Println(errf) // same as Println(errf.Error())
}
