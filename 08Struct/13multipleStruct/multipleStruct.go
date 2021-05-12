// Learned by Github YoonBaek
// 10Export struct참고하기
// 여러 구조체 만들어서 import 하기, field type으로 struct사용하기

package main

import (
	"fmt"

	magazine "github.com/YoonBaek/learngo/08Struct/10ExportStruct"
)

func main() {
	// 새로운 구조체 Address에 값을 할당해서 미리 구현한 SubScriber의 내부 struct로 할당합니다.
	addr := magazine.Address{Street: "Gangnamdaero, Seochogu",
		City:       "Seoul",
		State:      "South Korea",
		PostalCode: "12345"}

	subscriber := magazine.SubScriber{Name: "Yoon Baek", Home: addr}
	fmt.Println(subscriber)

	// Chaining을 활용하면, 내부 Struct의 필드값에도 접근할 수 있습니다.
	subscriber.Home.Street = "Dogok-ro"
	fmt.Println(subscriber.Home.Street)

	// 연산자 체이닝으로만 값 할당해보기.
	emp := magazine.Employee{Name: "Snowy"}
	emp.Home.City = "Seoul"
	emp.Home.Street = "Gangnamdaero"
	emp.Home.State = "South Korea"
	emp.Home.PostalCode = "12345"

	fmt.Println("Employee Name :", emp.Name)
	fmt.Println("Employee Salary :", emp.Salary)
	fmt.Println("Street :", emp.Home.Street)
	fmt.Println("City :", emp.Home.City)
	fmt.Println("State :", emp.Home.State)
	fmt.Println("PostalCode :", emp.Home.PostalCode)

}
