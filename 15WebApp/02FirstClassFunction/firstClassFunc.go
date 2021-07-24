// Learned by Githiub Yoon Baek
// 일급 함수로 다른 함수에 함수 전달하기

package main

import (
	"fmt"
)

func sayHi() {
	fmt.Println("hi!")
}

func sayBye() {
	fmt.Println("bye!")
}

// 함수를 파라미터로 줄 수 있습니다.
func twice(Function func()) {
	// 파라미터로 받은 함수를 두 번 호출합니다.
	Function()
	Function()
}

func main() {
	// 일급함수는 함수 호출이 아닌, 함수 그 자체로 전달합니다.
	// 메모리 상에 올라가 있는 변수라고 생각하면 되는듯...
	twice(sayHi)
	twice(sayBye)

	// 아래 코드는 타입 에러가 납니다.
	// http.HandleFunc("/hello", twice)

	// 이는 타입으로서의 함수 때문입니다.
}
