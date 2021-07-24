// learned by Github YoonBaek
// 타입으로서의 함수

// 다른 함수를 인자로 호출할 때 아무 함수나 호출할 수 없습니다.
// 파라미터에서 지정된 유형의 함수만 호출할 수 있습니다.

package main

import (
	"fmt"
)

func sayHi() {
	fmt.Println("Hi")
}

func divide(a int, b int) float64 {
	return float64(a) / float64(b)
}

func multiply(a int, b int) int {
	return a * b
}

func doMath(a, b int, paramFunc func(int, int) float64) {
	fmt.Println(paramFunc(a, b))
}

func main() {
	// 파라미터와 리턴 값이 필요 없는 유형의 함수를 저장하는 변수를 생성합니다.
	var greetingFunc func()
	// 파라미터와 리턴 값을 가진 유형의 함수를 저장합니다.
	var mathFunc func(int, int) float64

	greetingFunc = sayHi
	mathFunc = divide

	greetingFunc()
	fmt.Println(mathFunc(3, 5))

	// 이렇듯 일급함수를 활용하기 위해서는
	// 파라미터와 리턴 값의 개수와 타입이 모두 일치해야합니다.

	// 이는 함수의 argument로 들어가는 함수에도 마찬가지로 적용됩니다.
	// 1
	doMath(1, 3, divide)
	// 2 일부러 에러가 나게 짠 코드입니다. 실행하진 마세요.
	doMath(1, 3, multiply)

	// 2에 빨간줄이 표시됨을 볼 수 있습니다.
	// 이는 2의 리턴 값으로 float64가 아닌 int를 받고 있기 때문에 불일치하기 때문에 나오는 현상입니다

	// 앞서 http.HandleFunc의 인자로 sayHi를 넣었을 때 작동하지 않는 이유도 위와 같습니다.
	// http.HandleFunc의 doc을 확인하면, 파라미터 함수로 http.ResponseWriter와 *Request를 받는 함수만을
	// 받을 수 있도록 지정되어 있는 것을 알 수 있습니다.
}
