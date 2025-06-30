package astronomy

import (
	"fmt"
	"math"
)

func DecimalDegrees(h, m int, s float64) float64 {
	S := s / 60.0
	M := (float64(m) + S) / 60.0
	return float64(h) + M
}

func DecimalDegreesToDegreeHourMinute(d float64) (int, int, float64) {
	D := int(d)
	M := (d - float64(int(d))) * 60.0
	S := (M - float64(int(M))) * 60.0
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

func EquatorialToHorizontal(h, m, s, hh, mm, ss, l int) (int, int, int) {
	H := DecimalHour(h, m, float64(s))
	H = H * 15.0
	d := DecimalDegrees(hh, mm, float64(ss))
	sina := math.Sin(d*math.Pi/180.0)*math.Sin(float64(l)*math.Pi/180.0) + math.Cos(d*math.Pi/180.0)*math.Cos(float64(l)*math.Pi/180.0)*math.Cos(H*math.Pi/180.0)
	a := math.Asin(sina) * 180.0 / math.Pi
	fmt.Println(a)
	return 0, 0, 0
}
