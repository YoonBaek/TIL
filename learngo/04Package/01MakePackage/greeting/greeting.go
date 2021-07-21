// Author : Github YoonBaek
// Let's make package
// packagg also has same structure.
// package - import - function.
package greeting

// this package name should be same as folders name.
// If not, You can't import package as package name.
// it goes like this.
// import notsamepackagename "samepackagename"

import . "fmt"

// To make this functions accessable outside,
// Uppercase letter is used.
func Hello() {
	Println("Hello!")
}

func Hi() {
	Println("Hi!")
}
