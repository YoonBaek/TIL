// Learned by Github YoonBaek
// 고루틴 사용하기
package main

import (
	"fmt"
	"time"
)

func a() {
	for i := 0; i < 50; i++ {
		fmt.Print("a")
	}
}

func b() {
	for i := 0; i < 50; i++ {
		fmt.Print("b")
	}
}

func main() {
	go a()
	go b()
	// main 고루틴이 종료되면 다른 고루틴이 종료되지 않아도
	// 고루틴이 종료되게 설계되어 있습니다.
	// 따라서 다른 고루틴이 종료될 때까지 기다려줘야합니다.
	time.Sleep(time.Millisecond * 2)
	defer fmt.Println("end of main()")
}

// aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbbbbbbbbbbaaaaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbbbbbbend of main()
// 섞어서 출력되는 것을 볼 때 동시에 a()가 실행되는 중 b()도 실행됐다는 것을 알 수 있습니다.
