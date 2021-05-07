// Author : Github YoonBaek

package main

import "fmt"

func main() {
	ranks := map[string]int{
		"bronze": 3,
		"silver": 2,
		"gold":   1,
	}
	for medal, rank := range ranks {
		fmt.Printf("The %s medal's rank is %d\n", medal, rank)
	}
}
