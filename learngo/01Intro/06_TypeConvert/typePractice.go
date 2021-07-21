// Author : Github Yoon Baek

package main

import "fmt"

// Practice : make this work.
/*
	var price int = 100
	fmt.Println("Price is", price, "dollars.")

	var taxRate float64 = 0.08
	var tax float64 = price * taxRate // invalid.
	fmt.Println("Tax is", tax, "dollars.")

	var total float64 = price + tax
	fmt.Println("Total cost is", total, "dollars.")

	var availableFunds int = 120
	fmt.Println(availableFunds, "dollars are available.")
	fmt.Println("Within budget?", total <= availableFunds)
*/
func typePractice() {
	var price int = 100
	fmt.Println("Price is", price, "dollars.")

	priceF64 := float64(price)

	var taxRate float64 = 0.08
	var tax float64 = priceF64 * taxRate // invalid.
	fmt.Println("Tax is", tax, "dollars.")

	var total float64 = priceF64 + tax
	fmt.Println("Total cost is", total, "dollars.")

	var availableFunds int = 120
	fmt.Println(availableFunds, "dollars are available.")
	fmt.Println("Within budget?", total <= float64(availableFunds))
}
