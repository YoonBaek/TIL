// Solved by Github Yoon Baek
// Q : https://www.acmicpc.net/problem/10951

package main

import . "fmt"

func main() {
	var a, n int
	Scanln(&a)
	r := a
	for true {
		f := r % 10
		b := (r/10 + r%10) % 10
		r = 10*f + b
		n++
		if r == a {
			Println(n)
			break
		}
	}
}
