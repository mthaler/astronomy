package astronomy

import (
	"fmt"
	"math"
)

/*
 * To find a solution to Kepler’s equation E − e sin E = M for small values of e. All angles are
 * expressed in radians.
 */
func kepler(E0, e, M float64) float64 {
	E := E0
	for {
		d := E - e*math.Sin(M) - M
		ep := 0.000001
		if math.Abs(d) <= ep {
			return E
		} else {
			fmt.Println(E)
			for {
				DE := d / (1 - e*math.Cos(E0))
				E = E - DE
			}
		}
	}

}
