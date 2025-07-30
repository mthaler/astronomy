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
	E := kepler(Mc-0.8, 0.9673, Mc)
	fmt.Println(Mc)
	fmt.Println(E)
}

func Parabolic(y, m, d int) {

}
