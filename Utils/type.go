package util

import (
	"constraints"
)

type Numeric interface {
	constraints.Float | constraints.Integer
}
