// Learned by Github YoonBaek

package main

import (
	"fmt"
	"time"
)

type (
	Liters  float64
	Gallons float64
)

func (L Liters) ToGallons() Gallons {
	return Gallons(L * 0.264)
}

func main() {
	carFuel := Gallons(1.2)
	busFuel := Liters(2.5)
	// 타입간 연산은 할 수 없습니다. invalid operations
	// carFuel += Liters(9.0)
	// busFuel += Gallons(9.0)

	// 서로 다른 타입을 함께 연산하려면 먼저 타입 변환을 통해 타입부터 맞춰야 합니다.
	// 이전에 각 타입 간 변환했던 연산과 동일한 작업을 수행하는 ToGallons와 ToLiters를 만들어두면
	// 중복을 피하면서 필요할 때마다 불러와 사용할 수 있습니다.

	carFuel += ToGallons(Liters(40))
	busFuel += ToLiters(Gallons(30))

	fmt.Printf("Car Fuel : %.2f gallons\n", carFuel)
	fmt.Printf("Bus Fuel : %.2f liters\n", busFuel)

	// 그러나 이렇게 사용하게 되면 mililiter에서 gallon으로 변환할 때도 오버로딩이 생기는 문제가 발생하게 됩니다.
	// 그렇다고 매번 MiliLitersToGallon, LitersToGallon과 같은 함수를 생성해서 사용하는 것도 번거롭습니다.
	// 이때 사용할 수 있는 것이 메서드입니다.
	var now time.Time = time.Now()
	// time.Time 값은 연도를 반환하는 Year 메서드를 가지고 있습니다.
	var year = now.Year()
	fmt.Println(year)

	// 이것처럼 메서드를 만들어서 사용하면, ToGallons를 타입별로 여러번 정의할 수 있습니다.
	// 그리고 아래와 같은 연산도 가능해지죠.
	totalFuel := carFuel + busFuel.ToGallons()
	fmt.Println("Estimated total :", totalFuel)

	// 메서드는 리시버 매개변수를 통해 선언할 수 있는데, 다음 장에서 공부해보겠습니다.
}

func ToGallons(l Liters) Gallons {
	return Gallons(l * 0.264)
}

func ToLiters(g Gallons) Liters {
	return Liters(g * 3.785)
}
