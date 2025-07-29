package astronomy

import "fmt"

func Selenographic(y, m int, d, lo, la float64) {
	JD := JulianDay(y, m, d)
	T := (JD - 2451545.0) / 36525.0
	O := 125.044522 - 1934.136261*T
	fmt.Println(O)
	F := 93.271910 + 483202.0175*T
	F = normalizeAngle(F)
	fmt.Println(F)
}

func MoonPosition(y, m, d int) {
	D := Days(m, d)
	if y >= 2010 {
		for yy := y; yy > 2010; yy-- {
			if IsLeapYear(yy) {
				D += 366
			} else {
				D += 365
			}
		}
	}
	if y <= 2010 {
		for yy := y; yy < 2010; yy++ {
			if IsLeapYear(yy) {
				D -= 366
			} else {
				D -= 365
			}
		}
	}
	fmt.Println(D)
}

func MoonPhases() {

}

func MoonDistance() {

}
