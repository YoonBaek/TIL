// Learned by Github YoonBaek

// "defer"에 대해 학습하기

// 알고리즘 문제 풀 때의 vibe로 기계적으로
// scanner 파라미터에 os.Stdin 집어 넣었다가 무한 wait...

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func OpenFile(fileName string) (*os.File, error) {
	fmt.Println("Opening", fileName)
	return os.Open(fileName)
}

func CloseFile(file *os.File) {
	fmt.Println("Closing")
	file.Close()
}

func GetFloats(fileName string) (numbers []float64, err error) {
	file, err := OpenFile(fileName)
	if err != nil {
		return nil, err
	}
	defer CloseFile(file) // 옮긴 위치
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}
	// 코드에서 확인할 수 있듯, error발생 부분이 CloseFile 앞에 있어
	// 에러가 발생하면 아래 코드는 실행하지 않고 즉시 빠져 나갈 수 있게 설계 되어 있음.
	// 그래서 프로그램이 파일을 열어놓고 운영체제 자원을 계속 점유할 우려가 있습니다.
	// 이를 방지하기 위해 프로그램이 실행되더라도, 무슨 일이 있더라도 마지막에
	// 특정함수의 호출이 보장되게 하는 것이 "defer" 입니다.

	// 이 경우에는 파일을 반드시 닫아줘야 하므로 CloseFile은 무슨 일이 있더라도 실행해야 합니다.
	// 그래서 defer를 앞에 붙여주면, GetFloats가 어떤 형태로 종료되더라도
	// 끝에는 CloseFile을 실행해 줍니다.
	// 보통 이 경우 문맥을 고려해 알아보기 쉬운 자리에 넣어놓거나, 함수 맨 위에 넣는게 일반적입니다.
	// CloseFile(file) // 원래 CloseFile 위치. 이 경우 scan에서 에러가 나면 닫히지 않습니다.
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return numbers, nil
}

func main() {
	// 파일에서 읽어온 숫자의 슬라이스를 저장합니다.
	// 파일 명은 명령줄 인자에서 가져옵니다.
	numbers, err := GetFloats(os.Args[1])
	// 에러가 발생하면 보고한 뒤 프로그램을 종료합니다.
	if err != nil {
		log.Fatal(err)
	}
	var sum float64
	// 슬라이스에 담긴 모든 숫자를 더합니다.
	for _, number := range numbers {
		sum += number
	}
	// 총합을 출력합니다.
	fmt.Printf("Sum : %.2f\n", sum)
}
