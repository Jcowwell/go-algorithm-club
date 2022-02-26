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

func Max[T constraints.Ordered](elements ...T) (minima T) {
	if len(elements) < 1 {
		panic("Not enough arguments based for evaluation")
	} else {
		minima = elements[0]
		for _, element := range elements {
			if minima < element {
				minima = element
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

func Equal[T constraints.Ordered](a, b []T) bool {
	if a == nil || b == nil {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func Hash[A comparable](a A) uintptr {
	var m interface{} = make(map[A]struct{})
	hf := (*mh)(*(*unsafe.Pointer)(unsafe.Pointer(&m))).hf
	return hf(unsafe.Pointer(&a), 0)
}

func Swap[T any](x, y *T) {
	*x, *y = *y, *x
}

func Filler[T any](f T) T {
	return f
}
