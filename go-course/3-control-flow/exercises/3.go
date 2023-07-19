package main

import "fmt"

/*
Create a for loop using this syntax
● for condition { }
Have it print out the years you have been alive
*/

func main() {
	i := 1995
	for i < 2024 {
		fmt.Println(i)
		i++
	}
}