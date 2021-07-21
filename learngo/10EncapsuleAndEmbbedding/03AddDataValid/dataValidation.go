// Learned by Github YoonBaek
// error를 통한 유효성 검증.

package main

import (
	"errors" // 에러값을 생성하는 기능을 하는 패키지입니다.
	"fmt"
	"log" // 에러를 보고하고 종료하는 기능을 하는 패키지입니다.

	calendar "github.com/YoonBaek/learngo/10EncapsuleAndEmbbedding/04Calendar"
)

type Date struct {
	Year  int
	Month int
	Day   int
}

func (d *Date) SetYear(y int) error {
	if y < 1 {
		return errors.New("invalid year! year should be larger than 1")
	}
	d.Year = y
	return nil
}

func (d *Date) SetMonth(m int) error {
	if m < 1 || m > 12 {
		return errors.New("invalid month! month must be between 1 and 12")
	}
	d.Month = m
	return nil
}

func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("invalid day! day must be between 1 and 31")
	}
	d.Day = day
	return nil
}

//code snippet
func main() {
	date := Date{}
	err := date.SetYear(2021)
	// year에 잘못된 값을 전달할 경우 프로그램을 종료합니다.
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date)

	err = date.SetMonth(5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date)

	err = date.SetDay(32)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date)

	// 하지만 아직 리터럴로 직접 설정할 수 있는 여지가 있습니다!
	// 이는 필드를 숨겨서 해결할 수 있습니다.
	// 다음 장에서 만든 calendar 패키지를 활용해 해결해 보겠습니다.

	// calendar 에서는 필드를 숨겨 놓았습니다. 이렇게 리터럴 할당을 할 수 없습니다.
	// importedDate := calendar.Date{-1, 2, 3} // unexpected lit field error
	importedDate := calendar.Date{}

	// 이제 규정된 세터 메서드로만 값을 입력할 수 있습니다.
	// 따라서 유효성 검증도 빈틈없이 가능합니다.
	importedDate.SetYear(2020)
}
