package main

import "fmt"

// Go is a STATIC programming language
// A variable is declared to hold a value of a specific type

var y = 42
var z string = "Shaken, not stirred"

// create my type
type hotdog int
var b hotdog

func main() {
	fmt.Println(y)
	// %T prints the type of the variable, printf is a print function that allows to print a formatted string
	fmt.Printf("%T\n", y) 
	fmt.Printf("%b\n", y) // %b prints the value in binary 
	fmt.Printf("%x\n", y) // %x prints the value in hexadecimal
	fmt.Printf("%#x\n", y) // %#x prints the value in hexadecimal with 0x in front
	fmt.Printf(z)
	fmt.Printf("%T\n", z)

	s := fmt.Sprintf("%#x\t%b\t%x", y, y, y) // Sprintf returns a string
	fmt.Println(s)

	b = 43
	fmt.Printf("%T\n", b)
	//a = b // cannot use b (type hotdog) as type int in assignment
	a = int(b) // conversion, called "casting" in other languages

}
