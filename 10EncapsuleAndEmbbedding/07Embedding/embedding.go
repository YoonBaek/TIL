// Learned by Github YoonBaek

package main

import (
	"fmt"
	"log"

	calendar "github.com/YoonBaek/learngo/10EncapsuleAndEmbbedding/04Calendar"
)

func main() {
	event := calendar.Event{}

	// 임베디드 타입의 메서드, 필드 일지라도, 외부로 노출되지 않으면 승격되지 않습니다.
	// event.month = 5 //./embedding.go:11:7: event.month undefined (type calendar.Event has no field or method month)

	// 반면 임베딩 된 타입의 메서드 중 외부로 노출되어 승격된 경우에는 정상적으로 작동합니다.
	// 임베디드 타입을 처음부터 잘 설정하는 게 중요할 듯 합니다.
	// 이는 장점입니다. 보호한 필드 값이나 기능을 외부로 승격된 뒤에도 계속 보호할 수 있기 때문입니다.
	err := event.SetYear(2021)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetMonth(5)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetDay(21)
	if err != nil {
		log.Fatal(err)
	}

	// 접근자 메서드도 마찬가지 입니다.
	// 별도의 체이닝 없이도 호출되는 것을 보면 잘 임베딩 된 것을 볼 수 있습니다.
	fmt.Println(event.Year())
	// 체이닝도 됩니다.
	fmt.Println(event.Date.Month())
	fmt.Println(event.Date.Day())

	// 승격된 메서드는 외부 타입의 메서드와 공존합니다.
	// 따라서 Event type 내의 모든 메서드는 원래 event type에 있었던 것 처럼 동일하게 사용할 수 있습니다.
	// event에서 직접 정의
	err = event.SetTitle("Snowy's birth day")
	if err != nil {
		log.Fatal(err)
	}
	// Date로부터 승격
	err = event.SetYear(2021)
	if err != nil {
		log.Fatal(err)
	}
	// Date로부터 승격
	err = event.SetMonth(5)
	if err != nil {
		log.Fatal(err)
	}
	// Date로부터 승격
	err = event.SetDay(23)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(event.Title()) // event에서 직접 정의
	fmt.Println(event.Year())  // Date로부터 승격
	fmt.Println(event.Month()) // Date로부터 승격
	fmt.Println(event.Day())   // Date로부터 승격
}
