package main

import "fmt"

func main() {
	// print only even number
	x := 1
	for {

		x++
		if x > 100 {
			break
		}

		if x % 2 != 0 {
			continue
		}

		fmt.Println(x)
	}
}