// Solved by Github Yoon Baek
// Q : https://www.acmicpc.net/problem/10430
package main

import . "fmt"

func main() {
	var a, b, c int
	Scan(&a, &b, &c)
	Println((a + b) % c)
	Println(((a % c) + (b % c)) % c)
	Println((a * b) % c)
	Println(((a % c) * (b % c)) % c)
}
