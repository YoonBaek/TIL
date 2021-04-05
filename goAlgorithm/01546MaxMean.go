// Solved by Github Yoon Baek
// Q : https://www.acmicpc.net/problem/1546
package main

import (
	b "bufio"
	. "fmt"
	"os"
)

func main() {
	r := b.NewReader(os.Stdin)
	var n, score, max, sum int
	Fscanln(r, &n)
	scores := make([]int, n)
	for i := 0; i < n; i++ {
		Fscanf(r, "%d ", &score)
		scores[i] = score
	}
	for _, score := range scores {
		sum += score
		if score >= max {
			max = score
		}
	}
	Println(float64(sum) / float64(max*n) * 100)
}
