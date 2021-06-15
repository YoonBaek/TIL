// Learned by Github YoonBaek
// 바꾼 결과 확인하기.
package main

import (
	"fmt"
	"strings"

	prose "github.com/YoonBaek/learngo/14AutoTest/01Prose"
)

func JoinWithCommas(phrase []string) string {
	if len(phrase) == 2 {
		return strings.Join(phrase, " and ")
	}
	result := strings.Join(phrase[:len(phrase)-1], ", ")
	result += ", and "
	result += phrase[len(phrase)-1]
	return result
}

func main() {
	phrases := []string{"tomato", "banana", "melon"}
	fmt.Println("Fruits :", prose.JoinWithCommas(phrases))
	fmt.Println("Fruits :", JoinWithCommas(phrases))
	// 문자열이 두 개 일 때 마침표를 잘 제거한 채 동작
	phrases = []string{"my parents", "rodeo clown"}
	fmt.Println("A photo of", prose.JoinWithCommas(phrases))
	fmt.Println("A photo of", JoinWithCommas(phrases))
	// 문자열이 세 개 일 때도 변함없이 잘 동작한다.
	phrases = []string{"my parents", "rodeo clown", "a prize bull"}
	fmt.Println("A photo of", prose.JoinWithCommas(phrases))
	fmt.Println("A photo of", JoinWithCommas(phrases))
}
