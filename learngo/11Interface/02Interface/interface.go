// Learned by Github YoonBaek

package main

import "fmt"

// how to declare interface
type myInterface interface {
	MethodWithoutParams()
	MethodWithParams(float64)
	MethodWithReturn() string
}

// 파라미터가 없는 메서드를 가지고, float64를 매개변수로 갖는 메서드를 가지고, 리턴 값으로 string을 반환하는 메서드를
// 가진 타입은 이 인터페이스를 만족합니다.

// myInterface를 만족 시킬 타입을 하나 만들어봅시다.
type MyType int

// 메서드를 충족시킵니다.
// 정렬할 때 이 기능 자주 사용하니 알아두면 좋습니다.
func (m MyType) MethodWithoutParams() {
	fmt.Println("MethodWithoutParmas called")
}

func (m MyType) MethodWithParams(f float64) {
	fmt.Println("MethodWithParams called with", f)
}

func (m MyType) MethodWithReturn() string {
	return "Hi from MethodWithReturn"
}

// 인터페이스와 무관한 메서드가 있더라도 여전히 이 인터페이스를 만족할 수 있습니다.
func (m MyType) MethodUnrelatedWithInterface() {
	fmt.Println("MethodUnrelatedWithInterface called")
}

// 사용해보기
func main() {
	// 인터페이스 타입을 가진 변수를 선언합니다.
	var value myInterface
	// MyType의 값은 myInterface를 만족하고 있기 때문에
	// myInterface 타입의 변수에 할당할 수 있습니다.
	value = MyType(5)
	value.MethodWithoutParams()
	value.MethodWithParams(1.2)
	fmt.Println(value.MethodWithReturn())
}
