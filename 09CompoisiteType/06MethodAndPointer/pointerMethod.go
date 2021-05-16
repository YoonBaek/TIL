// Learned by Github YoonBaek
// 타입의 메서드에는 값 리시버와 포인터 리시버가 있습니다.
// 타입의 복사본이 아닌 타입 자체를 전달하고 싶을 때 포인터 리시버를 사용하면 됩니다.

// Important!
// 타입의 메서드 리시버를 지정할 때도 컨벤션이 있습니다.
// 가급적이면 포인터 리시버와 값 리시버의 혼용은 피하고 하나로만 메서드를 구성하는 것이 좋습니다.

package main

import "fmt"

type Number int

// value receiver
func (n Number) valueDouble() {
	n *= 2
}

// pointer receiver
func (n *Number) Double() {
	*n *= 2
}

func main() {
	num := Number(4)
	fmt.Println("original :", num)
	num.valueDouble()
	fmt.Println("after value receiver method :", num) // 값이 바뀌지 않았습니다.
	num.Double()
	fmt.Println("after calling double :", num) // 포인터를 사용한 메서드는 값이 바뀝니다.

	// 더불어 포인터는 값 자체에 생기는 것이 아니라 값이 저장된 변수에 생기는 것임을 유의해야합니다.
	// 따라서 아래와 같은 활용은 불가능합니다.
	// Number(4).Double()
	// 아래는 포인터 메서드가 아니기 때문에 가능합니다.
	Number(4).valueDouble()
}
