// Solved by Github Yoon Baek
// Q : https://www.acmicpc.net/problem/3052
package main

import (
	b "bufio"
	. "fmt"
	"os"
)

func search(s []int, i int) bool {
	for _, e := range s {
		if e == i {
			return true
			break
		}
	}
	return false
}

func main() {
	r := b.NewReader(os.Stdin)
	var a, mod int
	var mods = []int{}

	for i := 0; i < 10; i++ {
		Fscanf(r, "%d\n", &a)
		mod = a % 42
		if !search(mods, mod) {
			mods = append(mods, mod)
		}
	}
	Println(len(mods))
}
