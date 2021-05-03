// Author : Github YoonBaek
// 배열을 만들 때 배열에 포함된 모든 값은 배열이 가진 타입의
// 제로값으로 초기화됩니다.

package main

import "fmt"

func main() {
	// Let's see the default value of integer arr
	var primes [5]int
	primes[0] = 2

	for _, prime := range primes {
		fmt.Println(prime)
	} // you can see default = 0

	// default val of string arr
	var notes [5]string
	notes[0] = "be"
	notes[2] = "a"
	notes[4] = "gopher"
	for _, note := range notes {
		fmt.Println(note)
	} // you can see default = ""

	// Thanks to default value,
	// We can manage arrays safe.
	// For instance, we can do like this without assigning a value

	var cnts [3]int
	cnts[0]++
	cnts[1] += 2
	cnts[2] += 3
	for _, cnt := range cnts {
		fmt.Println(cnt)
	}
}
