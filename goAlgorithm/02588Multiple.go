// Solved by Github YoonBaek
// Q : https://www.acmicpc.net/problem/2588
package main

import (
	. "fmt"
)

func main() {
	a, b := 0, 0
	Scan(&a, &b)
	Println(a * (b % 10))
	Println(a * (b % 100 / 10))
	Println(a * (b / 100))
	Println(a * b)
}
