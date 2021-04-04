// Solved by Github Yoon Baek
// Q : https://www.acmicpc.net/problem/2739

package main

import . "fmt"

func main() {
	var n int
	Scan(&n)
	for i := 1; i <= 9; i++ {
		Printf("%d * %d = %d\n", n, i, n*i)
	}
}
