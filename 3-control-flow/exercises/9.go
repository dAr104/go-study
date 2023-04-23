package main

import "fmt"

/*
Create a program that uses a switch statement with the switch expression specified as a
variable of TYPE string with the IDENTIFIER “favSport”.
*/

func main() {
	favSport := "Swimming"
	switch {
	case favSport == "Soccer":
		fmt.Println("I love play soccer")
	case favSport == "Swimming":
		fmt.Println("I love swimming")	
	case favSport == "Tennis":
		fmt.Println("I love play tennis")
	}
}