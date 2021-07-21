// Learned by Github YoonBaek
package prose

import (
	"fmt"
	"strings"
	"testing"
)

// 헬퍼함수. Test로 시작하지 않으면 됩니다.
// errorString을 통해 중복을 줄일 수 있습니다.
func errorString(list []string, got, want string) string {
	//errorf를 활용해 추가정보 포함
	return fmt.Sprintf("JoinWithCommas(%#v) = \"%s\", want \"%s\"\n", list, got, want)
}

func TestTwoElements(t *testing.T) {
	list := []string{"apple", "orange"}
	want := "apple and orange"
	got := JoinWithCommas(list)
	if got != want {
		t.Error(errorString(list, want, got))
	}
}

func TestThreeElements(t *testing.T) {
	list := []string{"apple", "orange", "mango"}
	want := "apple, orange, and mango"
	got := JoinWithCommas(list)
	if got != want {
		t.Error(errorString(list, want, got))
	}
}

// 테스트하고자 하는 함수
func JoinWithCommas(phrase []string) string {
	// 케이스를 나눠 모두 패스했다.
	if len(phrase) == 2 {
		return strings.Join(phrase, " and ")
	}
	result := strings.Join(phrase[:len(phrase)-1], ", ")
	result += ", and "
	result += phrase[len(phrase)-1]
	return result
}
