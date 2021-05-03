// Author : Github YoonBaek
// 배열이 가질 값을 미리 알고 있는 경우에는 배열 리터럴을 사용하여 배열을 초기화할 수 있습니다.

package main

import "fmt"

func main() {
	// how to declare array with literals
	// [num of elems]type of elems{literals...}
	notes := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}

	for i := 0; i < len(notes); i++ {
		fmt.Printf("%s ", notes[i])
	}
	fmt.Println()

	// same with integer arrays.
	primes := [7]int{2, 3, 5, 7, 11, 13}
	for i := 0; i < len(primes); i++ {
		fmt.Printf("%d ", primes[i])
	}
	fmt.Println()

	// Also, you can change the line while declare literals.
	texts := [3]string{
		"Let's study golang.",
		"I'll do my own golang project till this year.",
		"Go together!", // should end with comma
	}

	for i := 0; i < len(texts); i++ {
		fmt.Printf("%s ", texts[i])
	}
	fmt.Println()
}
