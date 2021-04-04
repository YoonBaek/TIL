// Solved by Github Yoon Baek
// Q : https://www.acmicpc.net/problem/2742

package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	var n int
	Fscanln(r, &n)
	for n > 0 {
		Fprintln(w, n)
		n--
	}
	w.Flush()
}
