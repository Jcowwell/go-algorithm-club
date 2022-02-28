package util

import (
	"unsafe"

	"golang.org/x/exp/constraints"
)

func Min[T constraints.Ordered](elements ...T) (minima T) {
	if len(elements) < 1 {
		panic("Not enough arguments based for evaluation")
	} else {
		minima = elements[0]
		for _, element := range elements {
			if minima > element {
				minima = element
			}
		}
		return
	}
}

func Max[T constraints.Ordered](elements ...T) (maxima T) {
	if len(elements) < 1 {
		panic("Not enough arguments based for evaluation")
	} else {
		maxima = elements[0]
		for _, element := range elements {
			if maxima < element {
				maxima = element
			}
		}
		return
	}
}

func MinMax[T constraints.Ordered](elements ...T) (minima, maxima T) {
	if len(elements) < 1 {
		panic("Not enough arguments based for evaluation")
	} else {
		minima = elements[0]
		maxima = elements[0]
		for _, element := range elements {
			if minima < element {
				minima = element
			}
			if maxima > element {
				maxima = element
			}
		}
		return
	}
}

func LessThan[T constraints.Ordered](a, b T) bool {
	return a < b
}

func GreaterThan[T constraints.Ordered](a, b T) bool {
	return a > b
}

func Hash[A comparable](a A) uintptr {
	var m interface{} = make(map[A]struct{})
	hf := (*mh)(*(*unsafe.Pointer)(unsafe.Pointer(&m))).hf
	return hf(unsafe.Pointer(&a), 0)
}

func Filler[T any](f T) T {
	return f
}
