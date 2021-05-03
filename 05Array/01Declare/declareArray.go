// Author : Github YoonBaek
package main

import (
	"fmt"
	"time"
)

func main() {
	// How to declare Array
	// [num of arr]types of elements

	// declare array with 7 strings
	var notes [7]string
	notes[0] = "do"
	notes[1] = "re"
	notes[2] = "mi"

	fmt.Println(notes[0])
	fmt.Println(notes[1])

	// decalre array with 5 integers
	var primes [5]int

	primes[0] = 2
	primes[1] = 3
	fmt.Println(primes[0])

	// declare array with 3 time.Times
	var dates [3]time.Time
	dates[0] = time.Unix(1257894000, 0)
	dates[1] = time.Unix(1447920000, 0)
	dates[2] = time.Unix(1508632200, 0)
	fmt.Println(dates[2])
}
