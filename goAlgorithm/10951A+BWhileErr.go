// Solved by Github Yoon Baek
// Q : https://www.acmicpc.net/problem/10951

package main

import . "fmt"

func main() {
	var a, b int
	for true {
		_, err := Scanln(&a, &b)
		if err != nil {
			break
		}
		Println(a + b)
	}
}
