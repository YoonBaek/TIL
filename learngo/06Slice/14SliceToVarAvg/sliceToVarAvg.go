// Author : Github Yoon Baek
// 가변인자함수에 슬라이스 전달하기

package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
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
	args := os.Args[1:]
	numbers := []float64{}

	for _, arg := range args {
		number, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}
	fmt.Println("Avg :", avg(numbers...))
	// 슬라이스 뒤에 ...을 붙임으로서 가변함수처럼 처리할 수 있습니다.
}
