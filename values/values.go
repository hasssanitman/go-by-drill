// Go has various value types including strings, integers, floats, booleans, etc. Here are a few basic examples.

package main

import "fmt"

func main() {

	// Strings, which can be added together with +.
	fmt.Println("go by " + "DRILLL!")

	// Integers and floats.
	fmt.Println("100 + 200 = ", 100+200)
	fmt.Println("7.0/ 3.0 = ", 7.0/3.0)
	fmt.Println("10*10 = ", 10*10)

	// Booleans, with boolean operators as you’d expect.
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}

// $ go run values/values.go
// go by DRILLL!
// 100 + 200 =  300
// 7.0/ 3.0 =  2.3333333333333335
// 10*10 =  100
// false
// true
// false
