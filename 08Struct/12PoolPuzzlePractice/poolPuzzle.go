// Solved by Github YoonBaek

package main

import (
	"fmt"

	"github.com/YoonBaek/learngo/08Struct/12PoolPuzzlePractice/geo"
)

func main() {
	location := geo.Coordinates{Latitude: 37.42, Longitude: -122.08}
	fmt.Println("Lat :", location.Latitude)
	fmt.Println("Lon :", location.Longitude)
}
