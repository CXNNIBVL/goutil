package constraints

import "golang.org/x/exp/constraints"

// Defines the set of natural numbers
type NaturalNumber interface {
	constraints.Unsigned
}

// Defines the set of all integers
type IntegerNumber interface {
	NaturalNumber | constraints.Signed
}

// Defines the set of all rational numbers
type RationalNumber interface {
	IntegerNumber | constraints.Float
}

// Defines any number
type AnyNumber interface {
	RationalNumber
}
