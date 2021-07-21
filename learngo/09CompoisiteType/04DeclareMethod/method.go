// Learned by Github YoonBaek
// 메서드 선언과 리시버 매개변수

package main

import (
	"fmt"
	"strings"
)

type User string

// 메서드는 리시버 위에서 호출한다는 점만 제외하면 함수와 유사합니다.
// func (리시버 매개변수) 함수명(함수매개변수){} 로 선언합니다.
// 또한 메서드도 마찬가지로 대문자로 시작하면 패키지 외부로 export 할 수 있습니다.
func (u User) SayHi() {
	fmt.Println("Hi from", u)
}

//함수와 마찬가지로 한 개 이상의 반환 값을 가질 수도 있습니다.
func (u User) WhatisUserName() (string, string) {
	names := strings.Split(string(u), " ")
	return names[0], names[1]
}

func main() {
	user := User("Yoon Baek")
	user.SayHi()
	secondUser := User("Snowy Baek")
	secondUser.SayHi()
	first, second := user.WhatisUserName()
	fmt.Printf("%s's first name is %s, last name is %s\n", user, first, second)

}
