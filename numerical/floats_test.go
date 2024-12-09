package numerical

import "testing"

func TestAlmostEqualFloats(t *testing.T) {
	tests := []struct {
		a, b, epsilon float64
		expected      bool
	}{
		{1.0, 1.0, 0.0001, true},        // Exact equality
		{1.0, 1.00005, 0.0001, true},    // Within epsilon
		{1.0, 1.00015, 0.0001, false},   // Outside epsilon
		{-1.0, -1.0, 0.0001, true},      // Exact equality for negatives
		{-1.0, -1.00005, 0.0001, true},  // Within epsilon for negatives
		{-1.0, -1.00015, 0.0001, false}, // Outside epsilon for negatives
		{0.0, 0.0, 0.0001, true},        // Zero comparison
		{0.0, 0.00005, 0.0001, true},    // Near zero within epsilon
		{0.0, 0.00015, 0.0001, false},   // Near zero outside epsilon
	}

	for _, test := range tests {
		if got := almostEqualFloats(test.a, test.b, test.epsilon); got != test.expected {
			t.Errorf("almostEqualFloats(%v, %v, %v) = %v, want %v", test.a, test.b, test.epsilon, got, test.expected)
		}
	}
}
