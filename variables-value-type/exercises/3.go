package main

import "fmt"

/*
Using the code from the previous exercise,
1. At the package level scope, assign the following values to the three variables
a. for x assign 42
b. for y assign “James Bond”
c. for z assign true
2. in func main
Todd McLeod - Learn To Code Go on Udemy - Part 1 - Page 26
a. use fmt.Sprintf to print all of the VALUES to one single string. ASSIGN the
returned value of TYPE string using the short declaration operator to a
VARIABLE with the IDENTIFIER “s”
b. print out the value stored by variable “s”
*/

var xx int = 42
var yy string = "James Bond"
var zz = true

func main() {
	s := fmt.Sprintf("%v\t%v\t%v", xx, yy, zz)
	fmt.Println(s)
}