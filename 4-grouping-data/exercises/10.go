package main

import "fmt"

/*
Using the code from the previous example, delete a record from your map. Now print the map
out using the “range” loop
*/

func main() {

	m := map[string][]string{
		`bond_james`: []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`: []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	if v, ok := m["bond_james"]; ok {
		fmt.Println("Deleting", v)
		delete(m, "bond_james")
	}

	for k, v := range m {
		fmt.Println("Record for the key:", k)
		for i, e := range v {
			fmt.Println("\t", i, e)
		}
	}
}