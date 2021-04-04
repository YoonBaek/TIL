// Solved by Github Yoon Baek
// Q : https://www.acmicpc.net/problem/11021

package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var t, a, b, i int
	Fscanln(r, &t)

	for i < t {
		i++
		Fscanln(r, &a, &b)
		Fprintf(w, "Case #%d: %d\n", i, a+b)
	}

	w.Flush()
}
