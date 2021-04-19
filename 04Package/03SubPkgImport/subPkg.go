// Author : Github YoonBaek

package main

import (
	"github.com/YoonBaek/learngo/04Package/01MakePackage/greeting"
	"github.com/YoonBaek/learngo/04Package/01MakePackage/greeting/deutsch"
	// Subpkg imported, we use last folder's pkg name.
	// Once imported, go refers just packagename
)

func main() {
	greeting.Hello()
	greeting.Hi()
	deutsch.Hallo()
	deutsch.GutenTag()
}
