// Author : Github YoonBaek
// Why do we need pointer?
// here's the reason.

package main

import "fmt"

func double(number int) {
	number *= 2 // double the value of number
	fmt.Println(number)
}

func main() {
	amount := 6
	double(amount)      // 12
	fmt.Println(amount) // 6
}

/*
We can't change the value of amount
Because double() changed the copy of amount inside the function.
So we need to learn how to transfer the original value of amount,
and that is POINTER
*/
