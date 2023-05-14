// Go has various value types including strings, integers, floats, booleans, etc. Here are a few basic examples.

// Strings, which can be added together with +.

// Integers and floats.

// Booleans, with boolean operators as you’d expect.

package main

import "fmt"

func main() {
	fmt.Println("go by " + "DRILLL!")
	fmt.Println("100 + 200 = ", 100+200)
	fmt.Println("7.0/ 3.0 = ", 7.0/3.0)
	fmt.Println("10*10 = ", 10*10)

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
