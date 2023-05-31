// Command-line arguments are a common way to parameterize
// execution of programs. For example, go run hello.go uses
// run and hello.go arguments to the go program.

package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Args provides access to raw command-line arguments. Note
	// that the first value in this slice is the path to the program, and
	// os.Args[1:] holds the arguments to the program.
	argsWithProg := os.Args
	argsWitoutProg := os.Args[1:]

	// You can get individual args with normal indexing.
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWitoutProg)
	fmt.Println(arg)

}

// To experiment with command-line arguments it’s best to build
// a binary with go build first.

// $ go build -o cla command_line_arguments/command_line_arguments.go

// $ ./cla a b c d
// [./cla a b c d]
// [a b c d]
// c

// Next we’ll look at more advanced command-line processing
// with flags.
