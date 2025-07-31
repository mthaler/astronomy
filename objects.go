package astronomy

import (
	"fmt"
	"math"
)

func AngleBetween(alh, alm int, als float64, adh, adm int, ads float64, blh, blm int, bls float64, bdh, bdm int, bds float64) (int, int, float64) {
	a1 := DecimalHour(alh, alm, als)
	d1 := DecimalDegrees(adh, adm, ads)
	a2 := DecimalHour(blh, blm, bls)
	d2 := DecimalDegrees(bdh, bdm, bds)
	cosd := math.Sin(d1*math.Pi/180.0)*math.Sin(d2*math.Pi/180.0) + math.Cos(d1*math.Pi/180.0)*math.Cos(d2*math.Pi/180.0)*math.Cos((a1-a2)*math.Pi/180.0*15)
	d := math.Acos(cosd)
	return DecimalDegreesToDegreeHourMinute(d * 180.0 / math.Pi)
}

func RisingSetting(ah, am int, as float64, dd, dm int, ds float64, y, m, d int, la, lo float64) {
	nu := 34.0 / 60.0 // degrees
	a := DecimalHour(ah, am, as)
	a *= 15.0
	de := DecimalDegrees(dd, dm, ds)
	fmt.Println(de)
	cosH := -(math.Sin(nu*math.Pi/180.0) + math.Sin(la*math.Pi/180.0)*math.Sin(lo*math.Pi/180.0)) / (math.Cos(la*math.Pi/180.0) * math.Cos(lo*math.Pi/180.0))
	fmt.Println(cosH)
	H := math.Acos(cosH)
	H /= 15.0
	fmt.Println(H * 180.0 / math.Pi)
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
