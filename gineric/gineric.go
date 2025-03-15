package genutils

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func Print[T constraints.Ordered](s []T) {
	for _, v := range s {
		fmt.Print(v)
	}
}

func Reverse[T constraints.Ordered](s []T) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
