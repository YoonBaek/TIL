// Solved By Github YoonBaek
// nitroBoost, doublePackage 함수가 제대로 작동하지 않는 현상 고치기

package main

import "fmt"

type car struct {
	name     string
	topSpeed float64
}

type part struct {
	description string
	count       int
}

func nitroBoost(c *car) {
	c.topSpeed += 50
}

func doublePack(p *part) {
	p.count *= 2
}

func main() {
	var ferrari car
	ferrari.name = "Ferrari Rome"
	ferrari.topSpeed = 320

	nitroBoost(&ferrari)
	fmt.Printf("%s boosts to %.2fkm/h\n", ferrari.name, ferrari.topSpeed)

	var fuses part
	fuses.description = "Fuses"
	fuses.count = 5

	doublepack(&fuses)
	fmt.Printf("%s double package has %d parts.\n", fuses.description, fuses.count)
}
