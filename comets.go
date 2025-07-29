package astronomy

import "fmt"

const CometEpoch = 1986.112

func CometPosition(y int, Tp float64) {
	Y := float64(y) - CometEpoch
	fmt.Println(Y)
	Mc := 360.0 * float64(y) / Tp
	fmt.Println(Mc)
}

func Parabolic(y, m, d int) {

}
