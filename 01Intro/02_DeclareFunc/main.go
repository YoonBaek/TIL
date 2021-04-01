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
	// see the answer of poolpuzzle quiz
	// you can ignore this.
	// but in same package, you can import without the uppercase
	poolpuzzle()
}

// if you want to compare, make line 18-19 comments.
