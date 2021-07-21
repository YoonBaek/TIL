// Author : Github YoonBaek
// Reference : Head First Go
/*
	모든 슬라이스는 내부 배열을 기반으로 구현되었습니다.
	내부 배열은 슬라이스의 데이터가 실제로 저장되는 공간이며,
	슬라이스는 단지 이 배열 원소의 일부 또는 전체에 대한 추상화 된 뷰일 뿐입니다.

	make함수 또는 슬라이스 리터럴로 슬라이스를 생성하면 내부배열이 자동으로 생성됩니다.
	슬라이스를 거치지 않고 직접 내부 배열에 접근할 수는 없습니다.
	그러나 배열을 직접 생성한 다음 슬라이스 연산자를 사용하면
	아래 예시와 같이 해당배열을 기반으로 하는 슬라이스를 만들 수도 있습니다.

	mySlice := myArray[1:3]
*/
package main

import (
	"fmt"
	"reflect"
)

func main() {
	// how to slice Slice
	underlyingArray := [5]string{"a", "b", "c", "d", "e"}
	slice1 := underlyingArray[0:3]
	fmt.Println(reflect.TypeOf((slice1)), slice1)

	/*
		슬라이싱 인덱스의 문법은 파이썬과 동일합니다.
		i부터 j-1까지 인덱싱.
	*/
}
