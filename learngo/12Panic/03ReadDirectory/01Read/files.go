package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	files, err := os.ReadDir("mydir")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			fmt.Println("Directory :", file.Name())
		} else {
			fmt.Println("File :", file.Name())
		}
	}
}
