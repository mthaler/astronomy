package astronomy

func Degrees(d, m, s int) float64 {
	S := float64(s) / 60.0
	M := (float64(m) + S) / 60.0
	return float64(d) + M
}
