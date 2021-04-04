// Solved by Github Yoon Baek
// Q : https://www.acmicpc.net/problem/10950

package main

import . "fmt"

func main() {
	var t, a, b int
	Scan(&t)
	for i := 0; i < t; i++ {
		Scan(&a, &b)
		Println(a + b)
	}
}
