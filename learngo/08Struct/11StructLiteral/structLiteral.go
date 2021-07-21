// Learned by Github YoonBaek
// 구조체 리터럴
// Struct를 선언한 뒤 필드에 값을 하나씩 할당해주는 것은 따분한 일입니다.
// Golang은 슬라이스나 맵과 마찬가지로 Struct Literal을 지원하며 덕분에
// Struct 생성과 동시에 초깃값을 할당할 수 있습니다.
package main

import (
	"fmt"

	magazine "github.com/YoonBaek/learngo/08Struct/10ExportStruct"
)

func main() {
	// How To Declare
	// car 라는 struct 타입이 있다는 가정하에,
	// 타입명{필드명:"필드값"}이라는 형식으로 저장할 수 있습니다.
	// Ferrari := car{name: "Ferrari Rome",
	// 	topSpeed: 320}

	// 지금까지는 struct를 선언할 때 필드값도 하나하나 선언해서
	// 다소 긴 코드로 작성해야했지만 아래와 같이 리터럴을 활용하면
	// 단축변수 선언을 해줄 수 있습니다.
	subscriber := magazine.SubScriber{Name: "Gopher Yoon", Rate: 4.99, Active: true}
	fmt.Printf("%#v\n", subscriber)

	// 필드 선언을 생랴할 수도 있습니다.
	// 생략하게 되면 각 필드별 타입에 대한 디폴트값이 들어갑니다.
	subscriber = magazine.SubScriber{Name: "Gopher Baek"}
	fmt.Printf("%#v\n", subscriber)
}
