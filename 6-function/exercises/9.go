package main

import "fmt"

/*
A “callback” is when we pass a func into a func as an argument. For this exercise,
● pass a func into a func as an argument
*/



func main() {
	callback(foo, 4)
}

func callback(f func(x int) int, y int) {
	fmt.Println("Callback:", f(y))
}

func foo(x int) int {
	return x
}

