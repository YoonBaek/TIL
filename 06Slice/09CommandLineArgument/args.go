// Author : Github YoonBaek
// data.txt 파일을 매번 수정하는 것은 번거롭습니다.
// 그렇다면 명령줄에서 인자를 입력해서 가져오는 것은 어떨까요?

package main

import (
	"fmt"
	"os"
)

func main() {
	// This Returns file name and command line arguments
	// fmt.Println(os.Args)
	// remove filename.
	fmt.Println(os.Args[1:])
}
