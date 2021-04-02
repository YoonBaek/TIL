// Solved by Github Yoon Baek
// Q : https://www.acmicpc.net/problem/2884

package main

import . "fmt"

func main() {
	var h, m, e int
	Scan(&h, &m)
	e = 45
	if m >= e {
		m -= e
	} else {
		if h == 0 {
			h += 24
		}
		h--
		m += 15
	}
	Println(h, m)
}
