package astronomy

import (
	"math"
)

func Easter(y int) (int, int) {
	a := y % 19
	b := y / 100
	c := y % 100
	d := b / 4
	e := b % 4
	f := (b + 8) / 25
	g := (b - f + 1) / 3
	h := (19*a + b - d - g + 15) % 30
	i := c / 4
	k := c % 4
	l := (32 + 2*e + 2*i - h - k) % 7
	m := (a + 11*h + 22*l) / 451
	n := (h + l - 7*m + 114) / 31
	p := (h + l - 7*m + 114) % 31

	return n, p + 1
}

func JulianDay(year, month int, day float64) float64 {
	y := year
	m := month
	d := day
	if m == 1 || m == 2 {
		y = year - 1
		m = month + 12
	}
	B := 0
	if year > 1582 {
		A := int(y / 100)
		B = 2 - A + int(A/4)
		if m >= 10 {
			if d >= 15 {

			}
		}
	} else if year == 1582 && m > 10 {
		A := int(y / 100)
		B = 2 - A + int(A/4)
	} else if year == 1582 && m == 10 && d >= 15 {
		A := int(y / 100)
		B = 2 - A + int(A/4)
	}
	C := 0
	if y < 0 {
		C = int(math.Trunc((365.25 * float64(y)) - 0.75))
	} else {
		C = int(math.Trunc(365.25 * float64(y)))
	}
	D := int(30.6001 * float64(m+1))
	JD := float64(B+C+D) + d + 1720994.5

	return JD
}

func JulianDayToGreenwichCalendarDate(jd float64) (int, int, float64) {
	I := int(jd + 0.5)
	F := jd + 0.5 - float64(I)
	A := 0.0
	B := float64(I)
	if I > 2299160 {
		A = float64(int((B - 1867216.25) / 36524.25))
		B = B + A - float64(int(A/4.0)) + 1

	}
	C := B + 1524.0
	D := math.Floor(((C - 122.1) / 365.25))
	E := math.Floor(365.25 * D)
	G := math.Floor((C - E) / 30.6001)
	d := C - E + F - math.Floor(30.6001*G)
	m := G - 1.0
	if G >= 13.5 {
		m = G - 13.0
	}
	y := D - 4715.0
	if m > 2.5 {
		y = D - 4716.0
	}
	return int(y), int(m), d
}

func Weekday(year, month int, day float64) string {
	d := int(day)
	jd := JulianDay(year, month, float64(d))
	A := (jd + 1.5) / 7
	B := int(A)
	C := int(math.Round((A - float64(B)) * 7.0))
	switch C {
	case 0:
		return "Sunday"
	case 1:
		return "Monday"
	case 2:
		return "Tuesday"
	case 3:
		return "Wednesday"
	case 4:
		return "Thursday"
	case 5:
		return "Friday"
	case 6:
		return "Saturday"
	default:
		return "undefined"
	}
}

func DecimalHour(h, m int, s float64) float64 {
	S := s / 60.0
	M := (float64(m) + S) / 60.0
	H := float64(h) + M
	return H
}

func DecimalHourToHourMinuteSecond(t float64) (int, int, float64) {
	h, f := math.Modf(t)

	if isInteger(f) {
		f = math.Round(f)
	}
	f = f * 60.0
	m, s := math.Modf(f)
	if isInteger(s) {
		s = math.Round(s)
	}
	s = s * 60.0
	if isInteger(s) && int(s)%60 == 0 {
		n := int(s) / 60
		m += float64(n)
		s = 0.0
	}
	return int(h), int(m), s
}

func LocalTimeToUT(y, m, d, h, mm int, s float64, dstc, o int) (int, int, int, float64) {
	hh := h - dstc
	UT := DecimalHour(hh, mm, s)
	UT -= float64(o)
	GD := float64(d) + UT/24.0
	JD := JulianDay(y, m, GD)
	Y, M, D := JulianDayToGreenwichCalendarDate(JD)
	DD, _ := math.Modf(D)
	t := (D - DD) * 24.0
	return Y, M, int(D), t
}

func UTToLocalTime(y, m, d, h, mm int, s float64, dstc, o int) (int, int, int, int, int, float64) {
	lct := DecimalHour(h, mm, s)
	lct += float64(o + dstc)
	ljd := JulianDay(y, m, float64(d)) + lct/24.0
	yy, mm, dd := JulianDayToGreenwichCalendarDate(ljd)
	ddd, f := math.Modf(dd)
	f *= 24.0
	hh, mmm, ss := DecimalHourToHourMinuteSecond(f)
	return yy, mm, int(ddd), hh, mmm, ss
}

func GST(y, m, d, h, mm int, s float64) (int, int, float64) {
	jd := JulianDay(y, m, float64(d))
	S := jd - 2451545.0
	T := S / 36525.0
	T0 := 6.697374558 + (2400.051336 * T) + 0.000025862*T*T
	T0 = normalizeTime(T0)
	UT := DecimalHour(h, mm, s)
	A := UT * 1.002737909
	GST := normalizeTime(T0 + A)
	return DecimalHourToHourMinuteSecond(GST)
}

func GSTToUT(y, m, d, h, mm int, s float64) (int, int, float64) {
	jd := JulianDay(y, m, float64(d))
	S := jd - 2451545.0
	T := S / 36525.0
	T0 := 6.697374558 + (2400.051336 * T) + 0.000025862*T*T
	T0 = normalizeTime(T0)
	GST := DecimalHour(h, mm, s)
	B := normalizeTime(GST - T0)
	B *= 0.9972695663
	return DecimalHourToHourMinuteSecond(B)
}

func LST(h, m int, s, l float64) (int, int, float64) {
	GST := DecimalHour(h, m, s)
	d := l
	d /= 15
	GST = GST + d
	GST = normalizeTime(GST)
	return DecimalDegreesToDegreeHourMinute(GST)
}

func Days(m, d int) int {
	dd := 0
	for i := 0; i < m-1; i++ {
		dd += days(i + 1)
	}
	dd += d
	return dd
}

func days(m int) int {
	switch m {
	case 1:
		return 31
	case 2:
		return 28
	case 3:
		return 31
	case 4:
		return 30
	case 5:
		return 31
	case 6:
		return 30
	case 7:
		return 31
	case 8:
		return 31
	case 9:
		return 30
	case 10:
		return 31
	case 11:
		return 30
	case 12:
		return 31
	default:
		return 0
	}
}

func IsLeapYear(y int) bool {
	return ((y%4 == 0) && (y%100 != 0) || (y%400 == 0))
}
