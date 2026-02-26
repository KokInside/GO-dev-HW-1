package comparator

import "math"

func Float64Compare(a, b float64) bool {
	delta := 1e-9

	return math.Abs(a-b) < delta

}
