// Author : Github YoonBaek
// fmt 패키지 함수는 배열을 처리할 수 있습니다.

package main

import "fmt"

func main() {
	notes := [3]string{"do", "re", "mi"}
	primes := [5]int{2, 3, 5, 7, 11}

	// returns array itself
	fmt.Println(notes)
	fmt.Println(primes)

	// returns just like code
	fmt.Printf("%#v\n", notes)
	fmt.Printf("%#v\n", primes)
}
