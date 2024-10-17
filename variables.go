package main

import "fmt"

func main() {
	var a = "apple"
	var b, c = "banana", "cat"
	var d, e string = "dog", "elephant"
	var f int = 240
	var g int
	// var h wrong we have to define the type otherwise syntax error
	i := 10
	fmt.Println(a, b, c, d, e, f, g, i)
}

// In Go variables are explicitly declared and used by the compiler to e.g. check type-correctness of function calls.