// Author : Github Yoon Baek
// Import and use custom package!
package main

import (
	"fmt"
	"log"

	"github.com/YoonBaek/learngo/04Package/01MakePackage/keyboard"
)

func main() {
	fmt.Println("Enter a grade : ")
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}

	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println("A grade of", grade, "is", status)
}
