// Learned By Github YoonBaek
// 사용자 정의 타입 별로 단위가 기본 타입처럼 똑같이 표기 되면 구분이 생각보다 어렵습니다.
// 이를 구분하기 위해 출력 양식을 지원해주는 것이 빌트인 인터페이스인 Stringer 입니다.

package main

import "fmt"

/*
	인터페이스의 정의는 다음과 같습니다.
	type Stringer interface {
		String() string
	}

	문자열을 반환하는 String 메서드를 가진 모든 타입은 fmt.Stringer 인터페이스를 만족합니다.
*/
// 예시
type CoffeePot string

// Stringer 인터페이스를 충족시켜 줍니다.
func (c CoffeePot) String() string {
	return string(c) + " coffe pot"
}

// 이제 앞서 정의 했던 Gallon, Liter, Milliliter 등의 단위에도 적용할 수 있습니다.
type Gallons float64
type Liters float64
type Milliliters float64

// Sprintf 함수를 사용하면 지정 형식으로 바로 출력할 수 있습니다.
func (g Gallons) String() string {
	return fmt.Sprintf("%.2f gal", g)
}
func (l Liters) String() string {
	return fmt.Sprintf("%.2f L", l)
}
func (ml Milliliters) String() string {
	return fmt.Sprintf("%.2f mL", ml)
}
func main() {
	coffeepot := CoffeePot("LuxBrew")
	fmt.Println(coffeepot.String())

	fmt.Println(Gallons(12.06))
	fmt.Println(Liters(12.06))
	fmt.Println(Milliliters(12.06))

}
