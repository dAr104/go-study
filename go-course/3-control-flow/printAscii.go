package main

import "fmt"

func main() {
	for i := 33; i <= 122; i++ {
		// print UNICODE ascii and character corresponding to number i
		fmt.Printf("%v\t%#U\n", i, i)
	}
}