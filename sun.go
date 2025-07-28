package astronomy

import "fmt"

func CarringtonRotationNumber(y, m int, d float64) {
	JD := JulianDay(y, m, d)
	fmt.Println(JD)
}

func SunOrbit() {

}

func SunPosition(y, m, d, h, mm int, s float64) {
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
	fmt.Println(D)
}

func SunDistance() {

}

func SunriseSunset() {

}

func Twilight() {

}

func Time() {

}

func SolarElongations() {

}
