// Learned by Github YoonBaek

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Page struct {
	Url  string
	Size int
}

func responseSize(url string, channel chan int) {
	fmt.Println("Getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	channel <- len(body)
}

func responsePage(url string, channel chan Page) {
	fmt.Println("Getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	channel <- Page{Url: url, Size: len(body)}
}

func main() {
	sizes := make(chan int)

	urls := []string{"https://www.naver.com",
		"https://www.daum.net",
		"https://www.google.com",
		"https://www.github.com"}

	for _, url := range urls {
		go responseSize(url, sizes)
		// fmt.Println(<-sizes) 이곳에 넣으면 다음 고루틴 이전에 메인 고루틴을 반드시 실행해야 하기 때문에 한 번에 하나씩만 가져옵니다.
	}
	// 송신 고루틴과 수신 고루틴을 분리해줘야 메인 루틴에서 한 번에 합쳐줍니다.
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-sizes)
	}
	// 그런데 이렇게 하면 가져온 페이지가 어떤 페이지인지 알기 어렵습니다.
	// 그래서 둘을 매치시켜주는 구조체 채널을 만들 수도 있습니다.
	pages := make(chan Page)
	for _, url := range urls {
		go responsePage(url, pages)
	}
	for i := 0; i < len(urls); i++ {
		page := <-pages
		fmt.Printf("Size of %s : %d\n", page.Url, page.Size)
	}
	// 이제 url까지 깔끔하게 알 수 있게 됐으며, 고루틴으로 다음 요청을 동시에 처리했기 때문에 3배 더 빠른 시간 안에 작업을 끝냈습니다.
}
