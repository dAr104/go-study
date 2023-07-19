package main

import "fmt"

/*
Create a for loop using this syntax
● for { }
Have it print out the years you have been alive.
*/

func main() {
	i := 1995
	for {
		if i > 2023 {
			break
		}
		fmt.Println(i)
		i++
	}
}