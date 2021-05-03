// Author : Github YoonBaek
/*
for ... range 루프에는 복잡한 초기화문, 조건문 및 후처리문이 없습니다.
변수에는 각 원소의 값이 자동으로 할당되기 때문에 유효하지 않은 배열의 인덱스에 접근할 위험이 없습니다.
코드가 더 안전하고 읽기 쉽기 때문에 배열이나 다른 컬렉션의 데이터 구조를 다룰 대에는 range문 형태의
for 루프를 가장 많이 사용합니다.
*/

package main

import "fmt"

func main() {
	notes := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}

	for index, note := range notes {
		fmt.Println(index, note)
	}

	// if index is not needed, you can ignore it like this.
	for _, note := range notes {
		fmt.Println(note)
	}

	// it's also same with values.
	for idx, _ := range notes {
		fmt.Println(idx)
	}
}
