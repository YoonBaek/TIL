// Learned by Github YoonBaek

package prose

import "strings"

func JoinWithCommas(phrase []string) string {
	result := strings.Join(phrase[:len(phrase)-1], ", ")
	result += ", and "
	result += phrase[len(phrase)-1]
	return result
}
