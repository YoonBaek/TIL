package main

import (
	"fmt"

	magazine "github.com/YoonBaek/learngo/08Struct/10ExportStruct"
)

func main() {
	// var sub0 magazine.subscriber
	// subscriber not exported by package magazine

	// 대문자는 불러올 수 있습니다.
	// 타입명은 magazine.Subscriber로 써야합니다.
	var sub1 magazine.Subscriber
	fmt.Printf("%#v\n", sub1)

	// sub1.name = "Baek"
	// 그런데 field는 불러올 수 없습니다.
	// field도 마찬가지로 대문자 처리해서 export 가능하게 바꿔줘야 합니다.

	var sub2 magazine.SubScriber
	sub2.Rate = 4.99
	fmt.Printf("%#v\n", sub2)
	// 이제 struct 내부에 field값을 할당할 수 있습니다.
}
