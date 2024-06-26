package main

import "fmt"

// functions are central in go
// we'll learn about functions with a few different examples

// here's a function that takes two integers and returns their sum as an int
func plus(a int, b int) int {
	// Go requires explicit, i.e. it won't automatically return the value of the last expression
	return a + b
}

// when you have multiple consecutive parameters of the same type
// you may omit the type name for the like-typed parameters up to the final parameter that declares the type
func plusPlus(a, b, c int) int {
	return a + b + c
}

// call a funcion just as you'd expect, with name(args)
func main() {
	// call a function just as you'd expect, with name(args)
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	// there are several other features to Go functions
	// one is multiple return values, which we'll look at next
}
