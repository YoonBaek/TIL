// Solved by Github Yoon Baek
// Q : https://www.acmicpc.net/problem/10871

package main

import (
	b "bufio"
	. "fmt"
	"os"
)

func main() {
	r, w := b.NewReader(os.Stdin), b.NewWriter(os.Stdout)
	var a, i, n, x int
	Fscanln(r, &n, &x)
	for i < n {
		Fscanf(r, "%d ", &a)
		if a < x {
			Fprintf(w, "%d ", a)
		}
		i++
	}
	Fprintln(w)
	w.Flush()
}
