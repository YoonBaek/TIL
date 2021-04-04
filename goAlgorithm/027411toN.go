// Solved by Github Yoon Baek
// Q : https://www.acmicpc.net/problem/2741

package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	var n, i int
	Fscanln(r, &n)
	for n > i {
		i++
		Fprintln(w, i)
	}
	w.Flush()
}
