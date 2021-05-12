// Learned by Github YoonBaek
// Structure Embedding

/*
잘 이해가 안되니 다 써 가며 이해하기...

익명 필드는 구조체 정의에서 필드명을 생략할 수 있다는 것 이외에도 많은 것들을 제공합니다.

외부 Structure의 익명 필드로 선언된 내부 Structure를 외부 스트럭쳐 안에 Embedding되었다고 합니다.
임베딩된 Structure의 필드는 외부 스트럭쳐로 promoted되는데 승격되었다는 말은 내부 스트럭쳐의 필드를
마치 외부 스트럭쳐에 속해 있는 것처럼 접근할 수 있음을 의미합니다.

이제 Address스트럭쳐 타입은 subscriber와 employee 스트럭쳐타입에 임베딩 되었기 때문에
subscriber.Address.Citu대신 subscriber.City만으로 City필드의 값을 가져올 수 있습니다.
마찬가지로 employee.Address.State대신 employee.State만으로 State필드의 값을 가져올 수 있습니다.

다음은 Addres를 임베딩 타입으로 다루는 main.go의 최종 버전입니다. 이제 Address의 모든 필드가 자기자신이 임베딩된
외부 스트럭쳐에 속한 것 마냥 동작하기 때문에 마치 Address타입이 없는 듯한 코드를 작성할 수 있습니다.
*/
package main

import (
	"fmt"

	magazine "github.com/YoonBaek/learngo/08Struct/10ExportStruct"
)

type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
	magazine.Address
}

func main() {
	subscriber := Subscriber{Name: "Yoon Baek"}
	// Address의 필드가 마치 Subscriber 필드에 정의된 것 처럼 필드값을 할당하고 있습니다.
	subscriber.Street = "강남대로"
	subscriber.City = "서울시"
	subscriber.State = "KR"
	subscriber.PostalCode = "12345"
	// 할당 뿐만 아니라 가져오는 것도 마찬가지 입니다.
	fmt.Println("Street :", subscriber.Street)
	fmt.Println("City :", subscriber.City)
	fmt.Println("State :", subscriber.State)
	fmt.Println("Postal Code :", subscriber.PostalCode)
}

/* 여기서 명심할 것은 내부 스트럭쳐를 반드시 임베딩할 필요는 없다는 것입니다.
또한 항상 내부 스트럭쳐를 사용할 필요도 없습니다.
때로는 내부 스트럭쳐가 아닌 외부 스트럭쳐에 필드를 추가하는 것이 가장 깔끔한 방법일 수도 있습니다.
상황에 따라 가장 적절한 방법을 사용하는게 중요합니다. */
