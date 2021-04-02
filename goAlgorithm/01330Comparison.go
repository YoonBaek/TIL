// Solved by Github Yoon Baek
// Q : https://www.acmicpc.net/problem/1330

package main

import . "fmt"

func main() {
	var a, b int
	Scan(&a, &b)
	if a > b {
		Println(">")
	} else if a < b {
		Println("<")
	} else {
		Println("==")
	}
}
