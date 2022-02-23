package util

import (
	"fmt"
	"testing"
)

func preFormattedErrorString[T any](expected, got T) string {
	return fmt.Sprintf("expected: %+v got: %+v", expected, got)
}

func AssertTrue(value bool, t *testing.T) {
	t.Run(fmt.Sprintf("AssertTrue : %v", value), func(t *testing.T) {
		if value != true {
			t.Error(preFormattedErrorString(true, value))
		}
	})
}

func AssertFalse(value bool, t *testing.T) {
	t.Run(fmt.Sprintf("AssertFalse - %v", value), func(t *testing.T) {
		if value != false {
			t.Error(preFormattedErrorString(false, value))
		}
	})
}

func AssertEqual[T comparable](value, expected T, t *testing.T) {
	t.Run(fmt.Sprintf("AssertEqual - %v == %v", value, expected), func(t *testing.T) {
		if value != expected {
			t.Error(preFormattedErrorString(expected, value))
		}
	})
}

func AssertNotEqual[T comparable](value, expected T, t *testing.T) {
	t.Run(fmt.Sprintf("AssertNotEqual - %v != %v", value, expected), func(t *testing.T) {
		if value == expected {
			t.Error(preFormattedErrorString(expected, value))
		}
	})
}

func AssertNil[T comparable](value *T, t *testing.T) {
	t.Run(fmt.Sprintf("AssertNil - %v", value), func(t *testing.T) {
		if value != nil {
			t.Error(preFormattedErrorString(nil, value))
		}
	})

}

func AssertNotNil[T comparable](value *T, t *testing.T) {
	t.Run(fmt.Sprintf("AssertNil - %v", value), func(t *testing.T) {
		if value == nil {
			t.Error(preFormattedErrorString("not nil", "nil"))
		}
	})
}
