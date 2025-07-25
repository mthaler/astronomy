package astronomy

import "math"

func normalizeTime(n float64) float64 {
	r := n
	for {
		if r > 24.0 {
			r = r - 24.0
		}
		if r < 0.0 {
			r = r + 24.0
		}
		if r >= 0.0 && r <= 24.0 {
			break
		}
	}
	return r
}

func isInteger(x float64) bool {
	return math.Abs(x-math.Round(x)) < 0.001
}
