package iter_test

import (
	"testing"

	"github.com/CXNNIBVL/goutil/constraints"
	"github.com/CXNNIBVL/goutil/iter"
)

type IntervalGeneratorCfg[T constraints.AnyNumber] struct {
	first, last T
	numItems    int
	success     bool
}

func testGeneratorForIntervalGenerator[T constraints.AnyNumber](a, b T, expectedNumItems int, incr ...T) IntervalGeneratorCfg[T] {

	var increm T = 1

	if len(incr) > 0 {
		increm = incr[0]
	}

	var first, last T = a, b - increm

	cache := make([]T, 0)

	for number := range iter.Interval(a, b, increm) {
		cache = append(cache, number)
	}

	gotLen := len(cache)

	if gotLen == 0 {
		if gotLen != expectedNumItems {
			return IntervalGeneratorCfg[T]{0, 0, len(cache), false}
		}

		return IntervalGeneratorCfg[T]{0, 0, len(cache), true}
	}

	rf, rl := cache[0], cache[len(cache)-1]

	if gotLen != expectedNumItems || rf != first || rl != last {
		return IntervalGeneratorCfg[T]{rf, rl + increm, len(cache), false}
	}

	return IntervalGeneratorCfg[T]{0, 0, len(cache), true}
}

func TestGenerateIntervalNormalIntRangeDefaultIncrement(t *testing.T) {
	expected := IntervalGeneratorCfg[int]{
		first: 0, last: 4, numItems: 4, success: true,
	}
	if got := testGeneratorForIntervalGenerator(expected.first, expected.last, expected.numItems); !got.success {
		t.Errorf("expected %+v, got %+v\n", expected, got)
	}
}

func TestGenerateIntervalStartNegativeIntRangeCustomIncrement(t *testing.T) {
	expected := IntervalGeneratorCfg[int]{
		first: -4, last: 0, numItems: 2, success: true,
	}
	if got := testGeneratorForIntervalGenerator(expected.first, expected.last, expected.numItems, 2); !got.success {
		t.Errorf("expected %+v, got %+v\n", expected, got)
	}
}

func TestGenerateIntervalFloatRangeCustomIncrement(t *testing.T) {

	// Test float range with custom increment
	expected := IntervalGeneratorCfg[float64]{
		first: 0, last: 1.5, numItems: 3, success: true,
	}
	if got := testGeneratorForIntervalGenerator(expected.first, expected.last, expected.numItems, 0.5); !got.success {
		t.Errorf("expected %+v, got %+v\n", expected, got)
	}
}

func TestGenerateIntervalFloatRangeOneItemWithIncrementOutOfBounds(t *testing.T) {
	// Test immediate out of bounds increment
	expected := IntervalGeneratorCfg[int]{
		first: 0, last: 10, numItems: 1, success: true,
	}
	if got := testGeneratorForIntervalGenerator(expected.first, expected.last, expected.numItems, 10); !got.success {
		t.Errorf("expected %+v, got %+v\n", expected, got)
	}
}

func TestGenerateIntervalIntRangeOutOfBoundsNoItems(t *testing.T) {
	// Test immediate out of bounds
	expected := IntervalGeneratorCfg[int]{
		first: 11, last: 10, numItems: 0, success: true,
	}
	if got := testGeneratorForIntervalGenerator(expected.first, expected.last, expected.numItems); !got.success {
		t.Errorf("expected %+v, got %+v\n", expected, got)
	}
}
