package main

import (
	"fmt"
	"reflect"
)

func main() {
	// string & escape sequence
	// ""
	fmt.Println("\"string & escape sequence\"")
	fmt.Println("Be A Gopher!")
	fmt.Println("Be \nA Gopher!")
	fmt.Println("Be \tA Gopher!")
	fmt.Println("Be \\A Gopher!")
	fmt.Println("")

	// rune
	// ''
	fmt.Println("\"rune\"")
	fmt.Println("A :", 'A')
	fmt.Println("B :", 'B')
	fmt.Println("Is Korean 한 Okay? :", '한')
	fmt.Println("")

	// bool
	// true / false
	fmt.Println("\"bool\"")
	fmt.Println(true)
	fmt.Println(false)
	fmt.Println("")

	// int & float
	fmt.Println("\"int & float\"")
	fmt.Println("int :", 42)
	fmt.Println("float :", 3.1415)
	fmt.Println("")

	// operators
	// works as same way as other languages.
	fmt.Println("\"operators\"")
	fmt.Println("1 + 2 =\t", 1+2)
	fmt.Println("5.4 - 2.2 =", 5.4-2.2)
	fmt.Println("3 * 4 =\t", 3*4)
	fmt.Println("7.5 / 5 =", 7.5/5)
	fmt.Println("4 < 6\t", 4 < 6)
	fmt.Println("4 > 6\t", 4 > 6)
	fmt.Println("2 + 2 == 5", 2+2 == 5)
	fmt.Println("2 + 2 != 5", 2+2 != 5)
	fmt.Println("4 <= 6\t", 4 <= 6)
	fmt.Println("4 >= 4\t", 4 >= 4)
	fmt.Println("")

	fmt.Println("\"get type of values\"")
	fmt.Println("42\t", reflect.TypeOf(42))
	fmt.Println("3.1415\t", reflect.TypeOf(3.1415))
	fmt.Println("true\t", reflect.TypeOf(true))
	fmt.Println("BAG\t", reflect.TypeOf("Be a gopher!"))
	fmt.Println("")
}
