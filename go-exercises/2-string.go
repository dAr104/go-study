
package main

import (
	"fmt"
	"strings"
)

func main() {

	s := strings.Join([]string{"Hello", "World!", "joins", "this", "with", "comma", "separator"}, ", ")
	fmt.Println(s)
}