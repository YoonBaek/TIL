// Author : Github YoonBaek
package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// This function get input as a float from keyboard
func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}

/*
We're going to use this function in passfail.go
*/
