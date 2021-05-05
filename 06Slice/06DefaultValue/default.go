// Author : Github YoonBaek
package main

import "fmt"

func main() {
	// 슬라이스도 똑같이 선언된 크기의 디폴트 값을 갖는다.
	floatSlice := make([]float64, 10)
	boolSlice := make([]bool, 10)
	for i := 0; i < 10; i++ {
		fmt.Printf("%.1f ", floatSlice[i])
	}
	fmt.Println()
	for i := 0; i < 10; i++ {
		fmt.Printf("%t ", boolSlice[i])
	}
	fmt.Println()

	// 크기가 선언되지 않은 슬라이스는?
	// 아무 슬라이스도 할당되지 않은 슬라이스는 nil값을 갖게 된다.
	var intSlice []int
	var strSlice []string
	fmt.Printf("int: %#v, str: %#v\n", intSlice, strSlice)

	// 따라서 append나 len함수도 nil값을 갖고 있는 슬라이스는 빈 슬라이스처럼 처리합니다.
	fmt.Println(len(intSlice))
	intSlice = append(intSlice, 27)
	fmt.Printf("int: %#v\n", intSlice)

	// 결국 빈 슬라이스인지 nil 슬라이스인지는 크게 중요하지 않습니다.
	// 알아서 잘 하니까!
	var slice []string
	if len(slice) == 0 {
		slice = append(slice, "first elem")
	}
	fmt.Printf("%#v\n", slice)
}
