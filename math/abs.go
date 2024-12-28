package math

import "github.com/CXNNIBVL/goutil/constraints"

// Get the absolute value of a number
func Abs[T constraints.RationalNumber](v T) T {
	if v < 0 {
		return -v
	}
	return v
}
