// Author : Github Yoon Baek
// 빌트인 delete 함수 이용해 키/값 쌍 삭제하기
// delete(map[keytype]valtype, keytype) 의 방식으로 삭제

package main

import "fmt"

func main() {
	var ok bool
	ranks := make(map[string]int)
	rank := 0
	// 브론즈 키에 할당합니다.
	ranks["bronze"] = 3
	rank, ok = ranks["bronze"]
	// ok는 true가 됩니다.
	fmt.Printf("rank : %d, ok : %v\n", rank, ok)
	// 브론즈 키와 해당 키의 값을 삭제합니다.
	delete(ranks, "bronze")
	rank, ok = ranks["bronze"]
	// ok는 false가 됩니다.
	fmt.Printf("rank : %d, ok : %v\n", rank, ok)

	isPrime := make(map[int]bool)
	prime := false
	// 브론즈 키에 할당합니다.
	isPrime[5] = true
	prime, ok = isPrime[5]
	// ok는 true가 됩니다.
	fmt.Printf("prime : %v, ok : %v\n", prime, ok)
	// 브론즈 키와 해당 키의 값을 삭제합니다.
	delete(isPrime, 5)
	prime, ok = isPrime[5]
	// ok는 true가 됩니다.
	fmt.Printf("prime : %v, ok : %v\n", prime, ok)
}
