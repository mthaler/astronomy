package astronomy

import (
	"fmt"
	"math"
)

const StarEpoch = 1934.008

func BinaryStarOrbit(y, m, d int, T float64) {
	Y := float64(y) - StarEpoch
	M := 360.0 * Y / T
	M = normalizeAngle(M)
	M = M * math.Pi / 180.0
	fmt.Println(M)
}
