// Learned by Github Yoon Baek

package main

import "fmt"

type (
	Milliliters float64
	Liters      float64
	Gallons     float64
)

// 리시버 매개변수는 다른 언어의 self나 this 대신입니다.
func (l Liters) ToGallon() Gallons {
	return Gallons(l * 0.264)
}

func (m Milliliters) ToGallon() Gallons {
	return Gallons(m * 0.000264)
}

func (g Gallons) ToLiters() Liters {
	return Liters(g * 3.7854)
}

func (g Gallons) ToMilliliters() Milliliters {
	return Milliliters(g * 3785.41)
}

func main() {
	soda := Liters(2)
	fmt.Printf("%.2f liters of soda equals to %.2f gallons\n", soda, soda.ToGallon())
	water := Milliliters(500)
	fmt.Printf("%.2f milliliters of water equals to %.2f gallons\n", water, water.ToGallon())
	milk := Gallons(2)
	fmt.Printf("%.2f gallons of milk equals to %.2f liters\n", milk, milk.ToLiters())
	fmt.Printf("%.2f gallons of milk equals to %.2f milliliters\n", milk, milk.ToMilliliters())
}
