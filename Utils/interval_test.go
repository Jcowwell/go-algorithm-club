package util

import (
	"math"
	"testing"
)

/* SECTION: Interval Generation Testing  */

func TestGenerateEmptyInterval(t *testing.T) {
	interval := GenerateEmptyInterval[int]()
	AssertEqual(len(interval.values), 0, t)
	AssertEqual(interval.intervalType, EmptyInterval, t)
}

func TestGenerateOpenInterval(t *testing.T) {
	interval := GenerateOpenInterval(0, 9)
	AssertEqual(len(interval.values), 8, t)
	AssertEqual(interval.lowerBound.value, 0, t)
	AssertEqual(interval.lowerBound.pointType, OpenPoint, t)
	AssertEqual(interval.values[0], 1, t)
	AssertEqual(interval.upperBound.value, 9, t)
	AssertEqual(interval.upperBound.pointType, OpenPoint, t)
	AssertEqual(interval.values[len(interval.values)-1], 8, t)
	AssertEqual(interval.Count(), 8, t)
	AssertFalse(interval.Contains(0), t)
	AssertFalse(interval.Contains(9), t)
	AssertTrue(interval.Contains(5), t)
	AssertEqual(interval.intervalType, OpenInterval, t)
}

func TestGenerateClosedInterval(t *testing.T) {
	interval := GenerateClosedInterval(0, 9)
	AssertEqual(len(interval.values), 10, t)
	AssertEqual(interval.lowerBound.value, 0, t)
	AssertEqual(interval.lowerBound.pointType, ClosedPoint, t)
	AssertEqual(interval.values[0], 0, t)
	AssertEqual(interval.upperBound.value, 9, t)
	AssertEqual(interval.upperBound.pointType, ClosedPoint, t)
	AssertEqual(interval.values[len(interval.values)-1], 9, t)
	AssertEqual(interval.Count(), 10, t)
	AssertTrue(interval.Contains(0), t)
	AssertTrue(interval.Contains(9), t)
	AssertTrue(interval.Contains(5), t)
	AssertEqual(interval.intervalType, ClosedInterval, t)
}

func TestGenerateOpenClosedInterval(t *testing.T) {
	interval := GenerateOpenClosedInterval(0, 9)
	AssertEqual(len(interval.values), 9, t)
	AssertEqual(interval.lowerBound.value, 0, t)
	AssertEqual(interval.lowerBound.pointType, OpenPoint, t)
	AssertEqual(interval.values[0], 1, t)
	AssertEqual(interval.upperBound.value, 9, t)
	AssertEqual(interval.upperBound.pointType, ClosedPoint, t)
	AssertEqual(interval.values[len(interval.values)-1], 9, t)
	AssertEqual(interval.Count(), 9, t)
	AssertFalse(interval.Contains(0), t)
	AssertTrue(interval.Contains(9), t)
	AssertTrue(interval.Contains(5), t)
	AssertEqual(interval.intervalType, OpenClosedInterval, t)
}

func TestGenerateClosedOpenInterval(t *testing.T) {
	interval := GenerateClosedOpenInterval(0, 9)
	AssertEqual(len(interval.values), 9, t)
	AssertEqual(interval.lowerBound.value, 0, t)
	AssertEqual(interval.lowerBound.pointType, ClosedPoint, t)
	AssertEqual(interval.values[0], 0, t)
	AssertEqual(interval.upperBound.value, 9, t)
	AssertEqual(interval.upperBound.pointType, OpenPoint, t)
	AssertEqual(interval.values[len(interval.values)-1], 8, t)
	AssertEqual(interval.Count(), 9, t)
	AssertTrue(interval.Contains(0), t)
	AssertFalse(interval.Contains(9), t)
	AssertTrue(interval.Contains(5), t)
	AssertEqual(interval.intervalType, ClosedOpenInterval, t)
}

func TestGenerateGreaterThanInterval(t *testing.T) {
	interval := GenerateGreaterThanInterval(9)
	AssertEqual(len(interval.values), 0, t)
	AssertEqual(interval.lowerBound.value, 9, t)
	AssertEqual(interval.lowerBound.pointType, OpenPoint, t)
	AssertEqual(interval.upperBound.value, int(math.Inf(0)), t)
	AssertEqual(interval.upperBound.pointType, UnboundedPoint, t)
	AssertFalse(interval.Contains(9), t)
	AssertTrue(interval.Contains(10), t)
	AssertTrue(interval.Contains(500), t)
	AssertEqual(interval.intervalType, GreaterThanInterval, t)
}

func TestGenerateAtLeastInterval(t *testing.T) {
	interval := GenerateAtLeastInterval(9)
	AssertEqual(len(interval.values), 0, t)
	AssertEqual(interval.lowerBound.value, 9, t)
	AssertEqual(interval.lowerBound.pointType, ClosedPoint, t)
	AssertEqual(interval.upperBound.value, int(math.Inf(0)), t)
	AssertEqual(interval.upperBound.pointType, UnboundedPoint, t)
	AssertTrue(interval.Contains(9), t)
	AssertTrue(interval.Contains(10), t)
	AssertTrue(interval.Contains(500), t)
	AssertEqual(interval.intervalType, AtLeastInterval, t)
}

func TestGenerateLessThanInterval(t *testing.T) {
	interval := GenerateLessThanInterval(9)
	AssertEqual(len(interval.values), 0, t)
	AssertEqual(interval.lowerBound.value, int(math.Inf(-1)), t)
	AssertEqual(interval.lowerBound.pointType, UnboundedPoint, t)
	AssertEqual(interval.upperBound.value, 9, t)
	AssertEqual(interval.upperBound.pointType, OpenPoint, t)
	AssertFalse(interval.Contains(9), t)
	AssertTrue(interval.Contains(5), t)
	AssertTrue(interval.Contains(-500), t)
	AssertEqual(interval.intervalType, LessThanInterval, t)
}

func TestGenerateAtMostInterval(t *testing.T) {
	interval := GenerateAtMostInterval(9)
	AssertEqual(len(interval.values), 0, t)
	AssertEqual(interval.lowerBound.value, int(math.Inf(-1)), t)
	AssertEqual(interval.lowerBound.pointType, UnboundedPoint, t)
	AssertEqual(interval.upperBound.value, 9, t)
	AssertEqual(interval.upperBound.pointType, ClosedPoint, t)
	AssertTrue(interval.Contains(9), t)
	AssertTrue(interval.Contains(5), t)
	AssertTrue(interval.Contains(-500), t)
	AssertEqual(interval.intervalType, AtMostInterval, t)
}

func TestGenerateUnboundInterval(t *testing.T) {
	interval := GenerateUnboundedInterval[int]()
	AssertEqual(len(interval.values), 0, t)
	AssertEqual(interval.lowerBound.value, int(math.Inf(-1)), t)
	AssertEqual(interval.lowerBound.pointType, UnboundedPoint, t)
	AssertEqual(interval.upperBound.value, int(math.Inf(0)), t)
	AssertEqual(interval.upperBound.pointType, UnboundedPoint, t)
	AssertTrue(interval.Contains(9), t)
	AssertTrue(interval.Contains(5), t)
	AssertTrue(interval.Contains(-500), t)
	AssertEqual(interval.intervalType, UnboundedInterval, t)
}

/* !SECTION: Interval Generation Testing  */

/* SECTION: Intersection Testing */

/* SECTION: OpenInterval Intersection */

/* OpenInterval ∩ OpenInterval */
func TestOOIntervalIntersect(t *testing.T) {
	a := GenerateOpenInterval(1, 4)
	b := GenerateOpenInterval(2, 5)
	c := Intersect(a, b)
	AssertEqual(len(c.values), 1, t)
	AssertEqual(c.lowerBound.value, 2, t)
	AssertEqual(c.lowerBound.pointType, OpenPoint, t)
	AssertEqual(c.values[0], 3, t)
	AssertEqual(c.upperBound.value, 4, t)
	AssertEqual(c.upperBound.pointType, OpenPoint, t)
	AssertTrue(c.Contains(3), t)
	AssertFalse(c.Contains(1), t)
	AssertFalse(c.Contains(5), t)
	AssertEqual(c.intervalType, OpenInterval, t)
}

/* OpenInterval ∩ ClosedInterval */
func TestOCIntervalIntersect(t *testing.T) {
	a := GenerateOpenInterval(1, 4)
	b := GenerateClosedInterval(2, 5)
	c := Intersect(a, b)
	AssertEqual(len(c.values), 2, t)
	AssertEqual(c.lowerBound.value, 2, t)
	AssertEqual(c.lowerBound.pointType, ClosedPoint, t)
	AssertEqual(c.values[0], 2, t)
	AssertEqual(c.upperBound.value, 4, t)
	AssertEqual(c.upperBound.pointType, OpenPoint, t)
	AssertEqual(c.values[1], 3, t)
	AssertTrue(c.Contains(2), t)
	AssertTrue(c.Contains(3), t)
	AssertFalse(c.Contains(1), t)
	AssertFalse(c.Contains(5), t)
	AssertEqual(c.intervalType, ClosedOpenInterval, t)
}

/* OpenInterval ∩ OpenClosedInterval */
func TestOOCIntervalIntersect(t *testing.T) {
	a := GenerateOpenInterval(1, 4)
	b := GenerateOpenClosedInterval(2, 5)
	c := Intersect(a, b)
	AssertEqual(len(c.values), 1, t)
	AssertEqual(c.lowerBound.value, 2, t)
	AssertEqual(c.lowerBound.pointType, OpenPoint, t)
	AssertEqual(c.values[0], 3, t)
	AssertEqual(c.upperBound.value, 4, t)
	AssertEqual(c.upperBound.pointType, OpenPoint, t)
	AssertTrue(c.Contains(3), t)
	AssertFalse(c.Contains(1), t)
	AssertFalse(c.Contains(5), t)
	AssertEqual(c.intervalType, OpenInterval, t)
}

/* OpenInterval ∩ ClosedOpenInterval */
func TestOCOIntervalIntersect(t *testing.T) {
	a := GenerateOpenInterval(1, 4)
	b := GenerateClosedOpenInterval(2, 5)
	c := Intersect(a, b)
	AssertEqual(len(c.values), 2, t)
	AssertEqual(c.lowerBound.value, 2, t)
	AssertEqual(c.lowerBound.pointType, ClosedPoint, t)
	AssertEqual(c.values[0], 2, t)
	AssertEqual(c.upperBound.value, 4, t)
	AssertEqual(c.upperBound.pointType, OpenPoint, t)
	AssertEqual(c.values[1], 3, t)
	AssertTrue(c.Contains(3), t)
	AssertFalse(c.Contains(1), t)
	AssertFalse(c.Contains(5), t)
	AssertEqual(c.intervalType, ClosedOpenInterval, t)
}

/* OpenInterval ∩ GreaterThanInterval */
func TestOGTIntervalIntersect(t *testing.T) {
	a := GenerateOpenInterval(1, 4)
	b := GenerateGreaterThanInterval(2)
	c := Intersect(a, b)
	AssertEqual(len(c.values), 1, t)
	AssertEqual(c.lowerBound.value, 2, t)
	AssertEqual(c.lowerBound.pointType, OpenPoint, t)
	AssertEqual(c.values[0], 3, t)
	AssertEqual(c.upperBound.value, 4, t)
	AssertEqual(c.upperBound.pointType, OpenPoint, t)
	AssertTrue(c.Contains(3), t)
	AssertFalse(c.Contains(1), t)
	AssertFalse(c.Contains(5), t)
	AssertEqual(c.intervalType, OpenInterval, t)
}

/* OpenInterval ∩ AtLeastInterval */
func TestOALIntervalIntersect(t *testing.T) {
	a := GenerateOpenInterval(1, 4)
	b := GenerateAtLeastInterval(2)
	c := Intersect(a, b)
	AssertEqual(len(c.values), 2, t)
	AssertEqual(c.lowerBound.value, 2, t)
	AssertEqual(c.lowerBound.pointType, ClosedPoint, t)
	AssertEqual(c.values[0], 2, t)
	AssertEqual(c.upperBound.value, 4, t)
	AssertEqual(c.upperBound.pointType, OpenPoint, t)
	AssertEqual(c.values[1], 3, t)
	AssertTrue(c.Contains(3), t)
	AssertFalse(c.Contains(1), t)
	AssertFalse(c.Contains(5), t)
	AssertEqual(c.intervalType, ClosedOpenInterval, t)
}

/* OpenInterval ∩ LessThanInterval */
func TestOLTIntervalIntersect(t *testing.T) {
	a := GenerateOpenInterval(1, 4)
	b := GenerateLessThanInterval(2)
	c := Intersect(a, b)
	AssertEqual(len(c.values), 0, t)
	AssertEqual(c.lowerBound.value, 1, t)
	AssertEqual(c.lowerBound.pointType, OpenPoint, t)
	AssertEqual(c.upperBound.value, 2, t)
	AssertEqual(c.upperBound.pointType, OpenPoint, t)
	AssertFalse(c.Contains(1), t)
	AssertFalse(c.Contains(2), t)
	AssertEqual(c.intervalType, OpenInterval, t)
}

/* OpenInterval ∩ AtMostInterval */
func TestOAMIntervalIntersect(t *testing.T) {
	a := GenerateOpenInterval(1, 4)
	b := GenerateAtLeastInterval(2)
	c := Intersect(a, b)
	AssertEqual(len(c.values), 2, t)
	AssertEqual(c.lowerBound.value, 2, t)
	AssertEqual(c.lowerBound.pointType, ClosedPoint, t)
	AssertEqual(c.values[0], 2, t)
	AssertEqual(c.upperBound.value, 4, t)
	AssertEqual(c.upperBound.pointType, OpenPoint, t)
	AssertEqual(c.values[1], 3, t)
	AssertTrue(c.Contains(3), t)
	AssertFalse(c.Contains(1), t)
	AssertFalse(c.Contains(5), t)
	AssertEqual(c.intervalType, ClosedOpenInterval, t)
}

/* OpenInterval ∩ UnboundedInterval */
func TestOUIntervalIntersect(t *testing.T) {
	a := GenerateOpenInterval(1, 4)
	b := GenerateAtLeastInterval(2)
	c := Intersect(a, b)
	AssertEqual(len(c.values), 2, t)
	AssertEqual(c.lowerBound.value, 2, t)
	AssertEqual(c.lowerBound.pointType, ClosedPoint, t)
	AssertEqual(c.values[0], 2, t)
	AssertEqual(c.upperBound.value, 4, t)
	AssertEqual(c.upperBound.pointType, OpenPoint, t)
	AssertEqual(c.values[1], 3, t)
	AssertTrue(c.Contains(3), t)
	AssertFalse(c.Contains(1), t)
	AssertFalse(c.Contains(5), t)
	AssertEqual(c.intervalType, ClosedOpenInterval, t)
}

/* !SECTION: OpenInterval Intersection */

/* SECTION: ClosedInterval Intersection */

/* ClosedInterval ∩ ClosedInterval */
func TestCCIntervalIntersect(t *testing.T) {
	a := GenerateClosedInterval(1, 4)
	b := GenerateClosedInterval(2, 5)
	c := Intersect(a, b)
	AssertEqual(len(c.values), 3, t)
	AssertEqual(c.lowerBound.value, 2, t)
	AssertEqual(c.lowerBound.pointType, ClosedPoint, t)
	AssertEqual(c.values[0], 2, t)
	AssertEqual(c.upperBound.value, 4, t)
	AssertEqual(c.upperBound.pointType, ClosedPoint, t)
	AssertEqual(c.values[2], 4, t)
	AssertTrue(c.Contains(3), t)
	AssertFalse(c.Contains(1), t)
	AssertFalse(c.Contains(5), t)
	AssertEqual(c.intervalType, ClosedInterval, t)
}

/* ClosedInterval ∩ OpenClosedInterval */
func TestCOCIntervalIntersect(t *testing.T) {
	a := GenerateClosedInterval(1, 4)
	b := GenerateOpenClosedInterval(2, 5)
	c := Intersect(a, b)
	AssertEqual(len(c.values), 2, t)
	AssertEqual(c.lowerBound.value, 2, t)
	AssertEqual(c.lowerBound.pointType, OpenPoint, t)
	AssertEqual(c.values[0], 3, t)
	AssertEqual(c.upperBound.value, 4, t)
	AssertEqual(c.upperBound.pointType, ClosedPoint, t)
	AssertEqual(c.values[1], 4, t)
	AssertTrue(c.Contains(3), t)
	AssertFalse(c.Contains(1), t)
	AssertFalse(c.Contains(5), t)
	AssertEqual(c.intervalType, OpenClosedInterval, t)
}

/* ClosedInterval ∩ ClosedOpenInterval */
func TestCCOIntervalIntersect(t *testing.T) {
	a := GenerateClosedInterval(1, 4)
	b := GenerateClosedOpenInterval(2, 5)
	c := Intersect(a, b)
	AssertEqual(len(c.values), 3, t)
	AssertEqual(c.lowerBound.value, 2, t)
	AssertEqual(c.lowerBound.pointType, ClosedPoint, t)
	AssertEqual(c.values[0], 2, t)
	AssertEqual(c.upperBound.value, 4, t)
	AssertEqual(c.upperBound.pointType, ClosedPoint, t)
	AssertEqual(c.values[1], 3, t)
	AssertTrue(c.Contains(3), t)
	AssertFalse(c.Contains(1), t)
	AssertFalse(c.Contains(5), t)
	AssertEqual(c.intervalType, ClosedInterval, t)
}

/* ClosedInterval ∩ GreaterThanInterval */
func TestCGTIntervalIntersect(t *testing.T) {
	a := GenerateClosedInterval(1, 4)
	b := GenerateClosedOpenInterval(2, 5)
	c := Intersect(a, b)
	AssertEqual(len(c.values), 3, t)
	AssertEqual(c.lowerBound.value, 2, t)
	AssertEqual(c.lowerBound.pointType, ClosedPoint, t)
	AssertEqual(c.values[0], 2, t)
	AssertEqual(c.upperBound.value, 4, t)
	AssertEqual(c.upperBound.pointType, ClosedPoint, t)
	AssertEqual(c.values[2], 4, t)
	AssertTrue(c.Contains(3), t)
	AssertFalse(c.Contains(1), t)
	AssertFalse(c.Contains(5), t)
	AssertEqual(c.intervalType, ClosedInterval, t)
}

/* ClosedInterval ∩ AtLeastInterval */
func TestCATIntervalIntersect(t *testing.T) {
	a := GenerateClosedInterval(1, 4)
	b := GenerateAtLeastInterval(2)
	c := Intersect(a, b)
	AssertEqual(len(c.values), 3, t)
	AssertEqual(c.lowerBound.value, 2, t)
	AssertEqual(c.lowerBound.pointType, ClosedPoint, t)
	AssertEqual(c.values[0], 2, t)
	AssertEqual(c.upperBound.value, 4, t)
	AssertEqual(c.upperBound.pointType, ClosedPoint, t)
	AssertEqual(c.values[2], 4, t)
	AssertTrue(c.Contains(3), t)
	AssertFalse(c.Contains(1), t)
	AssertFalse(c.Contains(5), t)
	AssertEqual(c.intervalType, ClosedInterval, t)
}

/* ClosedInterval ∩ LessThanInterval */
func TestCLTIntervalIntersect(t *testing.T) {
	a := GenerateClosedInterval(1, 4)
	b := GenerateLessThanInterval(2)
	c := Intersect(a, b)
	AssertEqual(len(c.values), 1, t)
	AssertEqual(c.lowerBound.value, 1, t)
	AssertEqual(c.lowerBound.pointType, ClosedPoint, t)
	AssertEqual(c.values[0], 1, t)
	AssertEqual(c.upperBound.value, 2, t)
	AssertEqual(c.upperBound.pointType, OpenPoint, t)
	AssertTrue(c.Contains(1), t)
	AssertFalse(c.Contains(5), t)
	AssertEqual(c.intervalType, ClosedOpenInterval, t)
}

/* ClosedInterval ∩ AtMostInterval */
func TestCAMIntervalIntersect(t *testing.T) {
	a := GenerateClosedInterval(1, 4)
	b := GenerateAtMostInterval(2)
	c := Intersect(a, b)
	AssertEqual(len(c.values), 2, t)
	AssertEqual(c.lowerBound.value, 1, t)
	AssertEqual(c.lowerBound.pointType, ClosedPoint, t)
	AssertEqual(c.values[0], 1, t)
	AssertEqual(c.upperBound.value, 2, t)
	AssertEqual(c.upperBound.pointType, ClosedPoint, t)
	AssertEqual(c.values[1], 2, t)
	AssertTrue(c.Contains(1), t)
	AssertFalse(c.Contains(5), t)
	AssertEqual(c.intervalType, ClosedInterval, t)
}

/* ClosedInterval ∩ UnboundedInterval */
func TestCUIntervalIntersect(t *testing.T) {
	a := GenerateClosedInterval(1, 4)
	b := GenerateUnboundedInterval[int]()
	c := Intersect(a, b)
	AssertEqual(len(c.values), 4, t)
	AssertEqual(c.lowerBound.value, 1, t)
	AssertEqual(c.lowerBound.pointType, ClosedPoint, t)
	AssertEqual(c.values[0], 1, t)
	AssertEqual(c.upperBound.value, 4, t)
	AssertEqual(c.upperBound.pointType, ClosedPoint, t)
	AssertEqual(c.values[3], 4, t)
	AssertTrue(c.Contains(1), t)
	AssertFalse(c.Contains(5), t)
	AssertEqual(c.intervalType, ClosedInterval, t)
}

/* !SECTION: ClosedInterval Intersection */

/* SECTION: OpenInterval Intersection */

/* OpenClosedInterval ∩ OpenClosedInterval */
func TestOCOCIntervalIntersect(t *testing.T) {
	a := GenerateOpenClosedInterval(1, 4)
	b := GenerateOpenClosedInterval(2, 5)
	c := Intersect(a, b)
	AssertEqual(len(c.values), 2, t)
	AssertEqual(c.lowerBound.value, 2, t)
	AssertEqual(c.lowerBound.pointType, OpenPoint, t)
	AssertEqual(c.values[0], 3, t)
	AssertEqual(c.upperBound.value, 4, t)
	AssertEqual(c.upperBound.pointType, ClosedPoint, t)
	AssertEqual(c.values[1], 4, t)
	AssertTrue(c.Contains(3), t)
	AssertFalse(c.Contains(5), t)
	AssertEqual(c.intervalType, OpenClosedInterval, t)
}

/* OpenClosedInterval ∩ ClosedOpenInterval */
func TestOCCOIntervalIntersect(t *testing.T) {
	a := GenerateOpenClosedInterval(1, 4)
	b := GenerateClosedOpenInterval(2, 5)
	c := Intersect(a, b)
	AssertEqual(len(c.values), 3, t)
	AssertEqual(c.lowerBound.value, 2, t)
	AssertEqual(c.lowerBound.pointType, ClosedPoint, t)
	AssertEqual(c.values[0], 2, t)
	AssertEqual(c.upperBound.value, 4, t)
	AssertEqual(c.upperBound.pointType, ClosedPoint, t)
	AssertEqual(c.values[2], 4, t)
	AssertTrue(c.Contains(3), t)
	AssertFalse(c.Contains(5), t)
	AssertEqual(c.intervalType, ClosedInterval, t)
}

/* OpenClosedInterval ∩ GreaterThanInterval */
func TestOCGTIntervalIntersect(t *testing.T) {
	a := GenerateOpenClosedInterval(1, 4)
	b := GenerateGreaterThanInterval(2)
	c := Intersect(a, b)
	AssertEqual(len(c.values), 2, t)
	AssertEqual(c.lowerBound.value, 2, t)
	AssertEqual(c.lowerBound.pointType, OpenPoint, t)
	AssertEqual(c.values[0], 3, t)
	AssertEqual(c.upperBound.value, 4, t)
	AssertEqual(c.upperBound.pointType, ClosedPoint, t)
	AssertEqual(c.values[1], 4, t)
	AssertFalse(c.Contains(1), t)
	AssertFalse(c.Contains(5), t)
	AssertEqual(c.intervalType, OpenClosedInterval, t)
}

/* OpenClosedInterval ∩ AtLeastInterval */
func TestOCALIntervalIntersect(t *testing.T) {
	a := GenerateOpenClosedInterval(1, 4)
	b := GenerateAtLeastInterval(2)
	c := Intersect(a, b)
	AssertEqual(len(c.values), 3, t)
	AssertEqual(c.lowerBound.value, 2, t)
	AssertEqual(c.lowerBound.pointType, ClosedPoint, t)
	AssertEqual(c.values[0], 2, t)
	AssertEqual(c.upperBound.value, 4, t)
	AssertEqual(c.upperBound.pointType, ClosedPoint, t)
	AssertEqual(c.values[2], 4, t)
	AssertFalse(c.Contains(1), t)
	AssertFalse(c.Contains(5), t)
	AssertEqual(c.intervalType, ClosedInterval, t)
}

/* OpenClosedInterval ∩ LessThanInterval */
func TestOCLTIntervalIntersect(t *testing.T) {
	a := GenerateOpenClosedInterval(1, 4)
	b := GenerateLessThanInterval(2)
	c := Intersect(a, b)
	AssertEqual(len(c.values), 0, t)
	AssertEqual(c.lowerBound.value, 1, t)
	AssertEqual(c.lowerBound.pointType, OpenPoint, t)
	AssertEqual(c.upperBound.value, 2, t)
	AssertEqual(c.upperBound.pointType, OpenPoint, t)
	AssertFalse(c.Contains(1), t)
	AssertFalse(c.Contains(5), t)
	AssertEqual(c.intervalType, OpenInterval, t)
}

/* OpenClosedInterval ∩ AtMostInterval */
func TestOCAMIntervalIntersect(t *testing.T) {
	a := GenerateOpenClosedInterval(1, 4)
	b := GenerateAtMostInterval(2)
	c := Intersect(a, b)
	AssertEqual(len(c.values), 1, t)
	AssertEqual(c.lowerBound.value, 1, t)
	AssertEqual(c.lowerBound.pointType, OpenPoint, t)
	AssertEqual(c.upperBound.value, 2, t)
	AssertEqual(c.upperBound.pointType, ClosedPoint, t)
	AssertEqual(c.values[0], 2, t)
	AssertFalse(c.Contains(1), t)
	AssertFalse(c.Contains(5), t)
	AssertEqual(c.intervalType, OpenClosedInterval, t)
}

/* OpenClosedInterval ∩ UnboundedInterval */
func TestOCUIntervalIntersect(t *testing.T) {
	a := GenerateOpenClosedInterval(1, 4)
	b := GenerateUnboundedInterval[int]()
	c := Intersect(a, b)
	AssertEqual(len(c.values), 3, t)
	AssertEqual(c.lowerBound.value, 1, t)
	AssertEqual(c.lowerBound.pointType, OpenPoint, t)
	AssertEqual(c.upperBound.value, 4, t)
	AssertEqual(c.upperBound.pointType, ClosedPoint, t)
	AssertEqual(c.values[0], 2, t)
	AssertEqual(c.values[2], 4, t)
	AssertTrue(c.Contains(3), t)
	AssertFalse(c.Contains(1), t)
	AssertEqual(c.intervalType, OpenClosedInterval, t)
}

/* !SECTION: OpenClosedInterval Intersection */

/* SECTION: ClosedOpenInterval Intersection */

/* ClosedOpenInterval ∩ ClosedOpenInterval */
func TestCOCOIntervalIntersect(t *testing.T) {
	a := GenerateClosedOpenInterval(1, 4)
	b := GenerateClosedOpenInterval(2, 5)
	c := Intersect(a, b)
	AssertEqual(len(c.values), 2, t)
	AssertEqual(c.lowerBound.value, 2, t)
	AssertEqual(c.lowerBound.pointType, ClosedPoint, t)
	AssertEqual(c.values[0], 2, t)
	AssertEqual(c.upperBound.value, 4, t)
	AssertEqual(c.upperBound.pointType, OpenPoint, t)
	AssertEqual(c.values[1], 3, t)
	AssertTrue(c.Contains(3), t)
	AssertFalse(c.Contains(5), t)
	AssertEqual(c.intervalType, ClosedOpenInterval, t)
}

/* ClosedOpenInterval ∩ GreaterThanInterval */
func TestCOGTIntervalIntersect(t *testing.T) {
	a := GenerateClosedOpenInterval(1, 4)
	b := GenerateGreaterThanInterval(2)
	c := Intersect(a, b)
	AssertEqual(len(c.values), 1, t)
	AssertEqual(c.lowerBound.value, 2, t)
	AssertEqual(c.lowerBound.pointType, OpenPoint, t)
	AssertEqual(c.values[0], 3, t)
	AssertEqual(c.upperBound.value, 4, t)
	AssertEqual(c.upperBound.pointType, OpenPoint, t)
	AssertTrue(c.Contains(3), t)
	AssertFalse(c.Contains(5), t)
	AssertEqual(c.intervalType, OpenInterval, t)
}

/* ClosedOpenInterval ∩ AtLeastInterval */
func TestCOALIntervalIntersect(t *testing.T) {
	a := GenerateClosedOpenInterval(1, 4)
	b := GenerateAtLeastInterval(2)
	c := Intersect(a, b)
	AssertEqual(len(c.values), 2, t)
	AssertEqual(c.lowerBound.value, 2, t)
	AssertEqual(c.lowerBound.pointType, ClosedPoint, t)
	AssertEqual(c.values[0], 2, t)
	AssertEqual(c.upperBound.value, 4, t)
	AssertEqual(c.upperBound.pointType, OpenPoint, t)
	AssertEqual(c.values[1], 3, t)
	AssertTrue(c.Contains(3), t)
	AssertFalse(c.Contains(5), t)
	AssertEqual(c.intervalType, ClosedOpenInterval, t)
}

/* ClosedOpenInterval ∩ LessThanInterval */
func TestCOLTIntervalIntersect(t *testing.T) {
	a := GenerateClosedOpenInterval(1, 4)
	b := GenerateLessThanInterval(2)
	c := Intersect(a, b)
	AssertEqual(len(c.values), 1, t)
	AssertEqual(c.lowerBound.value, 1, t)
	AssertEqual(c.lowerBound.pointType, ClosedPoint, t)
	AssertEqual(c.values[0], 1, t)
	AssertEqual(c.upperBound.value, 2, t)
	AssertEqual(c.upperBound.pointType, OpenPoint, t)
	AssertFalse(c.Contains(2), t)
	AssertFalse(c.Contains(5), t)
	AssertEqual(c.intervalType, ClosedOpenInterval, t)
}

/* ClosedOpenInterval ∩ AtMostInterval */
func TestCOAMIntervalIntersect(t *testing.T) {
	a := GenerateClosedOpenInterval(1, 4)
	b := GenerateAtMostInterval(2)
	c := Intersect(a, b)
	AssertEqual(len(c.values), 2, t)
	AssertEqual(c.lowerBound.value, 1, t)
	AssertEqual(c.lowerBound.pointType, ClosedPoint, t)
	AssertEqual(c.values[0], 1, t)
	AssertEqual(c.upperBound.value, 2, t)
	AssertEqual(c.upperBound.pointType, ClosedPoint, t)
	AssertEqual(c.values[1], 2, t)
	AssertTrue(c.Contains(1), t)
	AssertFalse(c.Contains(5), t)
	AssertEqual(c.intervalType, ClosedInterval, t)
}

/* ClosedOpenInterval ∩ UnboundedInterval */
func TestCOUIntervalIntersect(t *testing.T) {
	a := GenerateClosedOpenInterval(1, 4)
	b := GenerateClosedOpenInterval(2, 5)
	c := Intersect(a, b)
	AssertEqual(len(c.values), 2, t)
	AssertEqual(c.lowerBound.value, 2, t)
	AssertEqual(c.lowerBound.pointType, ClosedPoint, t)
	AssertEqual(c.values[0], 2, t)
	AssertEqual(c.upperBound.value, 4, t)
	AssertEqual(c.upperBound.pointType, OpenPoint, t)
	AssertEqual(c.values[1], 3, t)
	AssertTrue(c.Contains(3), t)
	AssertFalse(c.Contains(5), t)
	AssertEqual(c.intervalType, ClosedOpenInterval, t)
}

/* !SECTION: ClosedOpenInterval Intersection */

/* SECTION: GreaterThanInterval Intersection */

/* GreaterThanInterval ∩ GreaterThanInterval */
func TestGTGTIntervalIntersect(t *testing.T) {
	a := GenerateGreaterThanInterval(1)
	b := GenerateGreaterThanInterval(2)
	c := Intersect(a, b)
	AssertEqual(len(c.values), 0, t)
	AssertEqual(c.lowerBound.value, 2, t)
	AssertEqual(c.lowerBound.pointType, OpenPoint, t)
	AssertEqual(c.upperBound.value, int(math.Inf(0)), t)
	AssertEqual(c.upperBound.pointType, UnboundedPoint, t)
	AssertFalse(c.Contains(1), t)
	AssertTrue(c.Contains(5), t)
	AssertEqual(c.intervalType, GreaterThanInterval, t)
}

/* GreaterThanInterval ∩ AtLeastInterval */
func TestGTALIntervalIntersect(t *testing.T) {
	a := GenerateGreaterThanInterval(1)
	b := GenerateAtLeastInterval(1)
	c := Intersect(a, b)
	AssertEqual(len(c.values), 0, t)
	AssertEqual(c.lowerBound.value, 1, t)
	AssertEqual(c.lowerBound.pointType, OpenPoint, t)
	AssertEqual(c.upperBound.value, int(math.Inf(0)), t)
	AssertEqual(c.upperBound.pointType, UnboundedPoint, t)
	AssertFalse(c.Contains(1), t)
	AssertTrue(c.Contains(5), t)
	AssertEqual(c.intervalType, GreaterThanInterval, t)
}

/* GreaterThanInterval ∩ LessThanInterval */
func TestGTLTIntervalIntersect(t *testing.T) {
	a := GenerateGreaterThanInterval(1)
	b := GenerateLessThanInterval(2)
	c := Intersect(a, b)
	AssertEqual(len(c.values), 0, t)
	AssertEqual(c.lowerBound.value, 1, t)
	AssertEqual(c.lowerBound.pointType, OpenPoint, t)
	AssertEqual(c.upperBound.value, 2, t)
	AssertEqual(c.upperBound.pointType, OpenPoint, t)
	AssertFalse(c.Contains(1), t)
	AssertFalse(c.Contains(2), t)
	AssertEqual(c.intervalType, OpenInterval, t)
}

/* GreaterThanInterval ∩ AtMostInterval */
func TestGTAMIntervalIntersect(t *testing.T) {
	a := GenerateGreaterThanInterval(1)
	b := GenerateAtMostInterval(2)
	c := Intersect(a, b)
	AssertEqual(len(c.values), 1, t)
	AssertEqual(c.lowerBound.value, 1, t)
	AssertEqual(c.lowerBound.pointType, OpenPoint, t)
	AssertEqual(c.upperBound.value, 2, t)
	AssertEqual(c.upperBound.pointType, ClosedPoint, t)
	AssertFalse(c.Contains(1), t)
	AssertTrue(c.Contains(2), t)
	AssertEqual(c.intervalType, OpenClosedInterval, t)
}

/* GreaterThanInterval ∩ UnboundedInterval */
func TestGTUIntervalIntersect(t *testing.T) {
	a := GenerateGreaterThanInterval(1)
	b := GenerateUnboundedInterval[int]()
	c := Intersect(a, b)
	AssertEqual(len(c.values), 0, t)
	AssertEqual(c.lowerBound.value, 1, t)
	AssertEqual(c.lowerBound.pointType, OpenPoint, t)
	AssertEqual(c.upperBound.value, int(math.Inf(0)), t)
	AssertEqual(c.upperBound.pointType, UnboundedPoint, t)
	AssertFalse(c.Contains(1), t)
	AssertTrue(c.Contains(5), t)
	AssertEqual(c.intervalType, GreaterThanInterval, t)
}

/* !SECTION: GreaterThanInterval Intersection */

/* SECTION: AtLeastInterval Intersection */

/* AtLeastInterval ∩ AtLeastInterval */
func TestALALIntervalIntersect(t *testing.T) {
	a := GenerateAtLeastInterval(2)
	b := GenerateAtLeastInterval(1)
	c := Intersect(a, b)
	AssertEqual(len(c.values), 0, t)
	AssertEqual(c.lowerBound.value, 2, t)
	AssertEqual(c.lowerBound.pointType, ClosedPoint, t)
	AssertEqual(c.upperBound.value, int(math.Inf(0)), t)
	AssertEqual(c.upperBound.pointType, UnboundedPoint, t)
	AssertFalse(c.Contains(1), t)
	AssertTrue(c.Contains(5), t)
	AssertEqual(c.intervalType, AtLeastInterval, t)
}

/* AtLeastInterval ∩ LessThanInterval */
func TestALLTIntervalIntersect(t *testing.T) {
	a := GenerateAtLeastInterval(2)
	b := GenerateLessThanInterval(1)
	c := Intersect(a, b)
	AssertEqual(len(c.values), 0, t)
	AssertEqual(c.lowerBound.value, 0, t)
	AssertEqual(c.lowerBound.pointType, OpenPoint, t)
	AssertEqual(c.upperBound.value, 0, t)
	AssertEqual(c.upperBound.pointType, OpenPoint, t)
	AssertFalse(c.Contains(1), t)
	AssertFalse(c.Contains(2), t)
	AssertEqual(c.intervalType, EmptyInterval, t)
}

/* AtLeastInterval ∩ AtMostInterval */
func TestALAMIntervalIntersect(t *testing.T) {
	a := GenerateAtLeastInterval(2)
	b := GenerateAtMostInterval(1)
	c := Intersect(a, b)
	AssertEqual(len(c.values), 0, t)
	AssertEqual(c.lowerBound.value, 0, t)
	AssertEqual(c.lowerBound.pointType, OpenPoint, t)
	AssertEqual(c.upperBound.value, 0, t)
	AssertEqual(c.upperBound.pointType, OpenPoint, t)
	AssertFalse(c.Contains(1), t)
	AssertFalse(c.Contains(2), t)
	AssertEqual(c.intervalType, EmptyInterval, t)
}

/* AtLeastInterval ∩ UnboundedInterval */
func TestALUIntervalIntersect(t *testing.T) {
	a := GenerateAtLeastInterval(2)
	b := GenerateUnboundedInterval[int]()
	c := Intersect(a, b)
	AssertEqual(len(c.values), 0, t)
	AssertEqual(c.lowerBound.value, 2, t)
	AssertEqual(c.lowerBound.pointType, ClosedPoint, t)
	AssertEqual(c.upperBound.value, int(math.Inf(0)), t)
	AssertEqual(c.upperBound.pointType, UnboundedPoint, t)
	AssertFalse(c.Contains(1), t)
	AssertTrue(c.Contains(2), t)
	AssertEqual(c.intervalType, AtLeastInterval, t)
}

/* !SECTION: AtLeastInterval Intersection */

/* SECTION: LessThanInterval Intersection */

/* LessThanInterval ∩ LessThanInterval */
func TestLTLTIntervalIntersect(t *testing.T) {
	a := GenerateLessThanInterval(2)
	b := GenerateLessThanInterval(1)
	c := Intersect(a, b)
	AssertEqual(len(c.values), 0, t)
	AssertEqual(c.lowerBound.value, int(math.Inf(-1)), t)
	AssertEqual(c.lowerBound.pointType, UnboundedPoint, t)
	AssertEqual(c.upperBound.value, 1, t)
	AssertEqual(c.upperBound.pointType, OpenPoint, t)
	AssertFalse(c.Contains(1), t)
	AssertFalse(c.Contains(2), t)
	AssertEqual(c.intervalType, LessThanInterval, t)
}

/* LessThanInterval ∩ AtMostInterval */
func TestLTAMIntervalIntersect(t *testing.T) {
	a := GenerateLessThanInterval(2)
	b := GenerateAtMostInterval(1)
	c := Intersect(a, b)
	AssertEqual(len(c.values), 0, t)
	AssertEqual(c.lowerBound.value, int(math.Inf(-1)), t)
	AssertEqual(c.lowerBound.pointType, UnboundedPoint, t)
	AssertEqual(c.upperBound.value, 1, t)
	AssertEqual(c.upperBound.pointType, ClosedPoint, t)
	AssertTrue(c.Contains(1), t)
	AssertFalse(c.Contains(2), t)
	AssertEqual(c.intervalType, AtMostInterval, t)
}

/* LessThanInterval ∩ UnboundedInterval */
func TestLTUIntervalIntersect(t *testing.T) {
	a := GenerateLessThanInterval(2)
	b := GenerateUnboundedInterval[int]()
	c := Intersect(a, b)
	AssertEqual(len(c.values), 0, t)
	AssertEqual(c.lowerBound.value, int(math.Inf(-1)), t)
	AssertEqual(c.lowerBound.pointType, UnboundedPoint, t)
	AssertEqual(c.upperBound.value, 2, t)
	AssertEqual(c.upperBound.pointType, OpenPoint, t)
	AssertTrue(c.Contains(1), t)
	AssertFalse(c.Contains(2), t)
	AssertEqual(c.intervalType, LessThanInterval, t)
}

/* !SECTION: LessThanInterval Intersection */

/* SECTION: AtMostInterval Intersection */

/* AtMostInterval ∩ AtMostInterval */
func TestAMAMIntervalIntersect(t *testing.T) {
	a := GenerateAtMostInterval(1)
	b := GenerateAtMostInterval(2)
	c := Intersect(a, b)
	AssertEqual(len(c.values), 0, t)
	AssertEqual(c.lowerBound.value, int(math.Inf(-1)), t)
	AssertEqual(c.lowerBound.pointType, UnboundedPoint, t)
	AssertEqual(c.upperBound.value, 1, t)
	AssertEqual(c.upperBound.pointType, ClosedPoint, t)
	AssertTrue(c.Contains(1), t)
	AssertFalse(c.Contains(2), t)
	AssertEqual(c.intervalType, AtMostInterval, t)
}

/* AtMostInterval ∩ UnboundedInterval */
func TestAMUIntervalIntersect(t *testing.T) {
	a := GenerateAtMostInterval(2)
	b := GenerateUnboundedInterval[int]()
	c := Intersect(a, b)
	AssertEqual(len(c.values), 0, t)
	AssertEqual(c.lowerBound.value, int(math.Inf(-1)), t)
	AssertEqual(c.lowerBound.pointType, UnboundedPoint, t)
	AssertEqual(c.upperBound.value, 2, t)
	AssertEqual(c.upperBound.pointType, ClosedPoint, t)
	AssertTrue(c.Contains(2), t)
	AssertTrue(c.Contains(-10), t)
	AssertEqual(c.intervalType, AtMostInterval, t)
}

/* !SECTION: AtMostInterval Intersection */

/* SECTION: Unbounded Intersection */

/* UnboundedInterval ∩ UnboundedInterval */
func TestUUIntervalIntersect(t *testing.T) {
	a := GenerateUnboundedInterval[int]()
	b := GenerateUnboundedInterval[int]()
	c := Intersect(a, b)
	AssertEqual(len(c.values), 0, t)
	AssertEqual(c.lowerBound.value, int(math.Inf(-1)), t)
	AssertEqual(c.lowerBound.pointType, UnboundedPoint, t)
	AssertEqual(c.upperBound.value, int(math.Inf(0)), t)
	AssertEqual(c.upperBound.pointType, UnboundedPoint, t)
	AssertTrue(c.Contains(1000000), t)
	AssertTrue(c.Contains(-1000000), t)
	AssertEqual(c.intervalType, UnboundedInterval, t)
}

/* !SECTION: Unbounded Intersection */

/* !SECTION: Intersection Testing */
