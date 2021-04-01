// Solved by Github Yoon Baek
// Q : https://www.acmicpc.net/problem/1001
package main

import . "fmt"

func main() {
	// buf way
	/*
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		var result int
		for _, value := range strings.Split(input, " ") {
			num, _ := strconv.Atoi(value)
			result -= num
		}
	*/

	// fmt.Scan
	a, b := 0, 0
	Scan(&a, &b)
	Println(a - b)
}
