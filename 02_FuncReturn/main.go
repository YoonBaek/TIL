// Author : Github YoonBaek
// If you want to see the result of the function,
// You have to print it.

package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	// this returns nothing
	math.Floor(2.75)
	strings.Title("this returns nothing.") // Make first letter uppercase
	// now you have return
	fmt.Println(math.Floor(2.75))
	fmt.Println(strings.Title("now this returns."))
}

// if you want to compare, make line 16-17 comments.
