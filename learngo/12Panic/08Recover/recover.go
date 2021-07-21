// Learned by Github YoonBaek
// scanDirectory 함수가 에러 대신 패닉을 사용하도록 코드를 수정하면 에러처리 코드가 아주 간단해집니다.
// 하지만 패닉은 복잡한 스택 트레이스와 함께 프로그램을 중단시킵니다.
// 사용자에게는 그저 간단한 에러메시지를 보여주는게 더 나을텐데 말이죠.

// Go에는 패닉상태에 빠진 프로그램을 복구할 수 있는 recover라는 내장함수가 있습니다.
// 패닉을 발생시키는 프로그램을 우아하게 종료하려면 recover를 사용해야합니다.
// recover를 정상적인 프로그램 실행 중에 호출하면 nil만  반환하고 아무 일도 수행하지 않습니다.
package main

import "fmt"

// 프로그램이 panic상태일 때 recover를 호출하면 패닉을 멈출 수 있습니다.
// 하지만 함수에서 panic을 호출하면 함수는 실행을 즉시 중단합니다.
// 따라서 패닉 상태가 계속 이어지기 떄문에 panic을 호출하는 함수와 동일한
// 함수에서  recover를 호출하는 것은 의미가 없습니다.
// 하지만 프로그램이 패닉상태일 때 recover를 호출하는 방법이 있습니다.
// 패닉이 발생하면 모든 지연된 함수 호출이 먼저 실행되기 때문에 recover를
// 별도의 함수에서 호출하도록 만든 뒤 패닉을 발생시키는 코드 앞에서
// 해당 함수의 호출을 지연시키면 됩니다.
func calmDown() {
	// 패닉이 발생한 경우, recover는 panic에 전달 된 모든 값을 반환합니다
	// 따라서 반환값을 사용해 패닉에 대한 정보를 가져올 수 있으며
	// 이는 복구작업을 수행하거나 사용자에게
	// 에러메시지를 보고할 때 유용하게 쓸 수 있습니다.

	// 다만 panic의 인자가  interface{}이듯,
	// recover의 반환값도 interface{}입니다.
	// 따라서 panic에 인자를 error 값으로 전달하더라도
	// interface{} 타입의 값이 나오므로 메서드를 호출할 수 없습니다.
	// 이 때 타입 단언을 활용해 줍니다.
	p := recover()
	err, ok := p.(error)

	if ok {
		fmt.Println(err.Error())
	}
}

// panic 실행부분 보다 먼저 앞서서 실행되는 +1 depth에서 recover가
// 실행되어야 합니다. defer자리에 recover() 사용불가.
// 무조건 panic앞에서 쓰여야합니다. 그러려면 recover()를 먼저 실행될
// 내부함수자리에 넣어주면 됩니다.
func freakOut() {
	defer calmDown()
	panic("omg")
	// recover를 내부에서 정상적으로 호출했다고 해서
	// 코드 실행이 재개되는 것은 아닙니다.
	// 패닉이 발생한 함수는 그 즉시 종료되며
	// 함수 내의 패닉 뒤에 위치한 모든 코드는 실행되지 않습니다.
	// 하지만 패닉이 발생한 함수를 빠져나온 후의 실행은 재개됩니다.
	fmt.Println("This will not be printed")
}

func main() {
	// 가장 바깥쪽의 recover defer가 마지막에 실행됩니다.
	defer calmDown()
	fmt.Println(recover()) // <nil>

	freakOut() // omg
	fmt.Println("This will be printed.")

	err := fmt.Errorf("there's an error")
	panic(err)
}
