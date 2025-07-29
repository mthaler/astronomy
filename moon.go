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

func MoonPosition() {

}

func MoonPhases() {

}

func MoonDistance() {

}

func MoonriseMoonset() {

}
