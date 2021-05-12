// Author : Github YoonBaek
// 구조체에 구독자 정보 저장하기

// 구조체 변수를 선언하고 구조체 필드에 값을 할당하는 방법을 배웠으니
// 이제 잡지의 구독자 정보를 저장할 구조체를 만들 수 있습니다.

package main

import "fmt"

func main() {
	// 먼저 subscriber 라는 이름의 변수를 정의합니다. 그 다음 subscriber에
	// 이름, 요금, 구독 여부 등을 저장합니다.
	var subscriber struct {
		name   string
		rate   float64
		active bool
	}
	subscriber.name = "Baek Gopher"
	subscriber.rate = 4.99
	subscriber.active = true
	fmt.Println("Name :", subscriber.name)
	fmt.Println("Rate :", subscriber.rate)
	fmt.Println("Active? :", subscriber.active)

	// 비록 구독자의 각 정보는 별도의 타입을 가진 필드에 저장되지만 구조체 덕분에 각 필드를 하나로 묶어 다룰 수 있습니다.
}
