package astronomy

import (
	"fmt"
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
	if y >= 1582 {
		if d >= 10 {
			if d >= 15 {
				A := int(y / 100)
				B = 2 - A + int(A/4)
			}
		}
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
	D := float64(int((C - 122.1) / 365.25))
	E := float64(int(365.25 * D))
	G := float64(int((C - E) / 30.6001))
	d := C - E + F - float64(int(30.6001*G))
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
	T := int(t)
	f := t - float64(T)
	m := f * 60
	s := (m - float64(int(m))) * 60
	return T, int(m), s
}

func LocalTimeToUT(y, m, d, h, mm int, s float64, dstc, o int) (int, int, float64) {
	UT := DecimalHour(h-dstc, mm, s) - float64(o)
	GD := float64(d) + UT/24.0
	fmt.Println(GD)
	JD := JulianDay(y, m, float64(d))
	Y, M, _ := JulianDayToGreenwichCalendarDate(JD)
	fmt.Println(M)
	return Y, M, float64(int(GD))
}

func GST(y, m, d, h, mm int, s float64) (int, int, float64) {
	jd := JulianDay(y, m, float64(d))
	S := jd - 2451545.0
	T := S / 36525.0
	T0 := 6.697374558 + (2400.051336 * T) + 0.000025862*T*T
	T0 = normalize(T0)
	UT := DecimalHour(h, m, s)
	fmt.Println(UT)
	A := UT*1.002737909 + T0
	GST := math.Mod(T0, 24.0) + T0 + A
	return DecimalHourToHourMinuteSecond(GST)
}

func normalize(n float64) float64 {
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
