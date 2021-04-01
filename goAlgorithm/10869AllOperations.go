// Solved by Github Yoon Baek
// Q : https://www.acmicpc.net/problem/10869
package main

import . "fmt"

func main() {
	a, b := 0, 0
	Scan(&a, &b)
	Printf("%d\n%d\n%d\n%d\n%d\n", a+b, a-b, a*b, a/b, a%b)
}
