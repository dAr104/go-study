package main

import "fmt"

func main() {
	x := "Bond"
	switch x {
	case "James":
		fmt.Println("this should not print")
	case "ond", "bond", "Bond":
		fmt.Println("this should print")	
	default:
		fmt.Println("this is default")
	}
}