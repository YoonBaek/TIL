// Author : Github Yoon Baek
// 루프 내에서 배열 원소 접근 하기

package main

import "fmt"

func main() {
	notes := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}

	for i := 0; i <= 2; i++ {
		fmt.Println(i, notes[i])
	}

	// take a good care of index.
	// If it exceeds range of index, 'panic' will be occurred.
	// 패닉이란 프로그램이 실행되는 동안 발생하는 에러를 말합니다.(컴파일 에러와는 다름)

	// 에러를 보고 싶으면 아래 3줄의 주석을 해제하세요.
	// for i := 0; i <= 7; i++ {
	// 	fmt.Println(i, notes[i])
	// } // panic: runtime error: index out of range [7] with length 7

	// If panic occurred, program stops.
	fmt.Println("If panic occurred, this will not be printed.")

	// Instead, use len() to figure out the size of the array.
	// Caution! num of loop should not be same as the size of the array.
	for i := 0; i < len(notes); i++ {
		fmt.Println(i, notes[i])
	}
}
