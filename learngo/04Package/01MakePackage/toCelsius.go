// Author : Github YoonBaek
package main

import (
	"fmt"
	"log"

	"github.com/YoonBaek/learngo/04Package/01MakePackage/keyboard"
)

func main() {
	fmt.Println("Enter a temperature in Fahrenheit :")
	fahr, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	celsius := (fahr - 32) * 5 / 9
	fmt.Printf("%0.2f degrees Celsius\n", celsius)
}
