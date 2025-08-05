package astronomy

import (
	"fmt"
	"math"
)

func Precession(ah, am int, as float64, dd, dm int, ds float64) (int, int, float64, int, int, float64) {
	a := DecimalHour(ah, am, as)
	d := DecimalDegrees(dd, dm, ds)
	ad := a * 15.0
	N := 29.5
	S := (3.07327 + 1.33617*math.Sin(ad*math.Pi/180.0)*math.Tan(d*math.Pi/180.0)) * N
	S /= 3600.0
	a1 := S + a
	ah2, am2, as2 := DecimalHourToHourMinuteSecond(a1)
	S2 := 20.0426 * math.Cos(ad*math.Pi/180.0) * N
	S2 /= 3600.0
	d2 := d + S2
	dd2, dm2, ds2 := DecimalDegreesToDegreeHourMinute(d2)
	return ah2, am2, as2, dd2, dm2, ds2
}

func Nutation(y, m int, d float64) (float64, float64) {
	JD := JulianDay(y, m, d)
	T := (JD - 2415020.0) / 36525.0
	A := 100.002136 * T
	L := 279.6967 + 360.0*(A-float64(int(A))) - 360.0
	L = normalizeAngle(L)
	B := 5.372617 * T
	O := 259.1833 - 360.0*(B-float64(int(B))) + 360.0
	DP := -17.2*math.Sin(O*math.Pi/180.0) - 1.3*math.Sin(2*L*math.Pi/180.0)
	De := 9.2*math.Cos(O*math.Pi/180.0) + 0.5*math.Cos(2*L*math.Pi/180.0)
	return DP, De
}

func Aberration(ld, lm int, ls float64, bd, bm int, bs float64) (int, int, float64, int, int, float64) {
	l := DecimalDegrees(ld, lm, ls)
	b := DecimalDegrees(bd, bm, bs)
	solar_longtitude := 165.562250
	Dl := -20.5 * math.Cos((solar_longtitude-l)*math.Pi/180) / math.Cos(b*math.Pi/180)
	Dl /= 3600.0
	Db := -20.5 * math.Sin((solar_longtitude-l)*math.Pi/180) * math.Sin(b*math.Pi/180)
	Db /= 3600.0
	fmt.Println(Db)
	ld2, lm2, ls2 := DecimalDegreesToDegreeHourMinute(l + Dl)
	bd2, bm2, bs2 := DecimalDegreesToDegreeHourMinute(b + Db)
	return ld2, lm2, ls2, bd2, bm2, bs2
}

func Refraction(h, m int, s float64, dd, md int, sd, P, T float64) {
	a := 19.334345
	A := 283.271027
	R := 0.0
	if a > 15.0 {
		R = 0.00452
	} else {
		R = 0.0
	}
	fmt.Println(R)
}
