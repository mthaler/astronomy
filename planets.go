package astronomy

import (
	"fmt"
	"math"
)

const Te = 0.999996
const epe = 99.556772
const oe = 103.2055
const ee = 0.016671

func PlanetCoordinates(y, m int, d, Tp, ep, o, e, a float64) {
	D := JulianDay(y, m, d) - JulianDay(2010, 1, 0.0)
	Np := 360.0 / 365.242191 * D / Tp
	Np = normalizeAngle(Np)
	Mp := Np + ep - o
	nup := Mp + 360.0/math.Pi*e*math.Sin(Mp*math.Pi/180.0)
	nup = normalizeAngle(nup)
	r := a * (1 - e*e) / (1 + e*math.Cos(nup*math.Pi/180.0))
	fmt.Println(r)
	Ne := 360.0 / 365.242191 * D / Te
	Ne = normalizeAngle(Ne)
	Me := Ne + epe - oe
	Me = normalizeAngle(Me)
	fmt.Println(Me)

}

func PlanetDistance() {

}

func PlanetLightTravelTime() {

}

func PlanetPhases() {

}

func PlanetBrightness() {
}
