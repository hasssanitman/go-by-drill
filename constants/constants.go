// Go supports constants of character, string, boolean, and numeric values

package main

import (
	"fmt"
	"math"
)

// const declares a constant value.
const s string = "My Constant"

func main() {
	fmt.Println(s)

	// A const statement can appear anywhere a var statement can.
	const n = 500_000_000

	// Constant expressions perform arithmetic with arbitrary
	// precision.
	const d = 3e20 / n
	fmt.Println(d)

	// A numeric constant has no type until it’s given one, such as by
	// an explicit conversion.
	fmt.Println(int64(d))

	// A number can be given a type by using it in a context that
	// requires one, such as a variable assignment or function call. For
	// example, here math.Sin expects a float64.
	fmt.Println(math.Sin(n))

}

// $ go run constants/constants.go
// My Constant
// 6e+11
// 600000000000
// -0.28470407323754404
