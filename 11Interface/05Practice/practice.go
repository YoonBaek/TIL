// Solved by Github YoonBaek
/*
	아래 코드에서는 Car와 Truck이라는 타입을 정의하며 각 타입은 Accelerate, Brake, Steer 메서드를 가지고 있습니다.
	빈 칸을 채워 이 메서드를 가진 Vehicle 인터페이스를 추가해 main 함수의 코드가 아래 보이는 값을 출력하도록 만들어 주세요.

	Speeding up
	Turning left
	Stopping
	Turning right
*/

package main

import "fmt"

type Car string
type Truck string

func (c Car) Accelerate() {
	fmt.Println("Speeding up")
}
func (c Car) Brake() {
	fmt.Println("Stopping")
}
func (c Car) Steer(direction string) {
	fmt.Println("Turning", direction)
}

func (t Truck) Accelerate() {
	fmt.Println("Speeding up")
}
func (t Truck) Brake() {
	fmt.Println("Stopping")
}
func (t Truck) Steer(direction string) {
	fmt.Println("speeding up")
}
func (t Truck) LoadCargo(cargo int) {
	fmt.Println("Loading", cargo)
}

type Vehicle interface {
	Accelerate()
	Brake()
	Steer(string)
}

func main() {
	var vehc Vehicle = Car("Porsche 911")
	vehc.Accelerate()
	vehc.Steer("left")

	vehc = Truck("Ford F180")
	vehc.Brake()
	vehc.Steer("right")
}
