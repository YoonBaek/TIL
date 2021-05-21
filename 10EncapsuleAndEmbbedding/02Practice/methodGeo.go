// Solved by Github YoonBaek

package main

import "fmt"

// this is coppied from "github.com/YoonBaek/learngo/08Struct/12PoolPuzzlePractice/geo"
// 메서드 만드는 문제인데 편의를 위해 그냥 가져와서 여기서 푼다.
type Coordinates struct {
	Latitude  float64
	Longitude float64
}

// 잊지말자 포인터 전달

func (c *Coordinates) SetLat(f float64) {
	c.Latitude = f
}

func (c *Coordinates) SetLon(f float64) {
	c.Longitude = f
}

func main() {
	coordinates := Coordinates{}
	coordinates.SetLat(37.42)
	coordinates.SetLon(-122.08)

	fmt.Println(coordinates)
}
