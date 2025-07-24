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

func MeanObliquity(y, m int, d float64) (int, int, float64) {
	jd := JulianDay(y, m, d)
	mjd := jd - 2451545.0
	T := mjd / 36525.0
	DE := 46.815*T + 0.0006*T*T - 0.00181*T*T*T
	DE = DE / 3600.0
	e := 23.439292 - DE
	return DecimalDegreesToDegreeHourMinute(e)
}

func EclipticToEquatorial(ld, lm int, ls float64, bd, bm int, bs float64, y, m int, d float64) {
	l := DecimalDegrees(ld, lm, ls)
	b := DecimalDegrees(bd, bm, bs)
	e := 23.438055
	sind := math.Sin(b*math.Pi/180.0)*math.Sin(e*math.Pi/180.0) + math.Cos(b*math.Pi/180.0)*math.Sin(e*math.Pi/180.0)*math.Sin(l*math.Pi/180.0)
	fmt.Println(l)
	fmt.Println(b)
	fmt.Println(e)
	fmt.Println(sind)
}

func EquatorialToEcliptic(ah, am int, as float64) {
	a := DecimalHour(ah, am, as)
	fmt.Println(a)
}

func EquatorialToGalactic(ah, am int, as float64, dd, dm int, ds float64) (int, int, float64, int, int, float64) {
	a := DecimalHour(ah, am, as)
	a = a * 15.0
	d := DecimalDegrees(dd, dm, ds)
	sinb := math.Cos(d*math.Pi/180.0)*math.Cos(27.4*math.Pi/180.0)*math.Cos((a-192.25)*math.Pi/180.0) + math.Sin(d*math.Pi/180.0)*math.Sin(27.4*math.Pi/180.0)
	b := math.Asin(sinb)
	y := math.Sin(d*math.Pi/180.0) - math.Sin(b)*math.Sin(27.4*math.Pi/180.0)
	x := math.Cos(d*math.Pi/180.0) * math.Sin(a*math.Pi/180.0-192.25*math.Pi/180.0) * math.Cos(27.4*math.Pi/180.0)
	l := math.Atan(y/x)*180.0/math.Pi + 33.0 + 180.0
	ld, lm, ls := DecimalDegreesToDegreeHourMinute(l)
	bd, bm, bs := DecimalDegreesToDegreeHourMinute(b * 180.0 / math.Pi)
	return ld, lm, ls, bd, bm, bs
}

func GalacticToEquatorial(ld, lm int, ls float64, bd, bm int, bs float64) {
	l := DecimalDegrees(ld, lm, ls)
	b := DecimalDegrees(bd, bm, bs)
	sind := math.Cos(b*math.Pi/180.0)*math.Cos(27.4*math.Pi/180.0)*math.Sin((l-33.0)*math.Pi/180.0) + math.Sin(b*math.Pi/180.0)*math.Sin(27.4*math.Pi/180.0)
	d := math.Asin(sind)
	fmt.Printf("d:%g\n", d*180.0/math.Pi)
	y := math.Cos(b*math.Pi/180.0) * math.Cos((l-33.0)*math.Pi/180.0)
	fmt.Printf("y:%g\n", y)
}
