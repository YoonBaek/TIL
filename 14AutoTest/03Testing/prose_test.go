package prose

// 표준 라이브러리
import (
	"strings"
	"testing"
)

func TestTwoElements(t *testing.T) {
	list := []string{"apple", "orange"}
	if JoinWithCommas(list) != "apple and orange" {
		t.Error("didn't match expected value")
	}
}

func TestThreeElements(t *testing.T) {
	list := []string{"apple", "orange", "mango"}
	if JoinWithCommas(list) != "apple, orange, and mango" {
		t.Error("didn't match expected value")
	}
}

func JoinWithCommas(phrase []string) string {
	result := strings.Join(phrase[:len(phrase)-1], ", ")
	result += ", and "
	result += phrase[len(phrase)-1]
	return result
}
