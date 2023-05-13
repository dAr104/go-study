package main

import "fmt"

/*
A deferred func runs after the func containing it exits.

A defer statement defers the execution of a function until the surrounding function returns.
The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
The first defer in order will be the last executed

Defer is used to ensure that a function call is performed later in a program’s execution,
usually for purposes of cleanup. defer is often used where e.g. ensure and finally would be used in other languages.
*/

func main(){
	defer foo()
	bar()
	who()
	fmt.Println("Return main")
}

func foo() {
	fmt.Println("Foo")
}

func bar() {
	fmt.Println("Bar")
}

func who() {
	fmt.Println("Who")
}