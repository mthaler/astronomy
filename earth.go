package astronomy

import "math"

func Nutation(y, m int, d float64) (float64, float64) {
	JD := JulianDay(y, m, d)
	T := (JD - 2415020.0) / 36525.0
	A := 100.002136 * T
	L := 279.6967 + 360.0*(A-float64(int(A))) - 360.0
	B := 5.372617 * T
	O := 259.1833 - 360.0*(B-float64(int(B))) + 360.0
	DP := -17.2*math.Sin(O*math.Pi/180.0) - 1.3*math.Sin(2*L*math.Pi/180.0)
	De := 9.2*math.Cos(O*math.Pi/180.0) + 0.5*math.Cos(2*L*math.Pi/180.0)
	return DP, De
}
