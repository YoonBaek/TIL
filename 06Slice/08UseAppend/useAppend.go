// Author : Github Yoon Baek
// 이제 3개 이상의 데이터가 들어와도 처리할 수 있습니다!
package main

import (
	"fmt"
	"log"

	readFloats "github.com/YoonBaek/learngo/06Slice/07ReadFloats"
)

func main() {
	numbers, err := readFloats.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	sum := 0.0
	for _, number := range numbers {
		sum += number
	}
	cnt := float64(len(numbers))
	fmt.Printf("Avg : %0.2f\n", sum/cnt)
}
