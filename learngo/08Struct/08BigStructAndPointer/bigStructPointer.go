// Learned by Github Yoon Baek
// 포인터를 사용한 큰 구조체 전달
// 큰 구조체를 매개변수로 쓰면 복사본 때문에 메모리를 많이 잡아먹는다.

package main

import "fmt"

type subscriber struct {
	name   string
	rate   float64
	active bool
}

// 포인터를 전달하게 되면 복사본이 생성되지 않기 때문에
// struct를 변경할 필요가 없는 경우라고 해도 이 방식이 더 좋습니다.
func printInfo(s *subscriber) {
	fmt.Println("Name :", s.name)
	fmt.Println("Monthly rate :", s.rate)
	fmt.Println("Active? :", s.active)
}

func defaultSubscriber(name string) *subscriber {
	var s subscriber
	s.name = name
	s.rate = 5.99
	s.active = true
	return &s
}

func applyDiscount(s *subscriber) {
	s.rate = 4.99
}

func main() {
	subscriber1 := defaultSubscriber("Gopher Baek")
	// 이미 defaultSubscriber는 Struct의 포인터 주소값을 반환하므로
	// 주소연산자 &를 사용할 필요가 없습니다.
	applyDiscount(subscriber1)
	printInfo(subscriber1)

	subscriber2 := defaultSubscriber("Gopher Yoon")
	printInfo(subscriber2)
}
