// Learned by Github YoonBaek

package main

// 함수 호출을 스택에 추가합니다.
func main() {
	one()
}

// 스택에 또다른 호출도 추가하빈다.
func one() {
	two()
}

// 세번째 호출도 추가합니다.
func two() {
	three()
}

// 패닉발생! 스택트레이스는 위의 모든 호출을 포함합니다.
func three() {
	panic("This call's stack is too deep for me")
}
