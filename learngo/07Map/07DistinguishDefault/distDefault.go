// Author : Github YoonBaek
/*
제로값이 값의 할당 여부를 판단하는 데에 유용하긴 하나,
제로값 만으로는 키 값이 제로값으로 할당된건지 아니면 키 자체가 조재하지 않는 건지 판단하기는 어렵습니다.
*/
package main

import "fmt"

var grades = map[string]float64{"Alma": 0, "Rohit": 86.5}

func status(name string) {
	grade := grades[name]
	if grade < 60 {
		fmt.Printf("%s is failing!\n", name)
	}
}

func statusFix(name string) {
	grade, ok := grades[name]
	if !ok {
		fmt.Printf("No grade recorded for %s.\n", name)
	} else if grade < 60 {
		fmt.Printf("%s is failing!\n", name)
	}
}

func main() {
	status("Alma")
	status("Carl")
	// Carl이 낙제했다는 잘못된 내용을 보고하고 있습니다.
	// map에는 Carl의 성적이 기록되지 않아 default값이 입력되었기 때문입니다.

	// 이런 상황을 피하기 위해 맵 키에 접근할 때 선택적으로 반환되는 bool값을 갖는 두번째 반환 값을 사용할 수 있습니다.
	// 이 값은 접근하고 있는 키 값이 맵에 존재하는 경우에는 true 값을, 존재하지 않는 경우에는 false 값을 갖습니다.
	// 대부분 Go 개발자는 이 값을 ok라는 변수에 할당하곤 합니다. 직관적이고 간단 명료하기 때문이죠.

	cnters := map[string]int{"a": 3, "b": 0}
	var value int
	var ok bool

	value, ok = cnters["a"]
	fmt.Println(value, ok)
	value, ok = cnters["b"]
	fmt.Println(value, ok)
	value, ok = cnters["c"]
	fmt.Println(value, ok) // ok가 false로 출력됩니다.

	// key의 존재 여부만 확인하고 싶을 때에는 _로 value를 무시하면 됩니다.

	// ok 값으로 함수를 수정해줍니다.
	statusFix("Alma")
	statusFix("Carl")
}
