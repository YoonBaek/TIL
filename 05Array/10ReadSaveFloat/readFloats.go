// Author : Github YoonBaek

package readFloat

import (
	"bufio"
	"os"
	"strconv"
)

// Get floats reads each line as float.
func GetFloats(fileName string) (numbers [3]float64, err error) {
	// open file as a filename
	file, err := os.Open(fileName)
	if err != nil {
		return numbers, err
	}
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		//convert string to float64.
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return numbers, err
		}
	}
	err = file.Close()
	if err != nil {
		return numbers, err
	}
	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}
	return numbers, nil
}
