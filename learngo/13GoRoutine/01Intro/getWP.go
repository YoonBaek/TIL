// Learned by Github YoonBaek
// 웹 페이지 로딩 속도를 위해 웹페이지의 크기를 바이트 단위로 측정하는 상황을 가정

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func responseSize(url string) {
	// 가져올 페이지의 URL로 http.Get함수를 호출합니다.
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	// 함수 종료시 네트워크 연결을 해제합니다.
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	// byte 슬라이스의 크기는 페이지의 크기와 같습니다.
	fmt.Printf("Getting %s...\n", url)
	fmt.Printf("Size : %d\n", len(body))
}

func main() {
	urls := []string{"https://www.naver.com",
		"https://www.daum.net",
		"https://www.google.com"}

	for _, url := range urls {
		responseSize(url)
	}
}

// Go Routine을 학습하며 써먹을 코드의 틀이 잡혔습니다.
