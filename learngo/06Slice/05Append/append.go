// Author : Github YoonBaek
// Different from array, slice can add elements
// 이는 기존 배열의 원소를 더 큰 배열로 복사하는 방식으로 이루어집니다.

package main

import "fmt"

func main() {
	slice := []string{"a", "b"}
	fmt.Println(slice, len(slice))

	slice = append(slice, "c")
	fmt.Println(slice, len(slice))

	// You can add multiple elems
	slice = append(slice, "d", "e")
	fmt.Println(slice, len(slice))

	// 유심히 보면, append는 기존 슬라이스 변수에 재할당하는 방식으로 작동하는 것을 알 수 있습니다.

	s1 := []string{"s1", "s1"}
	s2 := append(s1, "s2", "s2")
	s3 := append(s2, "s3", "s3")
	s4 := append(s3, "s4", "s4")
	s4[0], s3[1], s2[0] = "X", "Y", "W"
	fmt.Println(s1, s2, s3, s4)
	/* 보는 바와 같이 s3와 s4는 내부 원소를 공유하고 있음을 알 수 있습니다.
	이는 마지막 append 함수는 append 이전의 배열의 메모리를 공유하고 있기 때문입니다.
	따라서 값이 같이 변할 수가 있는 만큼, 하나의 변수로만 재할당하는 것을 원칙으로 합니다.
	아래와 같이 말이죠.*/
	s0 := []string{"s1", "s1"}
	s0 = append(s0, "s2", "s2")
	s0 = append(s0, "s3", "s3")
	s0 = append(s0, "s4", "s4")
	fmt.Println(s0)
}
