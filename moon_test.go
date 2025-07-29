package astronomy

import "testing"

func TestSelenographic(t *testing.T) {
	Selenographic(1988, 5, 1.0, 209.12, -3.08)
}

func TestMoonPhases(t *testing.T) {
	MoonPhases()
}

func TestMoonDistance(t *testing.T) {
	MoonDistance()
}

func TestMoonriseMoonset(t *testing.T) {
	MoonriseMoonset()
}
