// Author Github YoonBaek

package main

import (
	"fmt"

	"github.com/YoonBaek/learngo/04Package/02Const/dates"
)

func main() {
	days := 3
	fmt.Println("Your appointment is in", days, "days.")
	fmt.Println("with a follow-up in", days+dates.DaysInWeek, "days") // Use constant.
}
