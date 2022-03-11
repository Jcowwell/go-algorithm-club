package util

import (
	"errors"
	"math"

	"golang.org/x/exp/constraints"
)

type Pair[T constraints.Ordered] struct {
	x    T
	y    T
	init bool
}

func (self Pair[T]) Compare(other Pair[T]) int {
	if self.x < other.x {
		return -1
	}
	if self.x > other.x {
		return 1
	}
	if self.y < other.y {
		return -1
	}
	if self.y < other.y {
		return 1
	}
	return 0 // x and y must equal eachother
}

func PairInit[T constraints.Ordered](x, y T) *Pair[T] {
	return &Pair[T]{x: x, y: y, init: true}
}

type PointType int

const (
	OpenPoint      PointType = iota // exclusive Point
	ClosedPoint                     // inclusive Point
	UnboundedPoint                  // infinity
)

type Point[N Numeric] struct {
	value     N
	pointType PointType
}

// func (self Point[N]) Compare(point Point[N]) int {
// 	if self.pointType == point.pointType {
// 		if self.value < point.value {
// 			return -1
// 		} else if self.value > point.value {
// 			return 1
// 		} else {
// 			return 0
// 		}
// 	}

// }

type IntervalType int

const (
	EmptyInterval       IntervalType = iota
	OpenInterval                     // {x | a < x < b}
	ClosedInterval                   // {x | a <= x <= b}
	OpenClosedInterval               // {x | a < x <= b}
	ClosedOpenInterval               // {x | a <= x < b}
	GreaterThanInterval              // {x | x > a}
	AtLeastInterval                  // {x | x >= a}
	LessThanInterval                 // {x | x < b}
	AtMostInterval                   // {x | x <= b}
	UnboundedInterval                // {x}
)

type Interval[I constraints.Integer] struct {
	lowerBound   Point[I] // lowerBound
	upperBound   Point[I] // upperBound
	intervalType IntervalType
}

func IntervalInit[I constraints.Integer](lo, hi Point[I]) Interval[I] {
	intervalType := setIntervalType(lo.pointType, hi.pointType)
	return Interval[I]{lowerBound: lo, upperBound: hi, intervalType: intervalType}
}

func setIntervalType(loType, hiType PointType) IntervalType {
	if loType == OpenPoint && hiType == OpenPoint {
		return OpenInterval
	}
	if loType == ClosedPoint && hiType == ClosedPoint {
		return ClosedInterval
	}
	if loType == OpenPoint && hiType == ClosedPoint {
		return OpenClosedInterval
	}
	if loType == ClosedPoint && hiType == OpenPoint {
		return ClosedOpenInterval
	}
	if loType == OpenPoint && hiType == UnboundedPoint {
		return GreaterThanInterval
	}
	if loType == ClosedPoint && hiType == UnboundedPoint {
		return AtLeastInterval
	}
	if loType == UnboundedPoint && hiType == OpenPoint {
		return LessThanInterval
	}
	if loType == UnboundedPoint && hiType == ClosedPoint {
		return AtMostInterval
	}
	if loType == UnboundedPoint && hiType == UnboundedPoint {
		return UnboundedInterval
	}
	return EmptyInterval
}

/* Interval Generation Functions */

func GenerateEmptyInterval[I constraints.Integer]() Interval[I] {
	return Interval[I]{}
}

func GenerateOpenInterval[I constraints.Integer](x, y I) Interval[I] {
	xPoint := Point[I]{value: x, pointType: OpenPoint}
	yPoint := Point[I]{value: y, pointType: ClosedPoint}
	return IntervalInit(xPoint, yPoint)
}

func GenerateClosedInterval[I constraints.Integer](x, y I) Interval[I] {
	xPoint := Point[I]{value: x, pointType: ClosedPoint}
	yPoint := Point[I]{value: y, pointType: ClosedPoint}
	return IntervalInit(xPoint, yPoint)
}

func GenerateClosedOpenInterval[I constraints.Integer](x, y I) Interval[I] {
	xPoint := Point[I]{value: x, pointType: ClosedPoint}
	yPoint := Point[I]{value: y, pointType: OpenPoint}
	return IntervalInit(xPoint, yPoint)
}

func GenerateGreaterThanInterval[I constraints.Integer](x I) Interval[I] {
	xPoint := Point[I]{value: x, pointType: OpenPoint}
	yPoint := Point[I]{value: I(math.Inf(0)), pointType: UnboundedPoint}
	return IntervalInit(xPoint, yPoint)
}

func GenerateAtLeastInterval[I constraints.Integer](x I) Interval[I] {
	xPoint := Point[I]{value: x, pointType: ClosedPoint}
	yPoint := Point[I]{value: I(math.Inf(0)), pointType: UnboundedPoint}
	return IntervalInit(xPoint, yPoint)
}

func GenerateLessThanInterval[I constraints.Integer](y I) Interval[I] {
	xPoint := Point[I]{value: I(math.Inf(-1)), pointType: UnboundedPoint}
	yPoint := Point[I]{value: y, pointType: OpenPoint}
	return IntervalInit(xPoint, yPoint)
}

func GenerateAtMostInterval[I constraints.Integer](y I) Interval[I] {
	xPoint := Point[I]{value: I(math.Inf(-1)), pointType: UnboundedPoint}
	yPoint := Point[I]{value: y, pointType: ClosedPoint}
	return IntervalInit(xPoint, yPoint)
}

func GenerateUnboundedInterval[I constraints.Integer]() Interval[I] {
	xPoint := Point[I]{value: I(math.Inf(-1)), pointType: UnboundedPoint}
	yPoint := Point[I]{value: I(math.Inf(0)), pointType: UnboundedPoint}
	return IntervalInit(xPoint, yPoint)
}

func (self Interval[I]) Contains(value I) bool {
	switch self.intervalType {
	case EmptyInterval:
		return false
	case OpenInterval:
		return self.lowerBound.value < value && self.upperBound.value > value
	case ClosedInterval:
		return self.lowerBound.value <= value && self.upperBound.value >= value
	case OpenClosedInterval:
		return self.lowerBound.value < value && self.upperBound.value >= value
	case ClosedOpenInterval:
		return self.lowerBound.value <= value && self.upperBound.value > value
	case GreaterThanInterval:
		return self.lowerBound.value < value
	case AtLeastInterval:
		return self.lowerBound.value <= value
	case LessThanInterval:
		return self.upperBound.value > value
	case AtMostInterval:
		return self.upperBound.value >= value
	case UnboundedInterval:
		return true
	}
	return false
}

func (self Interval[I]) Count(value I) int {
	switch self.intervalType {
	case EmptyInterval:
		return 0
	case OpenInterval:
		return int((self.upperBound.value - 1) - (self.lowerBound.value + 1))
	case ClosedInterval:
		return int((self.upperBound.value) - (self.lowerBound.value))
	case OpenClosedInterval:
		return int((self.upperBound.value - 1) - (self.lowerBound.value))
	case ClosedOpenInterval:
		return int((self.upperBound.value) - (self.lowerBound.value + 1))
	case GreaterThanInterval, AtLeastInterval, LessThanInterval, AtMostInterval, UnboundedInterval:
		return int(math.Inf(0))
	}
	return 0
}

func (self Interval[I]) ToSlice() []I {
	interval := []I{}
	switch self.intervalType {
	case EmptyInterval:
		break
	case OpenInterval:
		for x := self.lowerBound.value + 1; x < self.upperBound.value; x++ {
			interval = append(interval, x)
		}
		break
	case ClosedInterval:
		for x := self.lowerBound.value; x <= self.upperBound.value; x++ {
			interval = append(interval, x)
		}
		break
	case OpenClosedInterval:
		for x := self.lowerBound.value + 1; x <= self.upperBound.value; x++ {
			interval = append(interval, x)
		}
		break
	case ClosedOpenInterval:
		for x := self.lowerBound.value + 1; x < self.upperBound.value; x++ {
			interval = append(interval, x)
		}
		break
	case GreaterThanInterval:
		for x := self.lowerBound.value + 1; x < math.MaxInt8; x++ {
			interval = append(interval, (x))
		}
		break
	case AtLeastInterval:
		for x := self.lowerBound.value; x < math.MaxInt8; x++ {
			interval = append(interval, x)
		}
		break
	case LessThanInterval:
		for x := self.upperBound.value - 1; x > math.MaxInt8; x-- {
			interval = append(interval, x)
		}
		break
	case AtMostInterval:
		for x := self.upperBound.value; x > math.MaxInt8; x-- {
			interval = append(interval, x)
		}
		break
	case UnboundedInterval:
		for x := math.MinInt8; x > math.MaxInt8; x-- {
			interval = append(interval, I(x))
		}
		break
	}
	return interval
}

func (self Interval[I]) Intersect(interval Interval[I]) {
	if interval.intervalType == EmptyInterval {
		return
	}
	a, b := self.lowerBound.value, self.upperBound.value
	c, d := interval.lowerBound.value, interval.upperBound.value

	if self.lowerBound.pointType == OpenPoint {
		a += 1
	}
	if self.upperBound.pointType == OpenPoint {
		b -= 1
	}
	if interval.lowerBound.pointType == OpenPoint {
		c += 1
	}
	if interval.upperBound.pointType == OpenPoint {
		d -= 1
	}
}

// func IntersectIntervals[I constraints.Integer](intervals ...Interval[I]) Interval[I] {
// 	itvl := GenerateEmptyInterval()
// 	for _, interval := range intervals {

// 	}
// }

type Range[T Numeric] []T

func (self Range[T]) IsEmpty() bool {
	return len(self) == 0
}

func (self Range[T]) Count() int {
	return len(self)
}
func (self Range[T]) UpperBound() (T, error) {
	if self.IsEmpty() {
		var value T
		return value, errors.New("Range is empty")
	}
	return self[len(self)], nil
}

func (self Range[T]) LowerBound() (T, error) {
	if self.IsEmpty() {
		var value T
		return value, errors.New("Range is empty")
	}
	return self[0], nil
}

func (self Range[T]) Generate(start, end T) {
	for i := start; i < end; i++ {
		self = append(self, i)
	}
}

func GenerateRange[T Numeric](start, end T) Range[T] {
	r := Range[T]{}
	r.Generate(start, end)
	return r
}
