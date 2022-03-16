package util

import (
	"errors"
	"math"

	"golang.org/x/exp/constraints"
)

type PointType int // Point Enum Type

const (
	OpenPoint      PointType = iota // exclusive point
	ClosedPoint                     // inclusive point
	UnboundedPoint                  // infinity
)

/* Point Type to represent a number on a number line. */
type Point[N Numeric] struct {
	value     N
	pointType PointType
}

/* Private function that calculates which start point should be used for interval intersection */
func startPointIntersect[N Numeric](a, b Point[N]) Point[N] {
	if a.value == b.value {
		if a.pointType < b.pointType {
			return a
		} else {
			return b
		}
	} else if a.value > b.value {
		return a
	} else {
		return b
	}
}

/* Private function that calculates which end point should be used for interval intersection */
func endPointIntersect[N Numeric](c, d Point[N]) Point[N] {
	if c.value == d.value {
		if c.pointType < d.pointType {
			return c
		} else {
			return d
		}
	} else if c.value < d.value {
		return c
	} else {
		return d
	}
}

type IntervalType int // Interval Enum Type

const (
	EmptyInterval       IntervalType = iota
	DegenerateInterval               // [a,a] = {a}
	OpenInterval                     // (a,b) = {x | a < x < b}
	ClosedInterval                   // [a,b] = {x | a <= x <= b}
	OpenClosedInterval               // (a,b] = {x | a < x <= b}
	ClosedOpenInterval               // [a,b) = {x | a <= x < b}
	GreaterThanInterval              // (a,+∞) = {x | x > a}
	AtLeastInterval                  // [a,+∞) = {x | x >= a}
	LessThanInterval                 // (-∞,b) = {x | x < b}
	AtMostInterval                   // (-∞,b] = {x | x <= b}
	UnboundedInterval                // (-∞,+∞) = {x}
)

type Interval[I constraints.Integer] struct {
	lowerBound   Point[I] // start point of interval
	upperBound   Point[I] // end point of interval
	values       []I      // values of an interval. For unbounded intervals values will be nil
	intervalType IntervalType
}

/*
	private construction function to create an interval.

	Parameters:
		lo Point[I] Lowerbound Value
		hi Point[I]	Higherbound Value
	Return:
		Interval[I] Interval Struct
*/
func createInterval[I constraints.Integer](start, end Point[I]) Interval[I] {
	if start.value > end.value {
		panic("The lowerbound endpoint cannot be higher than the upperbound endpoint")
	}
	interval := Interval[I]{lowerBound: start, upperBound: end}
	interval.setIntervalType()
	interval.setValues()
	return interval
}

/* Private void method to set an interval's type. */
func (self *Interval[I]) setIntervalType() {
	var loType, hiType PointType = self.lowerBound.pointType, self.upperBound.pointType

	/* OpenInterval */
	if loType == OpenPoint && hiType == OpenPoint {
		/* Empty Check: (a,a) = {} */
		if self.lowerBound.value == self.upperBound.value {
			self.intervalType = DegenerateInterval
			return
		}
		self.intervalType = OpenInterval

		/* ClosedInterval */
	} else if loType == ClosedPoint && hiType == ClosedPoint {
		/* Degenerate Check */
		if self.lowerBound.value == self.upperBound.value {
			self.intervalType = DegenerateInterval
			return
		}
		self.intervalType = ClosedInterval

		/* OpenClosedInterval Interval */
	} else if loType == OpenPoint && hiType == ClosedPoint {
		/* Empty Check: (a,a] = {} */
		if self.lowerBound.value == self.upperBound.value {
			self.intervalType = DegenerateInterval
			return
		}
		self.intervalType = OpenClosedInterval

		/* ClosedOpenInterval Interval */
	} else if loType == ClosedPoint && hiType == OpenPoint {
		/* Empty Check: [a,a) = {} */
		if self.lowerBound.value == self.upperBound.value {
			self.intervalType = DegenerateInterval
			return
		}
		self.intervalType = ClosedOpenInterval

		/* GreaterThanInterval Interval */
	} else if loType == OpenPoint && hiType == UnboundedPoint {
		self.intervalType = GreaterThanInterval

		/* AtLeastInterval Interval */
	} else if loType == ClosedPoint && hiType == UnboundedPoint {
		self.intervalType = AtLeastInterval

		/* LessThanInterval Interval */
	} else if loType == UnboundedPoint && hiType == OpenPoint {
		self.intervalType = LessThanInterval

		/* AtMostInterval Interval */
	} else if loType == UnboundedPoint && hiType == ClosedPoint {
		self.intervalType = AtMostInterval

		/* UnboundedInterval Interval */
	} else if loType == UnboundedPoint && hiType == UnboundedPoint {
		self.intervalType = UnboundedInterval
	}
}

/*
	Private void method to set an interval's set of values.

	NOTE:
	The initial implementaiton of this method populated endpoints with the interval's start and stop *inclusive* endpoints
	for a user to use (i.e iteration). This was problematic when met with edge cases such as (1,2) where,
	according to the previous logic, the endpoint slice would be []int{2,1}.
	This would defeat the original purpose of the method.
	Instead it will be a method to directy create a slice that holds the inclusive values of valid bounded intervals.
*/
func (self *Interval[I]) setValues() {
	switch self.intervalType {
	case DegenerateInterval:
		self.values = append(self.values, self.lowerBound.value)
	case OpenInterval:
		for n := self.lowerBound.value + 1; n < self.upperBound.value; n++ {
			self.values = append(self.values, n)
		}
		break
	case ClosedInterval:
		for n := self.lowerBound.value; n <= self.upperBound.value; n++ {
			self.values = append(self.values, n)
		}
		break
	case OpenClosedInterval:
		for n := self.lowerBound.value + 1; n <= self.upperBound.value; n++ {
			self.values = append(self.values, n)
		}
		break
	case ClosedOpenInterval:
		for n := self.lowerBound.value; n < self.upperBound.value; n++ {
			self.values = append(self.values, n)
		}
		break
	/* Unbounded (and Unbounded Hybrid) types cannot set a value type since they are not finite */
	case EmptyInterval, GreaterThanInterval, AtLeastInterval, LessThanInterval, AtMostInterval, UnboundedInterval:
		break
	}
}

/* SECTION: Interval Generation Functions */

/*
	Public Construction Function to generate an interval.

	Parameters:
		LowerBound Point[I] 	Start endpoint of interval.
		HigherBound Point[I]	End endpoint of interval.
	Return:
		Interval[I] Interval Struct
*/
func GenerateInterval[I constraints.Integer](lowerBound, upperBound Point[I]) Interval[I] {
	return createInterval(lowerBound, upperBound)
}

/*

	Public Function to generate an Empty Interval.

	Return:
		Interval[I] Interval Struct

	FIXME: Add support for Numerics {N} over integers {I}
	Emptend Intervals can be created when there is an intersection between
	two intervals that do not intersect (b_start > a_end || a_start > b_end),
	if the intervals intersection is open and closed around the same point ex : {(1,3] ∩ (-∞, 2) => (1,2)},
	or if the interval intersection is  at the same point but one point is Open
	and the other Closed ex: {(2,3] ∩ (-1, 2]}
*/
func GenerateEmptyInterval[I constraints.Integer]() Interval[I] {
	return Interval[I]{}
}

/*
	Public Function to generate an Open Interval.

	Parameters:
		start I 	Start endpoint of interval.
		end I		End endpoint of interval.
	Return:
		Interval[I] Interval Struct

	FIXME: Add support for Numerics {N} over integers {I}
	The current implementation of Intervals only support Integers.
	Once Golang supports a more mature Generics this method should be updated to support Numerics.
*/
func GenerateOpenInterval[I constraints.Integer](start, end I) Interval[I] {
	lowerBound := Point[I]{value: start, pointType: OpenPoint}
	higherBound := Point[I]{value: end, pointType: OpenPoint}
	return createInterval(lowerBound, higherBound)
}

/*
	Public Function to generate a Closed Interval.

	Parameters:
		start I 	Start endpoint of interval.
		end I		End endpoint of interval.
	Return:
		Interval[I] Interval Struct

	FIXME: Add support for Numerics {N} over integers {I}
	The current implementation of Intervals only support Integers.
	Once Golang supports a more mature Generics this method should be updated to support Numerics.
*/
func GenerateClosedInterval[I constraints.Integer](start, end I) Interval[I] {
	lowerBound := Point[I]{value: start, pointType: ClosedPoint}
	higherBound := Point[I]{value: end, pointType: ClosedPoint}
	return createInterval(lowerBound, higherBound)
}

/*
	Public Function to generate a OpenClosed Interval.

	Parameters:
		start I 	Start endpoint of interval.
		end I		End endpoint of interval.
	Return:
		Interval[I] Interval Struct

	FIXME: Add support for Numerics {N} over integers {I}
	The current implementation of Intervals only support Integers.
	Once Golang supports a more mature Generics this method should be updated to support Numerics.
*/
func GenerateOpenClosedInterval[I constraints.Integer](start, end I) Interval[I] {
	lowerBound := Point[I]{value: start, pointType: OpenPoint}
	higherBound := Point[I]{value: end, pointType: ClosedPoint}
	return createInterval(lowerBound, higherBound)
}

/*
	Public Function to generate a ClosedOpen Interval.

	Parameters:
		start I 	Start endpoint of interval.
		end I		End endpoint of interval.
	Return:
		Interval[I] Interval Struct

	FIXME: Add support for Numerics {N} over integers {I}
	The current implementation of Intervals only support Integers.
	Once Golang supports a more mature Generics this method should be updated to support Numerics.
*/
func GenerateClosedOpenInterval[I constraints.Integer](start, end I) Interval[I] {
	lowerBound := Point[I]{value: start, pointType: ClosedPoint}
	higherBound := Point[I]{value: end, pointType: OpenPoint}
	return createInterval(lowerBound, higherBound)
}

/*
	Public Function to generate a GreaterThan Interval.

	Parameters:
		start I 	Start endpoint of interval.
	Return:
		Interval[I] Interval Struct

	FIXME: Add support for Numerics {N} over integers {I}
	The current implementation of Intervals only support Integers.
	Once Golang supports a more mature Generics this method should be updated to support Numerics.
*/
func GenerateGreaterThanInterval[I constraints.Integer](start I) Interval[I] {
	lowerBound := Point[I]{value: start, pointType: OpenPoint}
	higherBound := Point[I]{value: I(math.Inf(0)), pointType: UnboundedPoint}
	return createInterval(lowerBound, higherBound)
}

/*
	Public Function to generate a AtLeast Interval.

	Parameters:
		start I 	Start endpoint of interval.
	Return:
		Interval[I] Interval Struct

	FIXME: Add support for Numerics {N} over integers {I}
	The current implementation of Intervals only support Integers.
	Once Golang supports a more mature Generics this method should be updated to support Numerics.
*/
func GenerateAtLeastInterval[I constraints.Integer](start I) Interval[I] {
	lowerBound := Point[I]{value: start, pointType: ClosedPoint}
	higherBound := Point[I]{value: I(math.Inf(0)), pointType: UnboundedPoint}
	return createInterval(lowerBound, higherBound)
}

/*
	Public Function to generate a LessThan Interval.

	Parameters:
		start I 	Start endpoint of interval.
	Return:
		Interval[I] Interval Struct

	FIXME: Add support for Numerics {N} over integers {I}
	The current implementation of Intervals only support Integers.
	Once Golang supports a more mature Generics this method should be updated to support Numerics.
*/
func GenerateLessThanInterval[I constraints.Integer](end I) Interval[I] {
	lowerBound := Point[I]{value: I(math.Inf(-1)), pointType: UnboundedPoint}
	higherBound := Point[I]{value: end, pointType: OpenPoint}
	return createInterval(lowerBound, higherBound)
}

/*
	Public Function to generate a AtMost Interval.

	Parameters:
		start I 	Start endpoint of interval.
	Return:
		Interval[I] Interval Struct

	FIXME: Add support for Numerics {N} over integers {I}
	The current implementation of Intervals only support Integers.
	Once Golang supports a more mature Generics this method should be updated to support Numerics.
*/
func GenerateAtMostInterval[I constraints.Integer](end I) Interval[I] {
	lowerBound := Point[I]{value: I(math.Inf(-1)), pointType: UnboundedPoint}
	higherBound := Point[I]{value: end, pointType: ClosedPoint}
	return createInterval(lowerBound, higherBound)
}

/*
	Public Function to generate a Unbounded Interval.

	Parameters:
		start I 	Start endpoint of interval.
	Return:
		Interval[I] Interval Struct

	FIXME: Add support for Numerics {N} over integers {I}
	The current implementation of Intervals only support Integers.
	Once Golang supports a more mature Generics this method should be updated to support Numerics.
*/
func GenerateUnboundedInterval[I constraints.Integer]() Interval[I] {
	lowerBound := Point[I]{value: I(math.Inf(-1)), pointType: UnboundedPoint}
	higherBound := Point[I]{value: I(math.Inf(0)), pointType: UnboundedPoint}
	return createInterval(lowerBound, higherBound)
}

/* !SECTION: Interval Generation Functions */

/*
	Public Boolean Method that returns true if a value is within an interval. False otherwise.

	Parameters:
		value I 	value of type Integer
	Return:
		bool

	FIXME: Add support for Numerics {N} over integers {I}
	The current implementation of Intervals only support Integers (I).
	Once Golang supports a more mature Generics this method should be updated to support Numerics.
*/
func (self *Interval[I]) Contains(value I) bool {
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

/*
	Public Integer Method that returns the amount of numbers in an interval. Unbounded intervals
	returns math.inf

	Return:
		int

	FIXME: Add support for Numerics {N} over integers {I}
	The current implementation of Intervals only support Integers (I).
	Once Golang supports a more mature Generics this method should be updated to support Numerics.
*/
func (self *Interval[I]) Count() int {
	switch self.intervalType {
	case EmptyInterval:
		return 0
	case DegenerateInterval:
		return 1
	case OpenInterval, ClosedInterval, OpenClosedInterval, ClosedOpenInterval:
		return len(self.values)
	case GreaterThanInterval, AtLeastInterval, LessThanInterval, AtMostInterval, UnboundedInterval:
		return int(math.Inf(0))
	}
	return 0
}

/*
	Public Interval Function that returns the intersect (∩) between two intervals.

	Parameters:
		a Interval[I]
		b Interval[I]
	Return:
		c Interval[I]	a ∩ b

	FIXME: Add support for Numerics {N} over integers {I}
	The current implementation of Intervals only support Integers (I).
	Once Golang supports a more mature Generics this method should be updated to support Numerics.
*/
func Intersect[I constraints.Integer](a, b Interval[I]) Interval[I] {
	/* an intersection involving an Empty Interval always results in an Empty Interval */
	if a.intervalType == EmptyInterval || b.intervalType == EmptyInterval {
		return GenerateEmptyInterval[I]()
	}
	/* adopted from: https://stackoverflow.com/a/325964/6427171. Checks if the two intervals intersect. */
	if Max(a.lowerBound.value, b.lowerBound.value) > Min(a.upperBound.value, b.upperBound.value) {
		return GenerateEmptyInterval[I]()
	} else { /* the intervals intersect */
		start, end := startPointIntersect(a.lowerBound, b.lowerBound), endPointIntersect(a.upperBound, b.upperBound)
		return GenerateInterval(start, end)
	}
}

/* DEPRECIATE: Use Interval instead. */

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
