package astronomy

import "testing"

func TestSelenographic(t *testing.T) {
	Selenographic(1988, 5, 1.0, 209.12, -3.08)
}

func TestMoonPosition(t *testing.T) {
	MoonPosition(2003, 9, 1)
}

func TestMoonPhases(t *testing.T) {
	MoonPhases()
}

func TestMoonDistance(t *testing.T) {
	MoonDistance()
}
