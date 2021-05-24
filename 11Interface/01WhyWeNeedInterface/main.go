// Learned by Github YoonBaek

package main

import "github.com/YoonBaek/learngo/11Interface/01WhyWeNeedInterface/gadget"

func playList(device gadget.TapePlayer, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	player := gadget.TapePlayer{}
	mixtape := []string{"Belief", "Still Feel Like Your Man", "New Light"}
	playList(player, mixtape)

	// 이와 같이 playList 함수가 잘 동작하고 있습니다. 이제 이게 같은 기능을 포함한 TapeRecorder 타입 (gadget 패키지 참조)
	// 에서도 잘 동작하길 바랍니다. 결국 테이프 레코더는 녹음 기능이 있는 플레이어이니까요.
	// 하지만, playList는 TapePlayer 타입으로 첫 번째 매개변수가 선언되었기 때문에, 다른 타입을 전달하려고 하면 컴파일 에러가 발생합니다.

	recorder := gadget.TapeRecorder{}
	// playList(recorder, mixtape) // 컴파일 에러 발생
}

/*
사실 playList 함수가 실질적으로 필요로 하는 TapePlayer라는 특정 타입이 아니라,
Play와 Stop 메서드를 정의한 타입입니다. 이렇게 메서드 매개변수를 정의하면 한 가지 타입의 매개변수 밖에 처리하지 못합니다.
그래서 인터페이스가 필요합니다.
*/
