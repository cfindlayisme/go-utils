package numerical

import "math"

func AlmostEqualFloats(a, b, epsilon float64) bool {
	return math.Abs(a-b) <= epsilon
}
