package astronomy

import "fmt"

func SunOrbit() {

}

func SunPosition(y, m, d, h, mm int, s float64) {
	jd := JulianDay(0, m, float64(d))
	fmt.Println(jd)
}

func SunDistance() {

}

func SunriseSunset() {

}
