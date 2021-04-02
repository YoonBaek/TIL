// Solved by Github YoonBaek
// Q : https://www.acmicpc.net/problem/2753

package main

import . "fmt"

func main() {
	var year, leap int
	Scan(&year)
	if year%4 == 0 {
		leap = 1
	}
	if year%100 == 0 {
		leap = 0
	}
	if year%400 == 0 {
		leap = 1
	}
	Println(leap)
}
