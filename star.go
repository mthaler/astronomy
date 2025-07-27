package astronomy

import "fmt"

const StarEpoch = 1934.008

func BinaryStarOrbit(y, m, d int) {
	Y := float64(y) - StarEpoch
	fmt.Println(Y)
}
