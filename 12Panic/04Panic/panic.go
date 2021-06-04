// Learned by Github YoonBaek
// 패닉 호출하는 법.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	defer awardPrize()

	rand.Seed(time.Now().Unix())

	// 선언법 참 쉽다. 그런데 어려운게 또 패닉이다.
	panic("Oh My God")
}

//패닉 사용 예시
func awardPrize() {
	doorNumber := rand.Intn(3) + 1
	if doorNumber == 1 {
		fmt.Println("You win a cruise!")
	} else if doorNumber == 2 {
		fmt.Println("You win a car!")
	} else if doorNumber == 3 {
		fmt.Println("You win a boat!")
	} else {
		panic("invalid or number")
	}
}

// 만약 doorNumber에 1,2,3 이외의 숫자가 저장된다면 이는 사용자의 오류가 아닌
// 프로그램의 버그입니다. 따라서 doorNumber에 잘못된 값이 들어가는 경우에는
// panic을 호출하는 게 맞습니다. 이런 일은 절대 발생해서는 안 되며
// 만에 하나 발생한다면 동작을 잘못하기 전에 프로그램을 중단하는게 좋습니다.
