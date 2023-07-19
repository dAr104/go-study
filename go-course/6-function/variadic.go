package main

import "fmt"

func main() {
	x := sum(1,2,3,4)
	fmt.Println("The final total is", x)

	xi := []int{2, 3, 4, 5, 6, 7, 8, 9}
	y := sum(xi...) // x... unfurling the slice xi in n ints: no new slices will be created, it will be point also tu xi
	fmt.Println("The final total of xi is", y)

	// a variadic parameter also accept zero values
	sum()
}


// funcion with variadic parameter: unlimited value (you have to specify the type)
// NOTE: the variadic parameter must to be the last parameter of the function
func sum(x ...int) int {
	// what is x? --> slice of int: you pass a set of single arguments and it store all in a slice
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		fmt.Println("for item in index position", i, "we are now adding", v, "to the total which is", sum)
		sum += v
	}
	return sum	
}