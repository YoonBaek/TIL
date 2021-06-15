// Learned by Github YoonBaek

package main

import (
	"fmt"

	prose "github.com/YoonBaek/learngo/14AutoTest/01Prose"
)

func main() {
	phrases := []string{"tomato", "banana", "melon"}
	fmt.Println("Fruits :", prose.JoinWithCommas(phrases))
	// 두 개 일 때는 쉼표가 없어야 하는데 들어 있습니다!
	phrases = []string{"my parents", "rodeo clown"}
	fmt.Println("A photo of", prose.JoinWithCommas(phrases))
	// 세 개일 때는 정상적으로 보입니다.
	phrases = []string{"my parents", "rodeo clown", "a prize bull"}
	fmt.Println("A photo of", prose.JoinWithCommas(phrases))
}
