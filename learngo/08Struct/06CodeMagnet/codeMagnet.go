// Solved by Github Yoon Baek
// Name : "Alonzo Cole"
// Grade : 92.3
// 위와 같이 출력되도록 코드 자석 퍼즐 맞추기
package main

import "fmt"

type student struct {
	name  string
	grade float64
}

func printInfo(s student) {
	fmt.Println("Name :", s.name)
	fmt.Println("Grade :", s.grade)
}

func main() {
	var s student
	s.name = "Alonzo"
	s.grade = 92.3

	printInfo(s)
}
