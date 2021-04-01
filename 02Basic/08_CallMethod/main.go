// Author : Github YoonBaek
/*
This looks quite similar to package function call,
But both are different.

package function call format :
packagename.Function()

method call format :
value.method()
*/

package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	// method of time.Time
	var now time.Time = time.Now() // time.Now returns time.Time value, which represents current datetime
	var year int = now.Year()      // Year is a method of time.Time
	fmt.Println("Now :", now, "\nYear :", year)

	// method of strings.NewReplacer
	var brokenLetter string = "G# R#cks!"
	var replacer *strings.Replacer = strings.NewReplacer("#", "o")

	var fixedLetter string = replacer.Replace(brokenLetter)
	fmt.Println(brokenLetter, "->", fixedLetter)
}
