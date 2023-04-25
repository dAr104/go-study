package main

import "fmt"

/*
Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional
slice:
		"James", "Bond", "Shaken, not stirred"
		"Miss", "Moneypenny", "Helloooooo, James."
Range over the records, then range over the data in each record
*/

func main() {

	x := []string{"James", "Bond", "Shaken, not stirred"}
	y := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	matrice := [][]string{x,y}

	for _, v := range matrice {
		for _, s := range v {
			fmt.Println(s)
		}
	}
}