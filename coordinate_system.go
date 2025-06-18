package astronomy

func DecimalDegrees(h, m, s int) float64 {
	S := float64(s) / 60.0
	M := (float64(m) + S) / 60.0
	return float64(h) + M
}

func DecimalDegreesToDegreeHourMinute(d float64) (int, int, int) {
	D := int(d)
	M := (d - float64(int(d))) * 60.0
	S := (M - float64(int(M))) * 60.0
	return D, int(M), int(S)
}

func DecimalDegreesToDeciamlHours(d float64) float64 {
	return d / 15.0
}

func DecimalHoursToDecimalDegrees(d float64) float64 {
	return d * 15
}

func RightAscension(h, m, s int) (int, int, int) {
	D := DecimalDegreesToDeciamlHours(DecimalDegrees(h, m, s))
	h, m, s = DecimalDegreesToDegreeHourMinute(D)
	return h, m, s
}

func EquatorialToHorizontal(h, H, d float64) (int, int, int) {
	return 0, 0, 0
}
