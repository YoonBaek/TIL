// Solved by Github Yoon Baek
// Q : https://www.acmicpc.net/problem/2438

package main

import (
	b "bufio"
	. "fmt"
	"os"
)

func main() {
	w := b.NewWriter(os.Stdout)

	var n, i int
	Scanln(&n)

	s := ""
	for i < n {
		i++
		s += "*"
		Println(s)
	}
	w.Flush()
}
