// Solved by Github YoonBaek

package main

import "fmt"

type Population int

func main() {
	var population Population = Population(572)

	fmt.Println("Sleepy Creek County population :", population)
	fmt.Println("Congratulation, Kevin and Anna, It's a girl!")
	population += 1
	fmt.Println("Sleepy Creek County population :", population)
}
