// Author : Github Yoon Baek
// Go에서는 데이터 컬렉션을 저장하는 또 다른 방식인 맵(map)이라고 하는 자료구조가 있습니다.
// map이란 저장된 값을 키를 통해 접근할 수 있는 컬렉션입니다.
// Python의 딕셔너리와 비슷...
// 키를 사용하면 맵의 데이터를 아주 편리하게 가져올 수 있습니다. 맵은 마치 널부러진 서류더미가 아닌 잘 라벨링된 서류 문서함과 같습니다.

// Array와 Slice에서는 인덱스 값으로 정수만 사용할 수 있는 반면 맵에서는 모든 타입의 값을 키로 사용할 수 있습니다.
// 단 , ==를 사용하여 서로 비교할 수 있는 값이어야 합니다. 여기에는 숫자, 문자열 등도 포함됩니다. 키는 모두 통일한 타입이어야하고
// 값 또한 모두 동일한 타입이어야 하지만 키와 값이 동일한 타입일 필요는 없습니다.

package main

import "fmt"

func main() {
	// how to declare
	// var name, map[key type]val type
	var ranks map[string]int
	fmt.Printf("%#v\n", ranks)
	// generate real value of ranks
	ranks = make(map[string]int)
	fmt.Printf("%#v\n", ranks)

	// 이렇게 한 번에 할 수도 있죠.
	ranksOneWay := make(map[string]int)
	fmt.Printf("%#v\n", ranksOneWay)

	// 키는 어떤 타입도 올 수 있습니다.
	// 키 : string, 값 : 정수
	ranks["gold"] = 1
	ranks["silver"] = 2
	ranks["bronze"] = 3

	fmt.Println(ranks["bronze"], ranks["gold"])

	// 키 : string 값 : string
	elems := make(map[string]string)
	elems["H"] = "Hydrogen"
	elems["Li"] = "Lithium"
	fmt.Println(elems)

	// 키 : int, 값 : bool
	isPrime := make(map[int]bool)
	isPrime[2] = true
	isPrime[3] = true
	isPrime[5] = true
	fmt.Println(isPrime)
}
