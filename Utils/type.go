package util

import (
	"unsafe"

	"golang.org/x/exp/constraints"
)

type Numeric interface {
	constraints.Float | constraints.Integer
}

type Transform[T, U any] func(T) U

type Predicate[T any] func(T) bool

///////////////////////////
/// stolen from runtime ///
///////////////////////////
type mh struct {
	_  uintptr
	_  uintptr
	_  uint32
	_  uint8
	_  uint8
	_  uint8
	_  uint8
	_  func(unsafe.Pointer, unsafe.Pointer) bool
	_  *byte
	_  int32
	_  int32
	_  unsafe.Pointer
	_  unsafe.Pointer
	_  unsafe.Pointer
	hf func(unsafe.Pointer, uintptr) uintptr
}

// Tree Node Interface
type TreeNode[T constraints.Ordered] interface {
	count()
	height() int
	depth() int
	isRoot() bool
	isLeaf() bool
}
