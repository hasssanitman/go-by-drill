// Our first program will print the classic “hello world” message. Here’s the full source code.

package main

import "fmt"

func main() {
	fmt.Println("Hello go by DRILLL!")
}

// HOW TO RUN/COMPILE

// To run the program, put the code in hello_world.go and use
// go run.
// $ go run hello_world/hello_world.go
// hello world

// Sometimes we’ll want to build our programs into binaries. We
// can do this using go build.
// $ go build hello_world/hello_world.go
// $ ls
// hello_world    hello_world.go

// We can then execute the built binary directly.
// $ ./hello_world
// hello world

// Now that we can run and build basic Go programs, let’s learn
// more about the language.
