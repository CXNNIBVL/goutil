package iter

import (
	"iter"

	"github.com/CXNNIBVL/goutil/constraints"
)

// Generates numbers in the half inclusive interval [a, b[ with an optional increment parameter (default is 1).
//
// The increment feature can also be abused to for an infinite generator. See below for details.
//
// For example:
//
//	// This will yield 0,1,2,3
//	for number := range iter.Interval(0, 4) {
//		...
//	}
//
//	// This will yield -4,-2
//	for number := range iter.Interval(-4, 0, 2) {
//		...
//	}
//	// This will yield -4, -5, -6, -7...-Inf, since iter < 0 is always true
//	for number := range iter.Interval(-4, 0, -1) {
//		...
//	}
func Interval[T constraints.AnyNumber](a, b T, increment ...T) iter.Seq[T] {
	var incr T = 1

	if len(increment) > 0 {
		incr = increment[0]
	}

	return func(yield func(T) bool) {
		for i := a; i < b; i = i + incr {
			if !yield(i) {
				return
			}
		}
	}
}
