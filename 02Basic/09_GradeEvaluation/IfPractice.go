// Author : Github Yoon Baek
// Try this! and get used to golang if expressions.

package main

import "fmt"

func ifPractice() {
	if true {
		fmt.Println("true")
	}
	if false {
		fmt.Println("false")
	}
	if !false {
		fmt.Println("!false")
	}
	if true {
		fmt.Println("It's true!")
	} else {
		fmt.Println("else")
	}
	if false {
		fmt.Println("if false")
	} else if true {
		fmt.Println("else if true")
	}
	if 12 == 12 {
		fmt.Println("12 == 12")
	} else {
		fmt.Println("12 != 12")
	}
	if 12 == 12 && 5.9 == 5.9 {
		fmt.Println("12 == 12 && 5.9 == 5.9")
	}
}
