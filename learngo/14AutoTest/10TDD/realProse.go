// 테스트 대상 패키지입니다.
package realProse

import "strings"

func JoinWithCommas(phrase []string) string {
	// 케이스를 나눠 모두 패스했다.
	if len(phrase) <= 1 {
		return strings.Join(phrase, "")
	}
	if len(phrase) == 2 {
		return strings.Join(phrase, " and ")
	}
	result := strings.Join(phrase[:len(phrase)-1], ", ")
	result += ", and "
	result += phrase[len(phrase)-1]
	return result
}
