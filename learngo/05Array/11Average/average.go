// Author : Github YoonBaek

package main

import (
	"fmt"
	"log"

	readFloats "github.com/YoonBaek/learngo/05Array/10ReadSaveFloat"
)

func main() {
	numbers, err := readFloats.GetFloats("../09ReadText/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	sum := 0.0
	for _, number := range numbers {
		sum += number
	}
	cnt := float64(len(numbers))
	fmt.Println("Average :", sum/cnt)
}

// Next Question...
// What about data which have more than 3?
// In this code, error must occur because readFloats package only process fixed size (3) array.
// In this situatuin we can use slice.
// We will take care about slice at next chapter.
