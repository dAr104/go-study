package main

import "fmt"

/*
Use the “defer” keyword to show that a deferred func runs after the func containing it exits.
*/

func main() {
	defer foo()
	bar()
	
}

func foo() {
	fmt.Println("This function will be defer than it will be run at the end")
}

func bar() {
	fmt.Println("This is the second function in order in the main")
}