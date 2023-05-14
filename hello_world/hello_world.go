// Our first program will print the classic “hello world” message. Here’s the full source code.

// To run the program, put the code in hello_world.go and use go run.

// Sometimes we’ll want to build our programs into binaries. We can do this using go build.

// We can then execute the built binary directly.

// Now that we can run and build basic Go programs, let’s learn more about the language.

package main

import "fmt"

func main() {
	fmt.Println("Hello go by DRILLL!")
}

// HOW TO RUN/COMPILE

// $ go run hello_world/hello_world.go
// hello world

// $ go build hello_world/hello_world.go
// $ ls
// hello_world    hello_world.go

// $ ./hello_world
// hello world
