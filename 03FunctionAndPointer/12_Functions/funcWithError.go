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
	Println(err.Error())
}
