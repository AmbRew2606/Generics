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

	fmt.Println("Печать:")
	genutils.Print(sint)
	fmt.Println()
	genutils.Print(sstr)
	fmt.Println()
	fmt.Println()
	fmt.Println("Реверс:")
	genutils.Reverse(sint)
	fmt.Println(sint)
	fmt.Println()
	genutils.Reverse(sstr)
	fmt.Println(sstr)

}
