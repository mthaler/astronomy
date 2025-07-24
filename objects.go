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
	nu := 34.0 / 60.0
	a := DecimalHour(ah, am, as)
	fmt.Println(a)
	de := DecimalDegrees(dd, dm, ds)
	fmt.Println(de)
	cosH := (math.Sin(nu*math.Pi/180.0 + math.Sin(a*math.Pi/180.0)*math.Sin(de*math.Pi/180.0))) / (math.Cos(a*math.Pi/180.0) * math.Cos(de*math.Pi/180.0))
	fmt.Println(cosH)
}

func Precession() {

}
