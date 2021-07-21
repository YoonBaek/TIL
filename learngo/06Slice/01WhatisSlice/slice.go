// Author : Gihub YoonBaek
// Go에는 값을 추가하여 확장할 수 있는 데이터 구조가 존재하는데 이를 슬라이스라고 합니다.
/*
배열과 마찬가지로 슬라이스도 복수 개의 원소로 이루어지며 모든 원소는 동일한 타입을 갖습니다.
그리고 배열과 달리 슬라이스의 끝에 원소를 추가할 수 있는 함수를 사용할 수 있습니다.
슬라이스 타입의 변수르 선언할 때에는 빈 대괄호 [] 다음에 원소의 타입을 지정하면 됩니다.
*/

package main

import "fmt"

func main() {
	// How to declare slice.
	var mySlice []string
	// compare to array
	var myArray [5]string

	// 배열 변수와는 달리 슬라이스 변수는 자동으로 슬라이스를 생성하지는 않습니다.
	fmt.Println("mySlice :", mySlice, "\nmyArray :", myArray)

	// 따라서 make 내장 함수를 사용해 슬라이스를 명시적으로 생성해 줘야 합니다.
	// make 함수에는 생성하려는 슬라이스 값의 타입
	// (생성된 값을 할당할 슬라이스 변수와 동일한 타입만 가능합니다)과 생성할 슬라이스의 크기를 전달합니다.
	notes := make([]string, 7)

	// 슬라이스가 생성되면 배열에서와 동일한 방식으로 값을 할당하고 가져올 수 있습니다.

	notes[0] = "do"
	notes[1] = "re"
	notes[2] = "mi"
	fmt.Println(notes[0], notes[1])

	// 단축 선언 역시 가능하며, len에 슬라이스를 전달하면
	// 배열과 동일하게 슬라이스의 길이를 반환합니다.
	primes := make([]int, 5)
	fmt.Println(len(notes), len(primes))

	// 그리고 for와 for ...range 루프 또한 배열에서와 동일하게 작동합니다.
	letters := []string{"a", "b", "c"}
	for i := 0; i < len(letters); i++ {
		fmt.Println(letters[i])
	}
	for _, letter := range letters {
		fmt.Println(letter)
	}
}
