package astronomy

import (
	"fmt"
	"math"
)

const eg = 279.557208
const og = 283.112438
const e = 0.016705

/*
D is the number of days since the epoch 2010.0
*/
func sunLongitude(D float64) float64 {
	fmt.Println(D)
	l := 360.0/365.242191*D + 360.0/math.Pi*e*(360.0/365.242191*D+eg-og) + eg
	l = normalizeAngle(l)
	fmt.Println(l)
	return l
}

func Heliographic(y, m int, d float64) {
	JD := JulianDay(y, m, d)
	T := (JD - 2415020) / 36525
	D := 84 * T / 60.0
	fmt.Println(D)
}

func CarringtonRotationNumber(y, m int, d float64) int {
	JD := JulianDay(y, m, d)
	CRN := math.Round(1690.0 + (JD-2444235.34)/27.2753)
	return int(CRN)
}

func SunOrbit() {

}

func SunPosition(y, m, d, h, mm int, s float64) (int, int, float64, int, int, float64) {
	D := Days(m, d)
	if y == 2010 {
		D += 365
	}
	if y > 2010 {
		for yy := y; yy < y; yy++ {
			D += 365
		}
	}
	if y < 2010 {
		for yy := y; yy > y; yy-- {
			D -= 365
		}
	}
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
	N := 360.0 / 365.242191 * float64(D)
	N = normalizeAngle(N)
	M := N + eg - og
	if M < 0 {
		M += 360
	}
	Ec := 360.0 / math.Pi * e * math.Sin(M*math.Pi/180.0)
	l := N + Ec + eg
	l = normalizeAngle(l)
	a := math.Atan((math.Sin(l*math.Pi/180.0)*math.Cos(e) - math.Tan(0.0)) / math.Cos(l*math.Pi/180.0))
	ah, am, as := DecimalDegreesToDegreeHourMinute(a * 180 / math.Pi)
	de := math.Asin(math.Sin(0.0)*math.Cos(e) + math.Cos(0.0)*math.Sin(e)*math.Sin(l*math.Pi/180.0))
	dd, dm, ds := DecimalHourToHourMinuteSecond(de)
	return ah, am, as, dd, dm, ds
}

func SunDistance(y, m int, d float64) {
	D := JulianDay(2010, 1, 0.0) - JulianDay(y, m, d)
	M := 360.0/365.242191*D + eg + og
	M = normalizeAngle(M)
	M *= math.Pi / 180.0
	nu := M + 360.0/math.Pi*e*math.Sin(M)
	fmt.Println(nu * 180 / math.Pi)
}

func SunriseSunset() {

}

func Twilight() {

}

func Time(y, m int, s float64) {
	RightAscension(0, 0, 0.0)
}

func SolarEclipse() {

}
