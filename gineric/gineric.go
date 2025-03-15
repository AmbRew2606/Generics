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
