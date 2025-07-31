package astronomy

import (
	"fmt"
	"math"
)

func Precession(ah, am int, as float64, dd, dm int, ds float64) {
	a := DecimalHour(ah, am, as)
	ad := a * 15
	fmt.Println(ad)
	d := DecimalDegrees(dd, dm, ds)
	N := 29.5
	S := (3.07327 + 1.33617*math.Sin(a*math.Pi/180.0)*math.Tan(d*math.Pi/180.0)) * N
	S = S / 3600.0
	a1 := S + a
	h, m, s := DecimalHourToHourMinuteSecond(a1)
	fmt.Println(a1)
	fmt.Println(h)
	fmt.Println(m)
	fmt.Println(s)
}

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
