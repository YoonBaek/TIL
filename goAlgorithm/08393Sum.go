// Solved by Github Yoon Baek
// Q : https://www.acmicpc.net/problem/8393

package main

import . "fmt"

func main() {
	var n, sum int
	Scan(&n)
	for n > 0 {
		sum += n
		n--
	}
	Println(sum)
}
