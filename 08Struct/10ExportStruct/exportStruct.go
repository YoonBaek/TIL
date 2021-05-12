// Learned By Github YoonBaek.
// Struct type export하기

package magazine

// 소문자로는 안됩니다.
type subscriber struct {
	name   string
	rate   float64
	active bool
}

// 이제 export할 수 있습니다.
type Subscriber struct {
	name   string
	rate   float64
	active bool
}

// 이제 field도 export할 수 있습니다.
type SubScriber struct {
	Name   string
	Rate   float64
	Active bool
	Home   Address
}

// 이 방식으로 노출시킬 필드와 감출 필드를 구분해줄 수 있습니다.
type Employee struct {
	Name   string
	Salary float64
	Home   Address
}

// struct를 생성해서 SubScriber와 Employee의 필드값으로 지정할 수도 있습니다.
type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}
