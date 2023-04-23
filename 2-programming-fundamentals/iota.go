package main

import "fmt"

// iota identifier to auto incrementig constant
const (
	a = iota
	b
	c 
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}