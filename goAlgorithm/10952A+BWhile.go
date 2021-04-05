// Solved by Github Yoon Baek
// Q : https://www.acmicpc.net/problem/10952

package main

import . "fmt"

func main() {
	var a, b int
	for true {
		Scanln(&a, &b)
		if a == 0 && b == 0 {
			break
		}
		Println(a + b)
	}
}
