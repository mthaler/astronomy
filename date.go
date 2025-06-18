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

func Date(jd float64) (int, int, float64) {
	jd = jd + 0.5
	I := int(jd)
	F := jd - float64(I)
	B := I
	if I >= 2299160 {
		A := int((float64(I) - 1867216.25) / 36524.25)
		B = I + A - A/4 + 1
	}
	C := B + 1524
	D := int((float64(C) - 122.1) / 365.25)
	E := int(365.25 * float64(D))
	G := int(float64(C-E) / 30.6001)
	d := float64(C) - float64(E) + F - float64(int(30.6001*float64(G)))
	m := G - 1
	if float64(G) >= 13.5 {
		m = G - 13
	}
	y := D - 4716
	if float64(m) < 2.5 {
		y = D - 4715
	}
	return y, m, d
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

func DecimalHour(h, m, s int) float64 {
	S := float64(s) / 60.0
	M := (float64(m) + S) / 60.0
	return float64(h) + M
}

func DecimalHourToHourMinuteSecond(t float64) (int, int, int) {
	T := int(t)
	f := t - float64(T)
	m := f * 60
	s := (m - float64(int(m))) * 60
	return T, int(m), int(s)
}

func GST(year, month int, day float64) {
	jd := JulianDay(year, month, float64(int(day)))
	S := jd - 2451545.0
	T := S / 36525.0
	T0 := math.Mod(6.697374558+(2400.051336*T)+0.000025862+T*T, 24.0)
	d := DecimalHour(year, month, int(day))
}
