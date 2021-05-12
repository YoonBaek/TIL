// Author : Github YoonBaek
// 도트 연산자를 통한 구조체 필드 접근
/* 아직은 구조체를 정의하기만하고 실제로 사용해보지는 않았습니다.
그럼 이제 구조체 필드에 값을 저장하고 가져오는 법을 배워보겠습니다.
우리는 지금까지 함수가 어떤 패키지에 속해 있는지 또는
메서드가 어떤 값에 속해있는지 나타내기 위해 도트 연산자를 사용해왔습니다. */
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("fmt패키지의 함수를 호출합니다.")
	var myTime time.Time
	// Time 값의 메서드를 호출합니다.
	fmt.Println(myTime.Year())

	// 마찬가지로 도트 연산자를 사용하면 구조체에 속해 있는 필드에 접근할 수 있으며
	// 필드에 값을 할당하거나 값을 가져올 수 있습니다.

	var myStruct struct {
		number float64
		word   string
		toggle bool
	}

	myStruct.number = 3.14
	myStruct.word = "pie"
	myStruct.toggle = true

	fmt.Printf("%#v\n", myStruct)
	//struct { number float64; word string; toggle bool }{number:3.14, word:"pie", toggle:true}
}
