// Learned by Github YoonBaek
// 웹 페이지 로딩 속도를 위해 웹페이지의 크기를 바이트 단위로 측정하는 상황을 가정

package main

import (
	"time"

	"github.com/YoonBaek/learngo/13GoRoutine/03ApplyingGoRoutine/resp"
)

func main() {
	urls := []string{"https://www.naver.com",
		"https://www.daum.net",
		"https://www.google.com"}

	for _, url := range urls {
		go resp.ResponseSize(url)
	}
	time.Sleep(time.Second)
}

// 함수 앞에 Go만 붙였을 뿐인데 세 함수가 동시에 실행됩니다!
// 실제로 실행해보면 훨씬 빠름을 알 수 있습니다.
// 더 이상 선행 함수 작업이 끝날 때 까지 기다릴 필요가 없기 때문입니다.
// 하지만 호출이 실행되는 순서는 제어할 수 없기 때문에
// 실행할 때마다 요청의 실행 순서가 바뀌는 것도 확인할 수 있습니다.
