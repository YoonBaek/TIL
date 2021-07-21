// Learned by Github YoonBaek
// This code is for learning "embedding"
package calendar

import (
	"errors"
	"unicode/utf8"
)

type Event struct {
	// Title string
	title string // title 필드 캡슐화 하기. 제목의 길이를 제한하는 유효성 검증을 하기 위함!
	Date
}

// Setter Method
func (e *Event) SetTitle(s string) error {
	if utf8.RuneCountInString(s) > 30 {
		// 제목의 길이가 30자를 넘어가면 에러를 반환합니다.
		return errors.New("invalid title. too long")
	}
	e.title = s
	return nil
}

// Access Method
func (e *Event) Title() string {
	return e.title
}
