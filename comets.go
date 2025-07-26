package astronomy

import "fmt"

const CometEpoch = 1986.112

func CometPosition(y int) {
	Y := float64(y) - CometEpoch
	fmt.Println(Y)
}
