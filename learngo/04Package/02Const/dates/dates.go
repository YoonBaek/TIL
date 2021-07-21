// Author : Github YoonBaek
// Const is for the stability of the code.
// Once declared, it never changes.
package dates

// This is the only way to declare const.
// No short way, no later assign.
const DaysInWeek int = 7

// Most of const is assigned at package level.
// If it's declared in func, it's usable just in the function.

func WeeksToDays(weeks int) int {
	return weeks * DaysInWeek
}

func DaysToWeeks(days int) float64 {
	return float64(days) / float64(DaysInWeek)
}
