// Author : Github YoonBaek
package main

import (
	"fmt"
	"log"
)

func paintNeeded(width float64, height float64) (float64, error) {
	// setting Errors
	if width < 0 {
		return 0, fmt.Errorf("A width of %0.2f is invalid.", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("A height of %0.2f is invalid.", height)
	}
	area := width * height
	// if there are no errors, return float value and nil.
	return area / 10.0, nil
}

func main() {
	amount, err := paintNeeded(4.2, -3.0)
	// We don't need 0.00 liters when error occured
	// So seperate the case using conditions
	if err != nil {
		fmt.Println(err)
	} else {
		// if error occuredm this code isn't gonna be executed.
		fmt.Println("0.2f liters needed\n", amount)
	}

	// Poolpuzzle
	quotient, err := divide(5.6, 0.0)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%0.2f\n", quotient)
	}
}
