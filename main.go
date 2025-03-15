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
	fmt.Println("Подсчет:")
	fmt.Println(len(sint))
	fmt.Println()
	fmt.Println()
	fmt.Println("Реверс:")
	genutils.Reverse(sint)
	fmt.Println(sint)
	fmt.Println()
	genutils.Reverse(sstr)
	fmt.Println(sstr)
	fmt.Println()
	fmt.Println()
	fmt.Println("Максимальные значения:")

	resultSintMax, err := genutils.MaxValue(sint)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultSintMax)
	}
	resultSstrMax, err := genutils.MaxValue(sstr)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultSstrMax)
	}
	fmt.Println()
	fmt.Println()
	fmt.Println("Преобразование:")
	resultSintMap := genutils.Map(sint, func(n int) string {
		return fmt.Sprintf("Number: %d", n)
	})
	fmt.Println(resultSintMap)
	resultSstrMap := genutils.Map(sstr, func(n string) int {
		return len(n)
	})
	fmt.Println(resultSstrMap)
}
