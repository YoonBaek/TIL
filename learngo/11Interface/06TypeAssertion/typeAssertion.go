// Learned by Github YoonBaek
// 타입 단언.
// 쉽게 말해 인터페이스로 정의한 타입이 사실 이거야. 라고 단언해주는 기능
package main

import (
	"fmt"
)

type Robot string

func (r Robot) MakeSound() {
	fmt.Println("Beep")
}

func (r Robot) Walk() {
	fmt.Println("Powering legs")
}

type NoiseMaker interface {
	MakeSound()
}

type TapeRecorder struct {
	Batteries   int
	Microphones int
}

type TapePlayer struct {
	Batteries int
}

func (t TapePlayer) Play(song string) {
	fmt.Println("Playing", song)
}

func (t TapePlayer) Stop() {
	fmt.Println("Stopped")
}

func (t TapeRecorder) Record() {
	fmt.Println("Recording...")
}

func (t TapeRecorder) Play(song string) {
	fmt.Println("Playing", song)
}

func (t TapeRecorder) Stop() {
	fmt.Println("Stopped")
}

type Player interface {
	Play(string)
	Stop()
}

func TryOut(p Player) {
	p.Play("test track")
	p.Stop()
	// panic 방지.
	recorder, ok := p.(TapeRecorder)
	if ok {
		recorder.Record()
	} else {
		fmt.Println("This is not a recorder.")
	}
}

func main() {
	// Robot 타입의 변수를 인터페이스 타입으로 생성합니다.
	var noisemaker NoiseMaker = Robot("Boston Dynamics Mk1")

	// How to Type Assertion
	// 쉽게 말해, 이 변수는 noisemaker 타입 중에서도 robot이라고 확신해 그러니까 써. 라고 단언
	var robot Robot = noisemaker.(Robot)

	// 타입 단언을 활용하면, interface에는 정의되어 있지 않지만 해당 타입에는 정의되어 있는 메서드를 호출할 수 있습니다.
	robot.Walk()

	// Type Assertion Failure
	TryOut(TapeRecorder{}) // ok
	// TryOut(TapePlayer{})   // failed. 단언하려는 타입과 일치하지 않으면 panic 발생

	// multiple return을 받으면 패닉이 발생하지 않습니다.
	var player Player = TapePlayer{}
	recorder, ok := player.(TapeRecorder)
	if ok {
		recorder.Record()
	} else {
		fmt.Println("Player was not a taperecorder")
	}
	// panic이 발생하지 않았씁니다.
	// 이렇게 단언하는 타입 값이 원래 무엇인지 확신하기 어려운 경우에는
	// 선택적 ok갑을 사용해 다른 타입을 가졌을 경우에 대한 대비책을 마련합니다.

	// TryOut에 반영시켜보기
	TryOut(TapePlayer{})
}
