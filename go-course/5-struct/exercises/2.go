package main	

import "fmt"

/*
Take the code from the previous exercise, then store the values of type person in a map with the
key of last name. Access each value in the map. Print out the values, ranging over the slice.
*/

type person struct {
	first string
	last string
	iceCreamFlavors []string
}

func main() {

	p1 := person{
		first: "Dario",
		last: "Rossi",
		iceCreamFlavors: []string{"Cocco", "Tiramisù"},
	}

	p2 := person{
		first: "Andrea",
		last: "Espposito",
		iceCreamFlavors: []string{"Stracciatella", "Limone", "Fragola"},
	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range m {
		fmt.Println(v.last, v.first)
		for i, e := range v.iceCreamFlavors {
			fmt.Println(i, e)
		}
		fmt.Println("-------")
	}
}