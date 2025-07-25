package astronomy

import (
	"fmt"
	"math"
)

func AngleBetween(alh, alm int, als float64, adh, adm int, ads float64, blh, blm int, bls float64, bdh, bdm int, bds float64) {
	a1 := DecimalHour(alh, alm, als)
	fmt.Println(a1)
	d1 := DecimalHour(adh, adm, ads)
	fmt.Println(d1)
	a2 := DecimalHour(blh, blm, bls)
	fmt.Println(a2)
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

func Nutation() {

}

func Aberration() {

}

func Refraction() {

}

func Parallax() {

}

func CarringtonRotation() {

}

func AtmosphericExtinction() {

}
