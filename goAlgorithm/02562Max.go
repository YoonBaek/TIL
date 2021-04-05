// Solved by Github Yoon Baek
// Q : https://www.acmicpc.net/problem/2562

package main

import (
	. "fmt"
)

func main() {
	var max, n, ix int
	for i := 0; i < 9; i++ {
		Scanln(&n)
		if n > max {
			max = n
			ix = i + 1
		}
	}
	Printf("%d\n%d\n", max, ix)
}
