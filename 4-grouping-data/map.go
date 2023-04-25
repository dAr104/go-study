package main		

import "fmt"

// key value store, unordered 
func main() {
	// maps hold references to an underlying data structure. If you pass a map to a function that changes the contents
	// of the map, the changes will be visibile in the caller
	m := map[string]int{ // key is a string, value is int
		"James": 32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["Dario"]) // no key "Dario" in the map, return 0 value (zero value of int)

	// if return 0 can be that value for that key is really zero. How check if the key exits?
	// i will use "comma ok" idiom: if "Dario" is present v will be set appropriately and ok will be true;
	// if not, v will be set to zero and ok will be false
	v, ok := m["Dario"]
	fmt.Println(v)
	fmt.Println(ok)

	// comma ok idioms with if condizion
	if v, ok := m["Dario"]; ok {
		fmt.Println(v)
	}

	// add new element in the map
	m["Todd"] = 33

	// range with map
	for k, v := range m {
		fmt.Printf("%v\t%v\n", k, v)
	}

	// delete from a map
	delete(m, "James")
	fmt.Println(m)

	// is better check if exits the key before deleting it
	if v, ok := m["Wee"]; ok {
		fmt.Println("Value exits: ", v)
		delete(m, "Wee")
	}
}