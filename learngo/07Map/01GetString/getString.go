// Author : Github Yoon Baek
// 파일에서 문자열을 읽어오는 함수.

package getString

import (
	"bufio"
	"log"
	"os"
)

func GetStrings(fileName string) (lines []string, err error) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return lines, nil
}
