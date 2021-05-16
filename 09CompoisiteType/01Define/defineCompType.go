// Learned by Github YoonBaek
// 보통 사용자 정의 타입은 struct를 많이 활용하지만, int, string, bool 등 모든 타입을 기본 타입으로 사용할 수 있습니다.
// 구분해야 하는 변수의 충돌을 막기 위해 주로 사용합니다.

package main

import "fmt"

type Liters float64
type Gallons float64

type Title string

func main() {
	var carFuel Gallons
	var busFuel Liters

	carFuel = Gallons(12.4)
	busFuel = Liters(240.0)

	fmt.Println(carFuel, busFuel)

	// 변환한 타입 간에는 서로 할당할 수 없습니다.
	// carFuel = Liters(1)

	// 기본 타입이 일치하는 경우에는 다른 타입으로의 변환도 가능하긴 합니다.
	carFuel = Gallons(Liters(47))
	// 그러나 이는 변환 에러를 방지하고자 하는 보호 기능을 무력화시킵니다.
	// 따라서 이런 식의 변환은 하지 않는 것이 좋습니다.

	// 1리터는 약 0.264 갤런이고,
	// 1갤런은 약 3.785 리터입니다. 따라서 변환할 때는 각각 값을 곱해주는게 맞습니다.
	carFuel = Gallons(Liters(47) * 0.264)
	busFuel = Liters(Gallons(63) * 3.785)
	fmt.Printf("Gallons : %.2f, Liters : %.2f\n", carFuel, busFuel)

	// 위에서 확인할 수 있듯, 사용자 정의 타입은 기본 정의 타입에서 사용할 수 있는 연산은 모두 사용 가능합니다.
	// 반대로 string 타입에서 빼기 연산은 안되기 때문에 string으로 만든 타입 역시 빼기 연산이 안됩니다.
	title1 := Title("Layla")
	title2 := Title("Wonderful Tonight")
	fmt.Println(title1 == title2)
	fmt.Println(title1 < title2)
	fmt.Println(title1 > title2)
	fmt.Println(title1 + " by Eric Clapton") // 여기까진 잘 됩니다.
	// fmt.Println(title2 - " Tonight") // 안됩니다.

	// 하지만 다른 타입과의 연산은 불가능하며, 설령 기본 타입이 같은 경우라도 안됩니다.
	// 이는 서로 다른 타입을 혼용할 여지를 미연에 방지해 줍니다.
	// fmt.Println(carFuel + busFuel)

	// 따라서 상호 연산이 필요한 경우에는, 해당 타입으로의 변환이 필수적입니다.
}
