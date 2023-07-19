package main

import "fmt"

/*Write a program that prints a number in decimal, binary, and hex*/

func main() {
	x := 33
	fmt.Printf("%d\n", x)
	fmt.Printf("%b\n", x)
	fmt.Printf("%x\n", x)
	fmt.Printf("%#x\n", x)
}