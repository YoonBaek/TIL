// Learned by Github YoonBaek

package main

import (
	"log"
	"net/http"
)

// writer : 브라우저에 전달할 응답을 수정하기 위한 값
// request : 브라우저에서 전달 받은 요청을 나타내는 값
func viewHandler(writer http.ResponseWriter, request *http.Request) {
	// 응답에 "Hello, web!"을 추가합니다.
	write(writer, "Hello, Web!")
}

func write(writer http.ResponseWriter, message string) {
	_, err := writer.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// "/hello"로 끝나는 url에 대한 요청을 받으면,
	// viewHandler 함수를 호출하여 응답을 생성합니다.
	http.HandleFunc("/hello", viewHandler)
	// 브라우저의 요청을 수신하고 응답합니다.
	// localhost를 사용하면 브라우저는 인터넷이 아닌 자신의 컴퓨터로 네트워크 연결을 생성합니다.
	// 또한 url의 일부로 포트를 지정해줘야 합니다.
	// 포트는 어플리케이션이 메시지를 수신할 수 있는 네트워크 통신채널을 숫자로 나타낸 값입니다.
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
