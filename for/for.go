// for is Go’s only looping construct. Here are some basic types of for loops.

package main

import "fmt"

func main() {

	// The most basic type, with a single condition
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	// A classic initial/condition/after for loop.
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// for without a condition will loop repeatedly until you break
	// out of the loop or return from the enclosing function.
	for {
		fmt.Println("inifinit loop if you don't break")
		break
	}

	// You can also continue to the next iteration of the loop.
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

// $ go run for/for.go
// 1
// 2
// 3
// 7
// 8
// 9
// inifinit loop if you don't break
// 1
// 3
// 5

// We’ll see some other for forms later when we look at range
// statements, channels, and other data structures.
