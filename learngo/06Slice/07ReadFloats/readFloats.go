// Author : Github YoonBaek
// 슬라이스와 append를 활용해 추가 데이터를 읽어올 것입니다.
package readFloats

import (
	"bufio"
	"os"
	"strconv"
)

// Get floats reads each line as float.
func GetFloats(fileName string) (numbers []float64, err error) {
	// open file as a filename
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		//convert string to float64.
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return numbers, nil
}
