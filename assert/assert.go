package assert

import "testing"
import "math"

const MaxDelta = 0.01

func ApproxEquals(t *testing.T, a, b float32) {
	delta := math.Abs(float64(a - b))
	if delta > MaxDelta {
		t.Error()
	}
}

func ApproxNotEquals(t *testing.T, a, b float32) {
	delta := math.Abs(float64(a - b))
	if delta < MaxDelta {
		t.Error()
	}
}
