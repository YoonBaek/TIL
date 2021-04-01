// Solved by Github Yoon Baek
// Q : https://www.acmicpc.net/problem/1000
package main

import "fmt"

func main() {
	// buf way
	/*
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		var result int
		for _, value := range strings.Split(input, " ") {
			num, _ := strconv.Atoi(value)
			result += num
		}
	*/

	// fmt.Scan
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(a + b)
}
