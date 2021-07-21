// Learned by Github YoonBaek
// 가장 먼저 call된 스택이 가장 나중에 defer 끝남.

package main

import "fmt"

func main() {
	one()
}

func one() {
	defer fmt.Println("first defer")
	two()
}

func two() {
	defer fmt.Println("second defer")
	panic("let's see what's been defered")
}
