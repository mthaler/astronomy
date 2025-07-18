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

func HorizontalToEquatorial(h, m int, s float64, hh, mm int, ss, l float64) (int, int, float64, int, int, float64) {
	A := DecimalDegrees(hh, mm, ss)
	a := DecimalDegrees(h, m, s)
	sind := math.Sin(a*math.Pi/180.0)*math.Sin(l*math.Pi/180.0) + math.Cos(a*math.Pi/180.0)*math.Cos(l*math.Pi/180.0)*math.Cos(A*math.Pi/180.0)
	d := math.Asin(sind)
	cosH := (math.Sin(a*math.Pi/180.0) - math.Sin(l*math.Pi/180.0)*sind) / (math.Cos(l*math.Pi/180.0) * math.Cos(d))
	H := math.Acos(cosH)
	cosA := (math.Sin(A*math.Pi/180.0) - math.Sin(l*math.Pi/180.0)*sind) / (math.Cos(l*math.Pi/180.0) * math.Cos(d))
	A = math.Acos(cosA)
	sinA := math.Sin(A)
	if sinA >= 0 {
		H = 360.0 - H*180.0/math.Pi
	}
	H = H / 15.0
	h2, m2, s2 := DecimalHourToHourMinuteSecond(H * 180.0 / math.Pi)
	h3, m3, s3 := DecimalDegreesToDegreeHourMinute(d * 180.0 / math.Pi)
	return h2, m2, s2, h3, m3, s3
}

func EclipticToEquatorial(y, m int, d float64) (int, int, float64) {
	jd := JulianDay(y, m, d)
	mjd := jd - 2451545.0
	T := mjd / 36525.0
	DE := 46.815*T + 0.0006*T*T - 0.00181*T*T*T
	DE = DE / 3600.0
	e := 23.439292 - DE
	return DecimalDegreesToDegreeHourMinute(e)
}
