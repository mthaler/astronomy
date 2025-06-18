package astronomy

func DecimalDegrees(d, m, s int) float64 {
	S := float64(s) / 60.0
	M := (float64(m) + S) / 60.0
	return float64(d) + M
}

func DecimalDegreesToDegreeHourMinute(d float64) (int, int, int) {
	D := int(d)
	M := (d - float64(int(d))) * 60.0
	S := (M - float64(int(M))) * 60.0
	return D, int(M), int(S)
}

func EquatorialToHorizontal(h, H, d float64) (int, int, int) {
	return 0, 0, 0
}
