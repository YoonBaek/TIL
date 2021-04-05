// Solved by Github Yoon Baek
// Q : https://www.acmicpc.net/problem/10818

package main

import (
	b "bufio"
	. "fmt"
	"os"
	"sort"
)

func main() {
	r := b.NewReader(os.Stdin)
	n := 0
	Fscanln(r, &n)
	in := make([]int, n)
	i := n
	for i > 0 {
		i--
		Fscanf(r, "%d ", &in[i])
	}
	sort.Slice(in, func(i, j int) bool { return in[i] < in[j] })
	Println(in[0], in[n-1])
}
