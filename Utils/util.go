package util

import (
	"unsafe"

	"golang.org/x/exp/constraints"
)

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
