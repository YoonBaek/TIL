// Author : Github YoonBaek
// 이제 명령줄에서 인자를 받아 평균을 구하는 프로그램을 만듭니다.
// 한 번 책 안 보고 만들어봄.

package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	sum := 0.0
	// args가 없으면 에러를 반환하고 종료
	errNoValue := errors.New("input please")
	if len(args) == 0 {
		log.Fatal(errNoValue)
	}
	for _, arg := range args {
		num, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			log.Fatal(err)
		}
		sum += num
	}
	cnt := float64(len(args))
	fmt.Printf("Avg : %0.2f\n", sum/cnt)
}
