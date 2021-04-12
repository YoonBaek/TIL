// Author Github YoonBaek
// Let's Fix double function in 01Why folder

package main

import "fmt"

func double(number *int) {
	*number *= 2
}

func main() {
	amount := 6
	double(&amount)
	fmt.Println(amount) // 12. It changed!
}

/*
By Using Pointer, var amount inside the function is no more a copy,
but same as var amount outside the function
*/
