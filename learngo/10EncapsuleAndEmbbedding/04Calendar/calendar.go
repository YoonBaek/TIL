// Learned by Github YoonBaek

package calendar

import (
	"errors" // 에러값을 생성하는 기능을 하는 패키지입니다.
)

type Date struct {
	// 필드를 외부에서 수정할 수 없도록 숨겨버립니다.
	year  int
	month int
	day   int
}

// Setter Method
// 메서드는 외부에서 접근할 수 있도록 유지해줍니다.
func (d *Date) SetYear(y int) error {
	if y < 1 {
		return errors.New("invalid year! year should be larger than 1")
	}
	d.year = y
	return nil
}

func (d *Date) SetMonth(m int) error {
	if m < 1 || m > 12 {
		return errors.New("invalid month! month must be between 1 and 12")
	}
	d.month = m
	return nil
}

func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("invalid day! day must be between 1 and 31")
	}
	d.day = day
	return nil
}

// Access Method
func (d *Date) Year() int {
	return d.year
}

func (d *Date) Month() int {
	return d.month
}

func (d *Date) Day() int {
	return d.day
}
