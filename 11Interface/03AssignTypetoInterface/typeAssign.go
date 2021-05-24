package main

import "fmt"

type whistle string

type horn string

// 참고로, 인터페이스 내 메서드의 이름, 매개변수 및 리턴 값의 타입이나 개수가 일치해야 합니다.
type NoiseMaker interface {
	MakeSound()
}

func (w whistle) MakeSound() {
	fmt.Println("tweet!")
}

func (h horn) MakeSound() {
	fmt.Println("honk!")
}

// 함수의 매개변수를 선언할 때도 인터페이스 타입을 선언할 수 있습니다.(결국 함수의 매개변수도 변수이므로)
// 예를 들어, NoiseMaker타입의 값을 매개변수로 받는 play 함수를 선언하면 MakeSound 메서드를 가진 모든 타입의 값을 play함수로 전달할 수 있습니다.
func play(n NoiseMaker) {
	n.MakeSound()
}

// 다만 인터페이스는 인터페이스 내의 정의 된 메서드만 호출할 수 있습니다.
type Robot string

func (r Robot) MakeSound() {
	fmt.Println("Beep!")
}

func (r Robot) Walk() {
	fmt.Println("Powering legs...")
}

func main() {
	var toy NoiseMaker

	toy = whistle("true bird")
	toy.MakeSound()
	toy = horn("car")
	toy.MakeSound()

	// 함수 매개변수 테스트
	play(whistle("true bird"))
	play(horn("car"))
	// 이렇게 타입이 다르더라도 할당이 됩니다!

	toy = Robot("Taekwon V")
	play(toy)
	// walk는 NoiseMaker 인터페이스에 없기 때문에 불러올 수 없습니다.
	// toy.Walk() // 불가능
	// 즉 인터페이스에 없는 메서드를 불러오지만 않는다면, 할당하는 데에는 아무 문제가 없습니다.
}
