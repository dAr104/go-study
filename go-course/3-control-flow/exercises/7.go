package main

import "fmt"

/*
Building on the previous hands-on exercise, create a program that uses “else if” and “else”
*/

func main() {
	learning := "JS"
	if learning == "Go" {
		fmt.Println("I'm learning Go")
	} else if learning == "Rust" {
		fmt.Println("I'm learning Rust")
	} else {
		fmt.Println("I'm not learning anything")
	}
}