package astronomy

func normalizeTime(n float64) float64 {
	r := n
	for {
		if r > 24.0 {
			r = r - 24.0
		}
		if r < 0.0 {
			r = r + 24.0
		}
		if r >= 0.0 && r <= 24.0 {
			break
		}
	}
	return r
}
