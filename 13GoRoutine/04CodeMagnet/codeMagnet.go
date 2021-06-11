// Solved by Github YoonBaek
package main

import (
	"fmt"
	"time"
)

func repeat(word string, count int) {
	for i := 0; i < count; i++ {
		fmt.Print(word)
	}
}

func main() {
	go repeat("x", 25)
	go repeat("y", 25)
	time.Sleep(time.Second)
	fmt.Print("\n")
}
