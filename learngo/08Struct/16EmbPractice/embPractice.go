// Solved by Github Yoon Baek
package main

import (
	"fmt"

	"github.com/YoonBaek/learngo/08Struct/12PoolPuzzlePractice/geo"
)

func main() {
	location := geo.Landmark{Name: "The Googleplex"}
	location.Latitude = 37.42
	location.Longitude = -122.08

	fmt.Println(location)
}

// Struct Done.
