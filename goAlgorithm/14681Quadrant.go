// Solved by Github Yoon Baek
// Q : https://www.acmicpc.net/problem/14681

package main

import . "fmt"

func main() {
	var x, y, q int
	Scan(&x, &y)
	if x > 0 && y > 0 {
		q = 1
	} else if x < 0 && y > 0 {
		q = 2
	} else if x < 0 && y < 0 {
		q = 3
	} else {
		q = 4
	}
	Println(q)
}
