// Learned by Github YoonBaek
// 내부 필드에 접근할 때 매번 이름을 통해 접근하는 것은 번거로울 수 있습니다.
// 따라서 Golang에서는 필드의 이름을 따로 지정하지 않고 타입만 Struct로 제공할 수 있는
// 익명 필드 annonymous field 기능을 제공합니다.

package main

import (
	"fmt"

	magazine "github.com/YoonBaek/learngo/08Struct/10ExportStruct"
)

type AnnonyTest struct {
	name             string
	rate             float64
	magazine.Address // 이게 anonymous field기능입니다.
}

func main() {
	subscriber := AnnonyTest{name: "Yoon Baek"}
	subscriber.rate = 0.99
	// Anonymous field 기능을 사용하면 아래와 같이 필드명을 생략해서 가져올 수도 있습니다.
	// 이걸 Embedding이라고 합니다.
	subscriber.City = "Seoul"
	// 정상적으로 Address.City에 들어가 있는 것을 볼 수 있습니다.
	fmt.Printf("%#v\n", subscriber)
	fmt.Println("Name :", subscriber.name)
	fmt.Println("City :", subscriber.City)
}
