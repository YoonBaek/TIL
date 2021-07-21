// Author : Github YoonBaek
// 개표 프로그램에서 맵 사용하기.

package main

import (
	"fmt"
	"log"

	getString "github.com/YoonBaek/learngo/07Map/01GetString"
)

func main() {
	lines, err := getString.GetStrings("../votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++
	}
	fmt.Println(counts)
}
