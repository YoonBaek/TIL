// Learned by Github YoonBaek
// 빈 인터페이스
/*
go doc fmt Println의 결과입니다.
func Println(a ...interface{}) (n int, err error)
    Println formats using the default formats for its operands and writes to
    standard output. Spaces are always added between operands and a newline is
    appended. It returns the number of bytes written and any write error
    encountered.

인자로 interface{}를 받는 것을 볼 수 있습니다.
이는 빈 인터페이스를 뜻합니다.
*/
package main

import (
	"fmt"
	"math"

	"github.com/YoonBaek/learngo/11Interface/01WhyWeNeedInterface/gadget"
)

// 빈 인터페이스란 ?
// 아래와 같이 메서드가 전혀 필요 없는 인터페이스를 말합니다.
type EmptyInterface interface{}

// 빈 인터페이스는 메서드가 없는 만큼 모든 타입을 인자로 만족시킬 수 있습니다!
// 그래서 fmt.Println()의 경우도 모든 타입의 인자를 출력할 수 있는 것이죠.

func AcceptAnyting(anything interface{}) {
	fmt.Println(anything)
}

// 그러나 빈 인터페이스는 메서드가 없는 만큼 이것으로 할 수 있는게 그리 많지 않습니다.
func WannaDoSomething(something interface{}) {
	// But you can't
	// something.Do() // ./EmptyInterface.go:36:11: something.Do undefined (type interface {} is interface with no methods)
}

// 물론 타입 Assertion으로 가져올 수는 있습니다.
type Robot string

func (r Robot) MakeSound() {
	fmt.Println("Beep")
}

type NoiseMaker interface {
	MakeSound()
}

func ReallyWannaDoSth(something interface{}) {
	// Now you can
	beep, ok := something.(Robot)
	if ok {
		beep.MakeSound()
	}
}

// 그러나 이렇게 할 바에야 그냥 Specific type을 받는게 나을 수도 있겠죠.
func NowIcanDo(r Robot) {
	fmt.Println(r)
	r.MakeSound()
}

// 그래도 여러 타입을 한 함수에 굳이 받아야 할 때 써봄직은 합니다.

func main() {
	AcceptAnyting(math.Pi)
	AcceptAnyting(gadget.TapeRecorder{})
	ReallyWannaDoSth(Robot("GoLang Destroyer")) // Beep!
	NowIcanDo(Robot("GoLang Master"))
}
