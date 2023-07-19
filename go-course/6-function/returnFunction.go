package main

import "fmt"

func main() {
	// assign to x the function
	x := foo()
	fmt.Printf("%T\n", x)

	y := x()
	fmt.Println(y)
}

// return a function 
func foo() func() int {
	return func() int{
		return 42
	}
}