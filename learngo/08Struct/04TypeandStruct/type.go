// Author : Github Yoon Baek
// 사용자 정의 타입과 구조체
// 구조체가 여러 타입을 담을 수 있다는 건 좋지만,
// 이때까지 해온 것으로는 새로운 변수를 선언할 때마다 구조체 정보를 계속 정의해줘야한다는 문제가 있습니다.

package main

import "fmt"

type subscriber struct {
	name   string
	rate   float64
	active bool
}

func main() {
	var sub1 struct {
		name   string
		rate   float64
		active bool
	}
	sub1.name = "Baek Gopher"
	fmt.Println(sub1.name)
	var sub2 struct {
		name   string
		rate   float64
		active bool
	}
	sub2.name = "Baek Golfer"
	fmt.Println(sub2.name)
	// 이것처럼요.
	// 이제 자료구조를 사용만 하지말고 기본 자료구조르 기반으로
	// 사용자 정의 타입을 만들어봅시다.

	// 타입은 보통 패키지 수준에서 선언하는게 일반적이나,
	// 가독성을 위해 main에서 정의
	type parts struct {
		description string
		count       int
	}

	type car struct {
		name     string
		topSpeed float64
	}

	// 이제 변수 선언 시 긴 구조체 정의를 일일이 작성할 필요 없이 사용자 정의 타입의 이름만 사용하면 됩니다.
	var porsche car
	porsche.name = "Porsche Carrera 4s"
	porsche.topSpeed = 323
	fmt.Println("Name :", porsche.name)
	fmt.Println("TopSpeed :", porsche.topSpeed)

	var bolts parts
	bolts.description = "Hex bolts"
	bolts.count = 24

	fmt.Println("Description :", bolts.description)
	fmt.Println("Count :", bolts.count)

	// 잡지 구독자 정보에 사용자 정의 타입 사용하기

	var subs1, subs2 subscriber

	subs1.name = "Baek"
	subs2.name = "Yoon"
	fmt.Println("Name :", subs1.name)
	fmt.Println("Name :", subs2.name)
}
