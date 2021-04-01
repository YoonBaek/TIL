// Author : Github Yoon Baek
// code magnet quiz

package main

import (
	"fmt"
	"log"
	"os"
)

func codeMagnet() {
	fileInfo, err := os.Stat("passFail.go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("size of passFail.go", fileInfo.Size()) // return size of the file.
}
