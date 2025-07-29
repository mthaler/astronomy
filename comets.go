package astronomy

import (
	"fmt"
	"math"
)

const CometEpoch = 1986.112

func CometPosition(y int, Tp float64) {
	Y := float64(y) - CometEpoch
	Mc := 360.0 * Y / Tp
	Mc = normalizeAngle(Mc)
	Mc = Mc * math.Pi / 180.0
	fmt.Println(Mc)
}

func Parabolic(y, m, d int) {

}
