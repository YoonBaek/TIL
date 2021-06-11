package resp

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func ResponseSize(url string) {
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
