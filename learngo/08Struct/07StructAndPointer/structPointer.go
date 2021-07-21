// Learned by Github Yoon Baek
// 함수에서 구조체 변경하기.

// Go는 pass by value언어이기 때문에,
// 함수 내부의 변수는 복사본에 지나지 않는다는 것을 변수에서 배웠습니다.
// struct 역시 마찬가지 입니다.
package main

import "fmt"

func changeInt(a int) {
	a = 2
}

func changeInt2(a *int) {
	// 포인터 주소에 있는 값을 변경합니다.
	*a = 2
}

type subscriber struct {
	name   string
	rate   float64
	active bool
}

func discount(s subscriber) {
	s.rate *= 0.8
}

func applyDiscount(s *subscriber) {
	// 도트 표기법은 struct 포인터에서도 작동하기 때문에
	// 따로 값을 *로 넘겨주지 않아도 됩니다!
	s.rate *= 0.8
	// 이는 사실
	// *(s).rate *= 0.8 과 동일한 표기법이지만
	// golang 내부적으로 이런 번거로움을 피할 수 있게
	// struct 포인터에서도 필드에 대한 접근을 허용합니다.
}

func main() {
	// 변수도 별도 포인터를 할당해 주지 않으면 바뀌지 않습니다.
	a := 1
	changeInt(a)
	fmt.Println(a)
	// 포인터의 주소를 매개변수로 넘겨줍니다.
	changeInt2(&a)
	fmt.Println(a)
	// 값이 바뀌었습니다.

	// struct도 똑같습니다.
	var s subscriber
	s.rate = 5.99
	discount(s)
	// 여전히 5.99로 변하지 않습니다.
	fmt.Println(s.rate)

	// struct의 포인터 주소값을 넘겨줍니다.
	applyDiscount(&s)
	// 20프로가 할인되었습니다.
	fmt.Println(s.rate)
}
