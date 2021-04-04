// Solved by Github Yoon Baek
// Q : https://www.acmicpc.net/problem/2439

package main

import (
	. "fmt"
)

func main() {
	var n int
	Scanln(&n)
	for i := 1; i <= n; i++ {
		s := ""
		j, k := n-i, i

		for j > 0 {
			s += " "
			j--
		}
		for k > 0 {
			s += "*"
			k--
		}
		Println(s)
	}
}
