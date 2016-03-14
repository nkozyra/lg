package main

import (
	"fmt"
)

var (
	first string
)

func access() {
	fmt.Println(first)
	// fmt.Println(second)
	// fmt.Println(third)
	// The above two lines will fail due to variable scope
}

func main() {
	//  This is global because it's declared in the var() section
	first = "Hello"
	fmt.Println("`first`:", first)

	/*
			   This is not global because it's bound to the scope of main().
		     Note that we're explicit with our type here - this is optional in many cases
	*/
	var second string = "World"
	fmt.Println("`second`:", second)

	/*
	   Like `second`, this variable is lexically scoped wthin main().
	   Here we're relying on the Go compiler for type inference
	*/

	third := "Hello, World!"
	fmt.Println("`third`:", third)

	// access()
	// uncomment the line above to play with the scope of the variables
}
