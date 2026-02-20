package comparator

import "math"

func Float64Compare(a, b float64) bool {
	var delta float64 = 1e-9

	if math.Abs(a-b) < delta {
		return true
	} else {
		return false
	}
}
