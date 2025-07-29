package astronomy

import "fmt"

const StarEpoch = 1934.008

func BinaryStarOrbit(y, m, d int, T float64) {
	Y := float64(y) - StarEpoch
	M := 360.0 * Y / T
	fmt.Println(M)
}
