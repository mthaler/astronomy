package astronomy

import (
	"fmt"
	"math"
)

func AngleBetween(alh, alm int, als float64, adh, adm int, ads float64, blh, blm int, bls float64, bdh, bdm int, bds float64) {
	a1 := DecimalHour(alh, alm, als)
	fmt.Println(a1)
	d1 := DecimalDegrees(adh, adm, ads)
	fmt.Println(d1)
	a2 := DecimalHour(blh, blm, bls)
	fmt.Println(a2)
	d2 := DecimalDegrees(bdh, bdm, bds)
	fmt.Println(d2)

}

func RisingSetting(ah, am int, as float64, dd, dm int, ds float64, y, m, d int) {
	nu := 34.0 / 60.0 // degrees
	a := DecimalHour(ah, am, as)
	fmt.Println(a)
	de := DecimalDegrees(dd, dm, ds)
	fmt.Println(de)
	cosH := (math.Sin(nu*math.Pi/180.0 + math.Sin(a*math.Pi/180.0)*math.Sin(de*math.Pi/180.0))) / (math.Cos(a*math.Pi/180.0) * math.Cos(de*math.Pi/180.0))
	fmt.Println(cosH)
}

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

func Aberration(ld, lm int, ls float64, bd, bm int, bs float64, sd, sm int, ss float64) {
	l := DecimalDegrees(ld, lm, ls)
	b := DecimalDegrees(bd, bm, bs)
	s := DecimalDegrees(sd, sm, ss)
	fmt.Println(l)
	fmt.Println(b)
	fmt.Println(s)
}

func Refraction(h, m int, s float64, dd, md int, sd float64) {
	a := DecimalHour(h, m, s)
	fmt.Println(a)
}

func Parallax() {

}

func CarringtonRotation() {

}
