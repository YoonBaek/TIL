// Learned by Github YoonBaek
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// 간단한 재귀함수 학습 코드
func count(start int, end int) {
	fmt.Printf("count(%d, %d) called\n", start, end)
	if start < end {
		count(start+1, end)
	}
	// 재귀함수에서 어떤 함수가 먼저 끝나는지   알 수 있습니다.
	fmt.Printf("Returning from count(%d, %d)\n", start, end)
}

func ScanDirectory(path string) error {
	// print current directory
	fmt.Println(path)
	// get dir content slice
	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, file := range files {
		// join path and current file name.
		filePath := filepath.Join(path, file.Name())
		// if subdir,
		if file.IsDir() {
			// recursive call by filePath
			err := ScanDirectory(filePath)
			if err != nil {
				return err
			}
		} else { // if not,
			// print path
			fmt.Println(filePath)
		}
	}
	return nil
}

func main() {
	count(0, 3)
	err := ScanDirectory("learngo")
	if err != nil {
		log.Fatal(err)
	}
}
