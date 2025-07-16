package astronomy

import (
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

func EquatorialToHorizontal(h, m, s, hh, mm, ss int, l float64) (int, int, float64, int, int, float64) {
	H := DecimalHour(h, m, float64(s))
	H = H * 15.0
	d := DecimalDegrees(hh, mm, float64(ss))
	sina := math.Sin(d*math.Pi/180.0)*math.Sin(float64(l)*math.Pi/180.0) + math.Cos(d*math.Pi/180.0)*math.Cos(float64(l)*math.Pi/180.0)*math.Cos(H*math.Pi/180.0)
	a := math.Asin(sina) * 180.0 / math.Pi
	cosA := (math.Sin(d*math.Pi/180.0) - math.Sin(float64(l)*math.Pi/180.0)*sina) / (math.Cos(float64(l)*math.Pi/180.0) * math.Cos(a/180.0*math.Pi))
	A := math.Acos(cosA) * 180 / math.Pi
	sinH := math.Sin(H * math.Pi / 180.0)
	if sinH > 0 {
		A = 360.0 - A
	}
	hhh, mmm, sss := DecimalDegreesToDegreeHourMinute(a)
	HH, M, S := DecimalHourToHourMinuteSecond(A)
	return hhh, mmm, sss, HH, M, S
}

func HorizontalToEquatorial(h, m int, s float64, hh, mm int, ss, l float64) (int, int, float64) {
	A := DecimalDegrees(hh, mm, ss)
	a := DecimalDegrees(h, m, s)
	sind := math.Sin(a*math.Pi/180.0)*math.Sin(l*math.Pi/180.0) + math.Cos(a*math.Pi/180.0)*math.Cos(l*math.Pi/180.0)*math.Cos(A*math.Pi/180.0)
	d := math.Asin(sind)
	sinH := (math.Sin(a*math.Pi/180.0) - math.Sin(l*math.Pi/180.0)*sind) / (math.Cos(l*math.Pi/180.0) * math.Cos(d))
	sinA := math.Sin(A * math.Pi / 180.0)
	A = math.Asin(sinA)
	if sinH >= 0 {
		A = 360.0 - A*180.0/math.Pi
	}
	h2, m2, s2 := DecimalHourToHourMinuteSecond(a)
	return h2, m2, s2
}

func EclipticToEquatorial() {

}
