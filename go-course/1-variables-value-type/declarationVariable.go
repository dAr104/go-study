package main

import "fmt"

// DECLARE USING var keyword

// declare the variable "y" and assign the value 23 --> declare & assign = initilization
var a = 23 // can be used in and outside the function scope

// Declare there is a variable with idetifier "z" and that the variable with identifier "z" is of type int, and assign the zero value
var z int // = var z int = 0

func main() {
	// SHORT DECLARATION OPERATOR
	x := 42 // Declare a variable and assigns value to it. It can be used only within the function scope
	fmt.Println(x)
	x = 99 // now i assign new value. I have to use := only when i declare a variable
	fmt.Println(x)
	/* a statement is the smallest standalone element of a program that expresses some action to be carried out.
	it is ana instruction that commands the computer to perform a specified action. A program is formed by a sequence of one or more statements */
	y := 100 + 24
	fmt.Println(y)
	fmt.Println(a, z)
}

// BEST PRACTICE: try to limite the scope of the variables and use short declaration operator