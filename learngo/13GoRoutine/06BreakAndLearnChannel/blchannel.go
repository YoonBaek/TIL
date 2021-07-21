// learned by Github YoonBaek

package main

import (
	"fmt"
	"time"
)

type ts struct {
	Greetng string
	When    time.Time
}

func greeting(myTs chan ts) {
	myTs <- ts{Greetng: "Halo", When: time.Now()}
	time.Sleep(time.Second * 3)
	myTs <- ts{Greetng: "Halo", When: time.Now()}

}

func welcome(myTs chan ts) {
	myTs <- ts{Greetng: "Wilkommen", When: time.Now()}
	myTs <- ts{Greetng: "Wilkommen", When: time.Now()}
	// time.Sleep(time.Second * 5)
}

func main() {
	myTS := make(chan ts)
	go greeting(myTS)
	go welcome(myTS)
	// 위를 풀어도 정상 작동하지 않습니다.
	fmt.Println(<-myTS)
	fmt.Println(<-myTS)
	fmt.Println(<-myTS)
	fmt.Println(<-myTS)
	// greeting 채널의 송신을 기다리다가 deadlock이 발생합니다.
}
