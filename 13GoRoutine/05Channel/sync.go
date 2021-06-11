// Learned by Github YoonBaek
// 고루틴의 동기화를 통해 예측 가능한 출력값을 획득할 수 있습니다.

package main

import (
	"fmt"
	"time"
)

func abc(channel chan string) {
	for _, str := range []string{"a", "b", "c"} {
		for i := 0; i < 20; i++ {
			channel <- str
		}
	}
}

func def(channel chan string) {
	for _, str := range []string{"d", "e", "f"} {
		for i := 0; i < 20; i++ {
			channel <- str
		}
	}
}

func nap(name string, delay int) {
	for i := 0; i < delay; i++ {
		fmt.Println(name, "sleeping for", i+1, "sec")
		time.Sleep(time.Second * 1)
	}
	fmt.Println(name, "wakes up!")
}

func send(myChannel chan string) {
	fmt.Println("***송신 어떻게 일어나는지 한 번 봐***")
	nap("sending go routine", 2)
	fmt.Println("***first sending***")
	myChannel <- "a"
	fmt.Println("***second sending***")
	myChannel <- "b"
}

func test() int {
	return 1
}

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)
	go abc(channel1)
	go def(channel2)
	// 수신 고루틴은 디른 고루틴이 값을 보내기 전까지 대기합니다.
	for i := 0; i < 60; i++ {
		fmt.Print(<-channel1)
		fmt.Print(<-channel2)
	}
	fmt.Println()

	myChan := make(chan string)
	go send(myChan)
	nap("receiving goroutine", 5)
	fmt.Println(<-myChan) // 수신 고루틴이 nap에 의해 5초 뒤로 강제됩니다.
	// 따라서 첫번째 송신이 이뤄진 3초 뒤에 두번째 송신이 일어나게 됩니다.
	fmt.Println(<-myChan)

	// for i := 0; i < 10; i++ {
	// 	go fmt.Println(test())
	// }
}

// 넣은 순서대로 출력됨...
// queue랑 똑같네?
