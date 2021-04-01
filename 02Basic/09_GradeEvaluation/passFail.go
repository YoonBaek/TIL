// Author : Github YoonBaek
// This program evaluate pass or fail based on grade.
// Grammar in here : if conditions, block, scope, short declare limits, type, methods... ...
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter a grade :")      // Tell users to enter a grade
	reader := bufio.NewReader(os.Stdin) // make buffer reader to read text from keyboard.

	/*
		reader.ReadString() returns two values.
		So, you have to declare two variables, string and error.
		But, Golang forces to use all declared variables.
		So, error should also be used.

		If we don't want to use all variables,
		We can ignore it by naming it underbar '_'

		input, _ := reader.ReadString('\n')

		But, We want to see the error, So we can do this.
	*/

	input, err := reader.ReadString('\n') // Return everything before Enter key ('\n')
	// '\n' is included in input variable.
	/*
		Why error is err, not error?
		error is a "type" in golang.
		If you use error as a name of a variable,
		Shadowing can be occured.
		It's not good. Because once you declare 'error',
		It's not anymore a type in your code.
		So you can not use type error.
	*/
	if err != nil {
		log.Fatal(err)
		// log.Fatal ends program. so we should make program not to be ended if there's no error.
	}
	fmt.Println("Your grade :", input)
	// But problem is...
	// input is 'string'
	// 1) remove space strings.
	input = strings.TrimSpace(input) // strings.TrimSpace removes space strings.
	// 2) parse strings to float
	grade, err := strconv.ParseFloat(input, 64) // 64 means float64
	if err != nil {
		log.Fatal(err)
	}

	/*
		Block & Scope.
		Scope belongs to Block.
		Each variable in Each block has scope.
		It's kind of focus range of a variable.
		if grade >= 60 {
			status := "passing"
		} else {
			status := "failing"
		}
		in this case, scope of variable "status" is just in if block.
		Because we declared in if Block.
		So we can't use it outer block.
		If you want to use it outer block, we can use it simply by declaring it outer block.
	*/
	var status string // belongs to func main block. so it's scope is func main.
	if grade >= 60 {
		status = "PASS"
	} else {
		status = "FAILED"
	}
	// Finally, Results.
	fmt.Println("A grade of", grade, "is", status)

	// if you want to see the result of ifPractice, codeMagnets remove comments.
	// ifPractice()
	// code Magnet quiz.
	// codeMagnet()
}
