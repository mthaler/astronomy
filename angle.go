package astronomy

import "math"

func DecimalDegrees(h, m int, s float64) float64 {
	S := s / 60.0
	M := (float64(m) + float64(S)) / 60.0
	if h > 0 {
		return float64(h) + M
	} else {
		return float64(h) - M
	}
}

func DecimalDegreesToDegreeHourMinute(d float64) (int, int, float64) {
	D := int(d)
	M := (d - math.Floor(d)) * 60.0
	S := (M - math.Floor(M)) * 60.0
	return D, int(M), S
}

func DecimalDegreesToDeciamlHours(d float64) float64 {
	return d / 15.0
}

func DecimalHoursToDecimalDegrees(d float64) float64 {
	return d * 15
}

func RightAscension(h, m int, s float64) (int, int, float64) {
	D := DecimalHoursToDecimalDegrees(DecimalDegrees(h, m, s))
	H, M, S := DecimalDegreesToDegreeHourMinute(D)
	return H, M, S
}

func HourAngle(h, m int, s float64) (int, int, float64) {
	D := DecimalDegreesToDeciamlHours(DecimalDegrees(h, m, s))
	H, M, S := DecimalDegreesToDegreeHourMinute(D)
	return H, M, S
}
