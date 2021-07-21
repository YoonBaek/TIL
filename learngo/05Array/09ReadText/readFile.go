// Author : Github YoonBaek

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// open and read file
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	// make scanner
	scanner := bufio.NewScanner(file)

	// Read line by line till scanner.Scan() returns false
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	// close file and returns resources
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}
