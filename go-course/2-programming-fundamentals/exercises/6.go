package main

import "fmt"

/*
Using iota, create 4 constants for the NEXT 4 years. Print the constant values
*/

const (
	_ = iota
	y = 2023 + iota
	yy = 2023 + iota
	yyy = 2023 + iota
	yyyy = 2023 + iota
)

func main() {

	fmt.Println(y, yy, yyy, yyyy)
}