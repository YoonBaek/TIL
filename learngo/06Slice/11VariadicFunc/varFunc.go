// Author : Github Yoon Baek
// Append나 fmt.Println과 같이 받을 수 있는 인자의 갯수가 정해져 있지 않은 함수를 만들어 보자

package main

import (
	"fmt"
	"math"
)

func severalInts(nums ...int) {
	fmt.Println(nums)
}

func severalStrings(strings ...string) {
	fmt.Println(strings)
}

func mixed(num int, flag bool, strings ...string) {
	fmt.Println(num, flag, strings)
}

//가변인자 사용한 본격 함수!
func maximum(numbers ...float64) float64 {
	// 최댓값을 가장 작은 음의 무한대로 설정
	max := math.Inf(-1)
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	return max
}

func inRange(min float64, max float64, numbers ...float64) (result []float64) {
	for _, num := range numbers {
		if num >= min && num <= max {
			result = append(result, num)
		}
	}
	return result
}

func main() {
	// 가변 인자는 여러 인수들을 슬라이스의 형태로 받습니다.
	severalInts(1, 2, 3, 4)

	// 가변함수는 인자를 전달하지 않아도 에러가 나지 않으며
	// 빈 슬라이스를 반환합니다.
	severalStrings()
	severalStrings("a", "b", "gopher")

	// 고정 인자와 함께 받을수도 있습니다.
	// 다만, 고정 인자는 반드시 전달해줘야 하며, 가변 인자의 앞에서만 사용할 수 있습니다.
	mixed(1, true, "Be", "A", "Gopher")
	fmt.Println(maximum(100, 200))
	fmt.Println(inRange(1, 100, -12.5, 3.2, 0, 50, 103.5))
}
