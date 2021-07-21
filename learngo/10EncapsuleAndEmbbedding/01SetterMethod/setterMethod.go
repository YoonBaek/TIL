// Learned by Github YoonBaek
// Setter Method를 통한
// data validation 데이터 유효성 검증

package main

import "fmt"

// 연,월, 일은 그자체로는 의미가 없기 때문에 하나로 묶기에 제 격입니다.
// struct를 만들어 봅니다.
type Date struct {
	Year  int
	Month int
	Day   int
}

// 설정자 메서드
// 포인터를 넘겨줌으로서 주소에 위치한 실제 값을 변경하도록 합니다.
func (d *Date) SetYear(y int) {
	d.Year = y
}

func (d *Date) SetMonth(m int) {
	d.Month = m
}

func (d *Date) SetDay(day int) {
	d.Day = day
}

func main() {
	// struct를 이용해 값을 만들어봅니다.
	date := Date{2021, 5, 20}
	fmt.Println(date)

	// 실행해 보면 값이 정상적으로 출력되고 있습니다.
	// 그러나, 데이터 유효성의 문제가 생깁니다.
	// 실제 날짜에서는 불가능한 아래와 같은 값의 입력도 가능하기 때문입니다.

	weirdDate := Date{Year: -30, Month: -1, Day: 0}
	fmt.Println(weirdDate)

	// Year 필드에 음수가 들어가거나,
	// Month 필드에 1 - 12 이외의 값이 들어가거나,
	// Day 필드에 1-31 이외의 값이 들어가도 막을 길이 없습니다.

	// 이를 막기 위해 필요한 일을 데이터 유효성 검증이라 합니다.

	// 유효성 검증을 위해서 설정자 메서드 (setter method)를 사용합니다.
	// go언어에서 내부적으로 처리해주는 덕분에
	// 일반적인 assign과 똑같이 값을 변경해줄 수 있습니다.
	weirdDate.SetYear(2021)
	// year가 잘 변경되어 있음을 확인할 수 있습니다.
	fmt.Println(weirdDate.Year)

	// 하지만 이렇게 한다 한들, 유효성 검증 문제가 해결이 되지는 않습니다.
	// 다음 장에서 알아보겠습니다.
}
