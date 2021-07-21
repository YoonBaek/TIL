// Author : Github YoonBaek
// 배열과 슬라이스와 마찬가지로 아무 값도 할당하지 않은 맵의 키에 접근하면 제로값이 반환됩니다.
// 개념 : value의 디폴트 값, 맵 자체의 디폴트값

package main

import "fmt"

func main() {
	numbers := make(map[string]int)
	numbers["Assigned"] = 12
	//할당된 값 출력
	fmt.Printf("assigned : %#v\n", numbers["Assigned"])
	//할당되지 않은 값 출력
	fmt.Printf("not assigned : %#v\n", numbers["NotAssigned"])

	// 슬라이스나 배열처럼 값의 타입에 따라 제로 값은 바뀝니다.
	words := make(map[string]string)
	words["Assigned"] = "hi"
	// 할당된 값 출력
	fmt.Printf("assigned : %#v\n", words["Assigned"])
	//할당되지 않은 값 출력
	fmt.Printf("not assigned : %#v\n", words["NotAssigned"])

	// 제로값 덕분에 할당하지 않은 값도 안전하게 처리되고 있음을 볼 수 있습니다.
	// 응용
	cnters := make(map[string]int)
	cnters["Be"]++
	cnters["A"]++
	cnters["Gopher"]++
	fmt.Println(cnters)

	// 슬라이스처럼 맵 변수 자체도 제로 값은 nil맵입니다.
	// 따라서 슬라이스처럼 변수 선언 이후 make나 빈 맵을 통해 생성해 주어야 panic이 발생하지 않습니다.

	// var emptyMap map[string]int
	// var emptySlice []int

	// emptySlice[1] = 0 // 해제하면 패닉이 일어납니다.
	// emptyMap["panic"] = 0 // 이것도 마찬가지.

	var notNilMap = make(map[int]string)
	notNilMap[3] = "three"
	fmt.Printf("%#v\n", notNilMap)
}
