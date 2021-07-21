// Solved by Github YoonBaek

package main

import (
	"fmt"
	"log"

	"github.com/YoonBaek/learngo/10EncapsuleAndEmbbedding/06Practices/geo2"
)

func main() {
	// practice1
	coordinates := geo2.Coordinates{}
	err := coordinates.SetLatitude(37.42)
	if err != nil {
		log.Fatal(err)
	}

	err = coordinates.SetLongitude(-1122.08)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(coordinates.Latitude())
	fmt.Println(coordinates.Longitude())

	// practice2
	location := geo2.Landmark{}
	err = location.SetName("The Googleplex")
	if err != nil {
		log.Fatal(err)
	}
	err = location.SetLatitude(37.42)
	if err != nil {
		log.Fatal(err)
	}
	err = location.SetLongitude(-122.08)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(location.Name())
	fmt.Println(location.Latitude())
	fmt.Println(location.Longitude())
}
