// Learned by Github YoonBaek

package gadget

import "fmt"

type TapePlayer struct {
	Batteries string
}

type Robot string

func (r Robot) MakeSound() {
	fmt.Println("Beep")
}

func (t TapePlayer) Play(song string) {
	fmt.Println("Playing", song)
}

func (t TapePlayer) Stop() {
	fmt.Println("Stopped")
}

type TapeRecorder struct {
	Microphones int
}

// 메서드 매개변수는 같은 기능일지라도 이처럼 변수 타입마다 메서드를 다시 선언해줘야합니다.
func (t TapeRecorder) Record() {
	fmt.Println("Recording...")
}

func (t TapeRecorder) Play(song string) {
	fmt.Println("Playing", song)
}

func (t TapeRecorder) Stop() {
	fmt.Println("Stopped")
}
