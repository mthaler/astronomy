package astronomy

import "fmt"

func PlanetCoordinates(y, m int, d, Tp, ep, o float64) {
	D := JulianDay(y, m, d) - JulianDay(2010, 1, 0.0)
	Np := 360.0 / 365.242191 * D / Tp
	Np = normalizeAngle(Np)
	Mp := Np + ep + o
	fmt.Println(Mp)
}

func PlanetDistance() {

}

func PlanetLightTravelTime() {

}

func PlanetPhases() {

}

func PlanetBrightness() {
}
