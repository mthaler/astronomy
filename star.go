package astronomy

import (
	"fmt"
	"math"
)

const StarEpoch = 1934.008

func BinaryStarOrbit(y, m, d int, T, e, a float64) {
	Y := float64(y) - StarEpoch
	M := 360.0 * Y / T
	M = normalizeAngle(M)
	M = M * math.Pi / 180.0
	E0 := 0.86
	E := kepler(E0, e, M)
	nu := 2 * math.Atan(math.Sqrt((1+e)/(1-e))*math.Tan(E/2))
	nu *= 180.0 / math.Pi
	r := a * (1 - e*math.Cos(E))
	fmt.Println(r)
}
