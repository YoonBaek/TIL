// Author : Github YoonBaek
// This program calculates average of numbers

package main

import "fmt"

func main() {
	numbers := [3]float64{71.8, 56.2, 89.5}
	sum := 0.0
	for _, number := range numbers {
		sum += number
	}
	fmt.Println("Sum :", sum)

	count := float64(len(numbers))
	fmt.Printf("Avg : %0.2f\n", sum/count)
}
