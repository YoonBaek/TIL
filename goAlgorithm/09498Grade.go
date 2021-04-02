// Solved by Github Yoon Baek
// Q : https://www.acmicpc.net/problem/9498

package main

import . "fmt"

func main() {
	var grade int
	score := ""
	Scan(&grade)
	if grade >= 90 {
		score = "A"
	} else if grade >= 80 {
		score = "B"
	} else if grade >= 70 {
		score = "C"
	} else if grade >= 60 {
		score = "D"
	} else {
		score = "F"
	}
	Println(score)
}
