// Solved by Github Yoon Baek
// Q : https://www.acmicpc.net/problem/2577
package main

import (
	b "bufio"
	. "fmt"
	"os"
	"strconv"
	s "strings"
)

func main() {
	r := b.NewReader(os.Stdin)
	var a, b, c int
	Fscanf(r, "%d\n%d\n%d\n", &a, &b, &c)
	str := strconv.Itoa(a * b * c)
	nums := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for _, num := range nums {
		Println(s.Count(str, num))
	}
}
