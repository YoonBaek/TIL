// Solved by Github YoonBaek
package main

import (
	"fmt"
	"log"
)

type Refrigerator []string

func (r Refrigerator) Open() {
	fmt.Println("Opening")
}

func Find(item string, slice []string) bool {
	for _, sliceItem := range slice {
		if sliceItem == item {
			return true
		}
	}
	return false
}

func (r Refrigerator) Close() {
	fmt.Println("Closing")
}

func (r Refrigerator) FindFood(food string) error {
	r.Open()
	// Modified
	defer r.Close()
	if Find(food, r) {
		fmt.Println("Found", food)
	} else {
		return fmt.Errorf("%s not found", food)
	}
	// original
	// r.Close()
	return nil
}

func main() {
	fridge := Refrigerator{"Milk", "Pizza", "Salsa"}
	for _, food := range []string{"Milk", "Bananas"} {
		err := fridge.FindFood(food)
		if err != nil {
			log.Fatal(err)
		}
	}
}
