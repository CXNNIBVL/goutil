package math_test

import (
	"testing"

	"github.com/CXNNIBVL/goutil/math"
)

func TestAbs(t *testing.T) {
	var v_pos, v_neg, v_abs int = 5, -5, 5

	if r := math.Abs(v_pos); r != v_abs {
		t.Errorf("absolute value of %d is %d, got %d", v_pos, v_abs, r)
	}

	if r := math.Abs(v_neg); r != v_abs {
		t.Errorf("absolute value of %d is %d, got %d", v_neg, v_abs, r)
	}
}
