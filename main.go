package main

import (
	genutils "Generics/gineric"
	"fmt"
)

func main() {
	// s := make([]int,10)
	var sint []int
	var sstr []string

	sint = append(sint, 1, 2, 3)
	sstr = append(sstr, "first", "second")

	genutils.Print(sint)
	fmt.Println()
	genutils.Print(sstr)
}
