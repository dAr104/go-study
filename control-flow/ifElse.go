package main

import "fmt"

func main() {
	x := 42;
	if x == 2 { // two statment in one line, use semicolon
		fmt.Println("x is 2")
	} else if x == 41{
		fmt.Printf("x is 41")
	} else {
		fmt.Println("x is not 41 or 2")
	}
}