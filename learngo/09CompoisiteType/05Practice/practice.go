// Solved by Github Yoon Baek

package main

import "fmt"

type Number int

func (n Number) Add(otherNumber int) int {
	return int(n) + otherNumber
}
func (n Number) Subtract(otherNumber int) int {
	return int(n) - otherNumber
}

func main() {
	ten := Number(10)
	fmt.Println(ten.Add(4))
	fmt.Println(ten.Subtract(5))
	four := Number(4)
	fmt.Println(four.Add(3))
	fmt.Println(four.Subtract(2))
}
