package numerical

import "math"

func almostEqualFloats(a, b, epsilon float64) bool {
	return math.Abs(a-b) <= epsilon
}
