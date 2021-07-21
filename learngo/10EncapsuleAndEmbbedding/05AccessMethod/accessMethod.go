// Learned by Github YoonBaek
// 접근자 메서드 : 숨겨진 내부 필드에 접근할 수 있도록 만드는 메서드!
// 접근자 메서드 컨벤션
// 1. 외부에서도 접근할 수 있도록 대문자로
// 2. 내부에 있는 값과 동일한 이름
// 3. 상관 없으나 Setter Method와 컨벤션을 맞추기 위해 포인터로 type을 전달
// 4. Get 컨벤션을 쓰지 않는다. Go는 접근자 메서드에서 Get을 사용하지 않고 필드 이름 그대로 접근 합니다.

// Access 메서드는 calendar 패키지에 구현해 놓았습니다.
package main

import (
	"fmt"
	"log"

	calendar "github.com/YoonBaek/learngo/10EncapsuleAndEmbbedding/04Calendar"
)

func main() {
	date := calendar.Date{}

	err := date.SetYear(2021)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetMonth(5)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetDay(21)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(date)
}

/*
지금까지 했던 필드를 숨기는 것과 같은 과정을 "캡슐화"라고 합니다.
캡슐화란 프로그램의 어느 한 영역에 있는 데이터를 다른 코드로부터 숨기는 것을 말합니다.
캡슐화는 지금까지 본 것 처럼 잘못된 데이터로부터 코드를 보호하는데 사용할 수 있기 때문에 중요하게 다뤄집니다.
또한 데이터에 직접 접근할 수 없기 때문에 캡슐화된 영역을 수정할 때 다른 코드에 미치는 영향에 대해서도 걱정할 필요가 없습니다.

다른 프로그래밍 언어와의 차이점은 다른 프로그래밍 언어는 클래스 단위로 캡슐화를 한다면, GOlang은 패키지 단위로도 가능하다는 것입니다.
일반적으로 Go 컨벤션에서는, 데이터 유효성 검증이 필요 없다고 판단되면 필드에 대한 직접 접근을 허용하는 편입니다.
*/
