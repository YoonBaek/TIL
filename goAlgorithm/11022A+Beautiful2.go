// Solved by Github Yoon Baek
// Q : https://www.acmicpc.net/problem/11022

package main

import (
	b "bufio"
	. "fmt"
	"os"
)

func main() {
	r := b.NewReader(os.Stdin)
	w := b.NewWriter(os.Stdout)

	var t, a, b, i int
	Fscanln(r, &t)

	for i < t {
		i++
		Fscanln(r, &a, &b)
		Fprintf(w, "Case #%d: %d + %d = %d\n", i, a, b, a+b)
	}
	w.Flush()
}
