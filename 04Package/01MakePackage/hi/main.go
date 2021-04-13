package main

import (
	"fmt"

	calc "github.com/YoonBaek/learngo/04Package/01MakePackage/calc"
	"github.com/YoonBaek/learngo/04Package/01MakePackage/greeting"
)

// to import user package, the most upper folder is inside "src" folder.
// if you put this repository inside the src

func main() {
	greeting.Hello()
	greeting.Hi()
	fmt.Println(calc.Add(1, 2))
}
