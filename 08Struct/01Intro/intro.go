// Author : Github YoonBaek
// 슬라이스와 맵은 한가지 타입의 값만 갖습니다.
// 구독자 이름, 구독료, 구독 여부라는 세가지 타입의 값을 한 번에 저장할 수 있을까?

package main

import "fmt"

func main() {
	/*
		구조체란 여러 타입의 값으로 구성된 값입니다.
		슬라이스가 string 타입의 값만 저장할 수 있거나 맵이 int타입의 값만 저장할 수 있는 것과는 다르게
		구조체는 string, int, float64, bool 등 여러 타입의 값들을 한데 묶어 저장할 수 있습니다.
	*/

	// How to declare struct
	var myStruct struct {
		// you can declare multple fileds in struct
		number float64
		word   string
		toggle bool
	}

	// 구조체의 필드별로 각자 자신의 타입에 맞는 디폴트 값으로 초기화 되어 있음을 확인할 수 있습니다.
	fmt.Printf("%#v\n", myStruct)
}
