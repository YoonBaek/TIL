// Author : Github YoonBaek
package main

import (
	"fmt"
	. "fmt"
)

func main() {
	width, height, paintArea := 0.0, 0.0, 0.0

	width, height, paintArea = 4.2, 3.0, 10.0
	area := width * height
	Println(area/paintArea, "litters needed.")

	width, height = 5.2, 3.5
	area = width * height
	Println(area/paintArea, "litters needed.")
	// This doesn't look good. Too much accurate to print value

	// We can manage this by fmt.Printf()
	Printf("%.2f litters needed\n", area/paintArea) // looks good.
	// or we can save this as string variable by fmt.Sprintf()
	resultString := fmt.Sprintf("%.2f litters needed", area/paintArea)
	Println(resultString)

	// Formatting Verbs.
	// Printf doesn't have lining options. So "\n" is needed.
	Printf("A float : %f\n", 3.1415)
	Printf("An Integer : %d\n", 15)
	Printf("A string : %s\n", "Gopher")
	Printf("A boolean : %t\n", true)
	Printf("Values : %v %v %v\n", 1.22, "\t", true)    // %v automatically set format.
	Printf("Values : %#v %#v %#v\n", 1.22, "\t", true) // %#v print just same as go code
	Printf("Types : %T %T %T\n", 1.22, "\t", true)
	Printf("Percent sign : %%\n\n")

	// Print pretty format by setting value width.
	Printf("%12s | %s\n", "Product", "Cost in Cents")
	Println("-----------------------------")
	Printf("%12s | %2d\n", "Stamps", 50)
	Printf("%12s | %2d\n", "Paper Clips", 5)
	Printf("%12s | %2d\n\n", "Tape", 99)

	// if size is exceeded by float, it is ignored.
	Printf("%12s | %s\n", "Product", "Cost in Bills")
	Println("-----------------------------")
	Printf("%12s | %5.2f\n", "Americano", 2.0)
	Printf("%12s | %5.3f\n", "Caffe Latte", 2.5)
	Printf("%12s | %5.4f\n", "Cold Brewed", 2.8)
}

// Basic FormatVerb is not quite different from Python, ...
// But %v, %#v, %T are interesting.
