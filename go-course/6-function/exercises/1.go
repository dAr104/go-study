package main

import "fmt"

/*
Hands on exercise
○ create a func with the identifier foo that returns an int
○ create a func with the identifier bar that returns an int and a string
○ call both funcs
○ print out their results
*/

func main() {
	x := foo(3)
	fmt.Println(x)
	y, s := bar(5, "Ciao")
	fmt.Println(y, s)
}

func foo(x int) int {
	return x
}

func bar(x int, s string) (int, string) {
	return x, s;
}