package realProse

import "testing"

// 테스트 파일 내에서도 타입 정의가 가능합니다.
type testData struct {
	list []string
	want string
}

// test 함수는 단일 매개변수로 testing.T 값의 포인터를 받습니다.
func TestJoinWithCommas(t *testing.T) {
	tests := []testData{
		testData{list: []string{}, want: ""},
		testData{list: []string{"apple"}, want: "apple"},
		testData{list: []string{"apple", "orange"}, want: "apple and orange"},
		testData{list: []string{"apple", "orange", "mango"}, want: "apple, orange, and mango"},
	}
	for _, test := range tests {
		got := JoinWithCommas(test.list)
		want := test.want
		if got != want {
			t.Errorf("JoinWithCommas(%#v) = \"%s\", want \"%s\"\n", test.list, got, want)
		}
	}
}
