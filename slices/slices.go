// Slices are an important data type in Go, giving a more powerful
// interface to sequences than arrays.

package main

import "fmt"

func main() {
	// Unlike arrays, slices are typed only by the elements they contain
	// (not the number of elements). To create an empty slice with
	// non-zero length, use the builtin make. Here we make a slice of
	// strings of length 3 (initially zero-valued).

	s := make([]string, 3)
	fmt.Println("Empty: ", s)

	// We can set and get just like with arrays.
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set: ", s)
	fmt.Println("get: ", s[2])

	// len returns the length of the slice as expected.
	fmt.Println("length: ", len(s))

	// In addition to these basic operations, slices support several more
	// that make them richer than arrays. One is the builtin append,
	// which returns a slice containing one or more new values. Note
	// that we need to accept a return value from append as we may
	// get a new slice value.
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("Append: ", s)

	// Slices can also be copy’d. Here we create an empty slice c of
	// the same length as s and copy into c from s.
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy: ", c)

	// Slices support a “slice” operator with the syntax
	// slice[low:high]. For example, this gets a slice of the
	// elements s[2], s[3], and s[4].
	l := s[2:5]
	fmt.Println("slice1: ", l)

	// This slices up to (but excluding) s[5].
	l = s[:5]
	fmt.Println("slice2: ", l)

	// And this slices up from (and including) s[2].
	l = s[2:]
	fmt.Println("slice3: ", l)

	// We can declare and initialize a variable for slice in a single line
	// as well.
	t := []string{"g", "h", "i"}
	fmt.Println("dcl: ", t)

	// Slices can be composed into multi-dimensional data structures.
	// The length of the inner slices can vary, unlike with multi-
	// dimensional arrays.
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerln := i + 1
		twoD[i] = make([]int, innerln)
		for j := 0; j < innerln; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

}

// Note that while slices are different types than arrays, they are
// rendered similarly by fmt.Println.
// $ go run slices/slices.go
// Empty:  [  ]
// set:  [a b c]
// get:  c
// length:  3
// Append:  [a b c d e f]
// copy:  [a b c d e f]
// slice1:  [c d e]
// slice2:  [a b c d e]
// slice3:  [c d e f]
// dcl:  [g h i]
// 2d:  [[0] [1 2] [2 3 4]]
