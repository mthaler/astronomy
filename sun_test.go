package astronomy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarringtonRotationNumber(t *testing.T) {
	assert.Equal(t, 1624, CarringtonRotationNumber(1975, 1, 27.0))
}

func TestSunOrbit(t *testing.T) {

}

func TestSunPosition(t *testing.T) {
	h, m, s := SunPosition(2003, 7, 27, 0, 0, 0.0)
	assert.Equal(t, 8, h)
	assert.Equal(t, 23, m)
	assert.Equal(t, 34, s)
}

func TestSunriseSunset(t *testing.T) {
	SunriseSunset()
}

func TestTwilight(t *testing.T) {
	Twilight()
}

func TestTime(t *testing.T) {
	Time(2010, 7, 27.0)
}

func TestSolarEclipse(t *testing.T) {
	SolarEclipse()
}
