package astronomy

import "testing"

func TestAngleBetween(t *testing.T) {
	AngleBetween(5, 13, 31.7, -8, 13, 30.0, 6, 44, 13.4, -16, 41, 11.0)
}

func TestRisingSetting(t *testing.T) {
	RisingSetting(23, 39, 20.0, 21, 42, 0.0, 2010, 8, 24)
}
