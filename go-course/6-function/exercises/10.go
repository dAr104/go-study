package main

import "fmt"

/*
Closure is when we have “enclosed” the scope of a variable in some code block. For this
hands-on exercise, create a func which “encloses” the scope of a variable:
*/



func main() {
	a := foo()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
}


func foo() func() int {
	var x int = 2
	return func() int {
		x = x * x
		return x
	}
}

