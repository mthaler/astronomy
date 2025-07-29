package astronomy

import "math"

/*
 * To find a solution to Kepler’s equation E − e sin E = M for small values of e. All angles are
 * expressed in radians.
 */
func kepler(E, e, M float64) float64 {
	d := M - e*math.Sin(M)
	ep := 0.000001
	if math.Abs(d) <= ep {
		return d
	} else {
		for {
			DE := d / (1 - e*math.Cos(E))
			E1 := E - DE
			d := E1 - e*math.Sin(E1)
			if math.Abs(d) <= ep {
				return d
			}
		}
	}
}
