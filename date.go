package astronomy

import "math"

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
