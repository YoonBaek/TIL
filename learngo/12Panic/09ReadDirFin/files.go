// Learned by Github YoonBaek

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func reportPanic() {
	// call recover and save return.
	p := recover()
	// No panic?
	if p == nil {
		// do nothing
		return
	} else { // if panic,
		err, ok := p.(error)
		if ok {
			fmt.Println(err)
		} else {
			// 예상치 못한 패닉을 다루기 위한 일반적인 방법은
			// 새로운 패닉을 발생시키는 것입니다.
			// 새로 발생시킨 패닉을 통해 프로그램의 실패 이유를 알 수 있습니다.
			panic(err)
		}
	}
}

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
	// recover를 처리하는 함수는 panic이 발생하는 코드 전에 와야합니다.
	defer reportPanic()
	scanDirectory("..")
	// 가상의 다른 이슈의 패닉을 만들어 기존 패닉 처리 외의 방식을 적용해봅니다.
	panic("some other issues.")
}
