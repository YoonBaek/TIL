// Solved by Github YoonBaek

package main

import "fmt"

func odd(intChan chan int) {
	intChan <- 1
	intChan <- 3
}

func even(intChan chan int) {
	intChan <- 2
	intChan <- 4
}

func main() {
	channelOdd := make(chan int)
	channelEven := make(chan int)

	go odd(channelOdd)
	go even(channelEven)

	fmt.Println(<-channelOdd)
	fmt.Println(<-channelOdd)

	fmt.Println(<-channelEven)
	fmt.Println(<-channelEven)
}
