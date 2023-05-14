// In Go, variables are explicitly declared and used by the
// compiler to e.g. check type-correctness of function calls.

package main

import "fmt"

func main() {

	// var declares 1 or more variables.
	var a = "initialize"
	fmt.Println(a)

	// You can declare multiple variables at once.
	var b, c int = 10, 20
	fmt.Println(b, c)

	// Go will infer the type of initialized variables.
	var d = true
	fmt.Println(d)

	// Variables declared without a corresponding initialization are
	//  zero-valued. For example, the zero value for an int is 0.
	var e int
	fmt.Println(e)

	// The := syntax is shorthand for declaring and initializing a
	// variable, e.g. for var f string = "apple" in this case. This
	//  syntax is only available inside functions.
	f := "Pineapple"
	fmt.Println(f)

}

// $ go run variables/variables.go
// initialize
// 10 20
// true
// 0
// Pineapple
