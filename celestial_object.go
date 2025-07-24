package astronomy

import "fmt"

func AngleBetween(alh, alm int, als float64, adh, adm int, ads float64, blh, blm int, bls float64, bdh, bdm int, bds float64) {
	a1 := DecimalHour(alh, alm, als)
	fmt.Println(a1)
	d1 := DecimalHour(blh, blm, bls)
	fmt.Println(d1)
}
