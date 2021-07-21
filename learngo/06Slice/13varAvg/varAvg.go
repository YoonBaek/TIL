// Author : Github YoonBaek
// 가변 인자 함수를 사용해 평균 계산하기

package main

import (
	"errors"
	"fmt"
	"log"
)

func avg(numbers ...float64) float64 {
	sum := 0.0
	err := errors.New("this function requires more than 1 value")
	if len(numbers) <= 0 {
		log.Fatal(err)
	}
	for _, number := range numbers {
		sum += number
	}
	cnt := float64(len(numbers))
	return sum / cnt
}

func main() {
	fmt.Println(avg(100, 50))
	fmt.Println(avg(90.7, 89.7, 98.5, 92.3))
	fmt.Println(avg())
}
