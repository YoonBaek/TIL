// Solved by Github Yoon Baek
// Q : https://www.acmicpc.net/problem/15552

package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var t, a, b int
	Fscanln(r, &t)

	for t > 0 {
		Fscanln(r, &a, &b)
		Fprintln(w, a+b)
		t--
	}

	w.Flush()
}
