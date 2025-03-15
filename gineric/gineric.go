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

func MaxValue[T constraints.Ordered](s []T) (T, error) {
	if len(s) == 0 {
		var zero T
		return zero, fmt.Errorf("срез пуст")
	}
	max := s[0]
	for _, v := range s {
		if v > max {
			max = v
		}
	}
	return max, nil
}

func Map[T, U any](s []T, fn func(T) U) []U {
	var result []U
	for _, v := range s {
		result = append(result, fn(v))
	}
	return result
}
