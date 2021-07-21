// Learned by Github YoonBaek
// go의 에러는 인터페이스 입니다!

package main

import (
	"fmt"
	"log"
)

// "에러 값은 문자열을 반환한하는 Error라는 메서드를 가진 값이다."
// 이는 인터페이스의 특징과도 일치함을 볼 수 있습니다.
// 네, error는 빌트인 인터페이스 입니다

type ComedyError string

func (c ComedyError) Error() string {
	return string(c)
}

// 과열을 추적하는 에러를 만들어봅시다.
type OverHeatError float64

func (o OverHeatError) Error() string {
	return fmt.Sprintf("Overheating by %0.2f degrees!", o)
}

func checkTemperature(actual float64, safe float64) error {
	excess := actual - safe
	// 초과 온도가 0을 넘으면...
	if excess > 0 {
		return OverHeatError(excess)
	}
	return nil
}

func main() {
	// 이렇게 에러 타입에 변수로 바로 할당할 수 있습니다.
	// Error()를 메서드로 갖고 있기 때문이죠. 인터페이스의 특징입니다.
	var err error = ComedyError("What's a programmer's favorite beer? Logger!") // 개노잼...

	fmt.Println(err)

	err = checkTemperature(121.379, 100.0)
	if err != nil {
		log.Fatal(err)
	}
}
