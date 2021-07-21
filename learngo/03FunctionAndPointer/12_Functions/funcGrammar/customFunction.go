// Author : Github Yoon Baek
package main

import (
	. "fmt"
)

/*
Naming function also follows camelCase conventions.
Every rule is same as naming variable rules. export, starting with letters... etc.
You can't call functions in main package like this. main.sayHello() ...
*/

// Declare Simple function
func sayHello() {
	Println("Hello")
}

// Declare function with parameters
// Paramter is a local Variable of the function, declared as the function declared.

// func FUNCNAME(PARAM1NAME PARAM1FORMAT, PARAM2NAME PARAM2FORMAT ...) {}
func repeatLine(line string, times int) {
	for i := 0; i < times; i++ {
		Println(line)
	}
}

// Use custom function in calculating paint
func paintNeeded(width, height float64) {
	area := width * height
	Printf("%.2f litters needed\n", area/outsideFunc)
}

// function also has a scope.
var outsideFunc float64 = 10

// make function which returns value, not print
// set the type of return value right before the block.
func double(number float64) float64 {
	return number * 2
}

// paintNeeded with Return
// return type must be declared.
func paintNeededReturn(width float64, height float64) float64 {
	area := width * height
	return area / 10.0 // return func should include and end with return.
	// and return type must match declared return type.
}

func main() {
	sayHello()
	repeatLine("line", 3)
	paintNeeded(3.2, 3.0) // 0.96
	paintNeeded(5.2, 3.5)
	paintNeeded(5.0, 3.3)

	// function also has a scope.
	// fmt.Println(area) // area is inside the paintNeeded func. but not defined outsidefunc.
	outsideFunc = 5       // We can change the value because we are in a scope of outsideFunc var
	paintNeeded(3.2, 3.0) // 1.92

	// we can assign return as a value of a variable,
	// or pass it to other function as well.
	dozen := double(6)
	Println(dozen)
	Println(double(3.2))

	// Use func return in Paint Calculation
	var amount, total float64 // now you may know default is 0.0
	amount = paintNeededReturn(4.2, 3.0)
	Printf("%.3f liters needed.\n", amount)

	// add to total.
	total += amount

	amount = paintNeededReturn(5.2, 3.5)
	Printf("%.3f liters needed.\n", amount)

	total += amount
	Printf("Total : %.3f liters total.", total)
}

// Next time, we'll figure out how to return errors and multiple values
