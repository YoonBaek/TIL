package main

import "fmt"

func codeMagnet(originalCount int, eatenCount int) {
	var apple = "apple"
	if originalCount > 1 {
		apple = "apples"
	}
	fmt.Println("I started with", originalCount, apple)
	if eatenCount > 1 {
		apple = "apples"
	}
	fmt.Println("Some jerk ate", eatenCount, apple)
	if originalCount-eatenCount > 1 {
		apple = "apples"
	}
	fmt.Println("There are", originalCount-eatenCount, apple, "left")
}
