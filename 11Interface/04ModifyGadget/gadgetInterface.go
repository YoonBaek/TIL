// Learned and Written by Github Yoon Baek
// 인터페이스로 파라미터의 타입간 호환성 문제 해결하기.
package main

import (
	"fmt"

	"github.com/YoonBaek/learngo/11Interface/01WhyWeNeedInterface/gadget"
)

type Player interface {
	Play(string)
	Stop()
}

func playList(p Player, songs []string) {
	for _, song := range songs {
		p.Play(song)
	}
	p.Stop()
}

// 타입이 포인터 메서드를 선언하고 있는 경우

type controller bool

func (c *controller) toggle() {
	if *c {
		*c = false
	} else {
		*c = true
	}
	fmt.Println(*c)
}

type toggleable interface {
	toggle()
}

func main() {
	album := []string{"Wonderful Tonight", "Tears In Heaven", "Layla"}
	player := gadget.TapePlayer{}
	recorder := gadget.TapeRecorder{}

	playList(player, album)
	playList(recorder, album)

	// 포인터 메서드의 경우에는 리시버를 포인터로 넘겨주기 때문에
	// 포인터 변수의 주소값만 할당해주면 문제 없습니다.
	ctrl := controller(true)
	var t toggleable = &ctrl
	t.toggle()
	t.toggle()
}
