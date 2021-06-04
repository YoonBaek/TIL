// Learned by Github YoonBaek
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// 이제 복잡하게 함수가 error를 반환하지 않아도 됩니다.
func scanDirectory(path string) {
	fmt.Println(path)
	files, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			scanDirectory(filePath)
		} else {
			fmt.Println(filePath)
		}
	}
}

func main() {
	scanDirectory("12Panic")
}
