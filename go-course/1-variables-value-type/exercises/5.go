package main

import "fmt"

/*
1. at the package level scope, using the “var” keyword, create a VARIABLE with the
IDENTIFIER “y”. The variable should be of the UNDERLYING TYPE of your custom
TYPE “x”
a. eg:
2. in func main
a. this should already be done
i. print out the value of the variable “x”
ii. print out the type of the variable “x”
iii. assign your own VALUE to the VARIABLE “x” using the “=” OPERATOR
iv. print out the value of the variable “x”
b. now do this
i. now use CONVERSION to convert the TYPE of the VALUE stored in “x”
to the UNDERLYING TYPE
1. then use the “=” operator to ASSIGN that value to “y”
2. print out the value stored in “y”
3. print out the type of “y”
*/

type myType int
var xxx myType
var yyy int

func main() {
	fmt.Println(xxx)
	fmt.Printf("%T\n", xxx)
	xxx = 42
	fmt.Println(xxx)

	yyy = int(xxx)
	fmt.Println(yyy)
	fmt.Printf("%T\n", yyy)
}
