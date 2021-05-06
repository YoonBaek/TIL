// Author : Github YoonBaek
// 배열이나 슬라이스와 마찬가지로 맵에서도 키와 값을 이미 알고 있는 경우 맵 리터럴을 사용할 수 있습니다.
// 맵 리터럴은 맵 타입으로 시작해 키/값 쌍들이 중괄호 안에 감싸여 따라옵니다.
// 키 값 쌍은 키, 콜론 그리고 값으로 구성되며 복수 개의 키/값 쌍들은 쉼표로 구분됩니다.

package main

import "fmt"

func main() {
	ranks := map[string]int{"bronze": 3, "silver": 2, "gold": 1}
	fmt.Println(ranks)

	elems := map[string]string{
		"H":  "Hyrdrogen",
		"Li": "Lithium",
	}
	fmt.Println(elems)

	// 또한 리터럴로 빈 map을 생성할 수도 있습니다.
	emptyMap1 := make(map[string]float64)
	emptyMap2 := map[string]float64{}
	fmt.Println(emptyMap1, emptyMap2)
}
