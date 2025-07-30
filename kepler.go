package astronomy

import (
	"math"
)

/*
 * To find a solution to Kepler’s equation E − e sin E = M for small values of e. All angles are
 * expressed in radians.
 */
func kepler(E0, e, M float64) float64 {
	ep := 0.000001
	E := E0
	for {
		d := E - e*math.Sin(E) - M
		if math.Abs(d) <= ep {
			return E
		} else {
			DE := d / (1 - e*math.Cos(E))
			E = E - DE
		}
	}
}
