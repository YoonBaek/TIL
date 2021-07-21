// Practice solved by github YoonBaek
// make function lie.

package main

import "fmt"

func negate(myBool *bool) {
	*myBool = !*myBool
}

func main() {
	truth := true
	negate(&truth)
	fmt.Println(truth)

	lie := false
	negate(&lie)
	fmt.Println(lie)
}
