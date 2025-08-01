package astronomy

import (
	"fmt"
	"math"
)

const Te = 0.999996
const epe = 99.556772
const oe = 103.2055
const ee = 0.016671
const ae = 0.999985

func PlanetCoordinates(y, m int, d, Tp, ep, o, e, a, i, O float64) {
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
	nue := Me + 360.0/math.Pi*ee*math.Sin(Me*math.Pi/180.0)
	L := nue + oe
	L = normalizeAngle(L)
	R := ae * (1 - ee*ee) / (1 + ee*math.Cos(nue*math.Pi/180.0))
	fmt.Println(R)
	Lp := nup + o
	P := math.Asin(math.Sin((Lp-O)*math.Pi/180.0) * math.Sin(i*math.Pi/180.0))
	fmt.Println(P * 180.0 / math.Pi)
	yy := math.Sin((Lp-O)*math.Pi/180.0) * math.Cos(i*math.Pi/180.0)
	x := math.Cos((Lp - O) * math.Pi / 180.0)
	l := math.Atan(yy/x) + O
	fmt.Println(l)
}

func PlanetDistance() {

}

func PlanetLightTravelTime() {

}

func PlanetPhases() {

}

func PlanetBrightness() {
}
