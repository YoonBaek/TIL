// Learned by : github Yoon Baek
// 함수에서 사용자 정의 타입 사용하기
// 사용자 정의 타입은 변수의 타입 뿐만 아니라 함수의 매개변수와 반환값의 타입으로도 사용할 수 있습니다.

package main

import "fmt"

type part struct {
	description string
	count       int
}

// 매개변수로 생성한 타입을 넣을 수도 있고,
func showInfo(p part) {
	fmt.Println("Description :", p.description)
	fmt.Println("Count :", p.count)
}

// 생성한 타입을 반환하는 함수를 만들 수도 있습니다.
func makePart(desc string, cnt int) (p part) {
	p.description = desc
	p.count = cnt
	return p
}

// 그럼 이제 잡지의 subscriber 타입을 사용하는 몇 가지 함수를 사용하겠습니다.
type subscriber struct {
	name   string
	rate   float64
	active bool
}

func printSubInfo(s subscriber) {
	fmt.Println("Name :", s.name)
	fmt.Println("Rate :", s.rate)
	fmt.Println("Active? :", s.active)
}

func defaultSubscriber(name string) (s subscriber) {
	// name 외의 디폴트 값들을 지정해 줍니다.
	s.name = name
	s.rate = 5.99
	s.active = true

	return s
}

func main() {
	var bolts part

	bolts.description = "Hex bolts"
	bolts.count = 24
	showInfo(bolts)

	p := makePart("Screws", 8)
	showInfo(p)

	// 그럼 이제 잡지의 subscriber 타입을 사용하는 몇 가지 함수를 사용하겠습니다.
	subscriber1 := defaultSubscriber("Yoon")
	// 별도의 구독료를 책정이 필요하다면 지정해줍니다.
	subscriber1.rate = 4.99
	printSubInfo(subscriber1)

	subscriber2 := defaultSubscriber("Baek")
	printSubInfo(subscriber2)

	// 커맨드 창에서 두 값이 다름을 확인할 수 있습니다.
	// 변수명을 설정할 때 타입명과 섀도잉이 되지 않도록 주의하세요!
}
