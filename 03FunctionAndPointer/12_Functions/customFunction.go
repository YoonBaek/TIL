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
	Printf("%.2f litters needed\n", area/10.0)
}

func main() {
	sayHello()
	repeatLine("line", 3)
	paintNeeded(3.2, 3.0)
	paintNeeded(5.2, 3.5)
	paintNeeded(5.0, 3.3)
}
