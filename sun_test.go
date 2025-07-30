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
	ah, am, as, dd, dm, ds := SunPosition(2003, 7, 27, 0, 0, 0.0)
	assert.Equal(t, 8, ah)
	assert.Equal(t, 23, am)
	assert.Equal(t, 34.0, as)
	assert.Equal(t, 19, dd)
	assert.Equal(t, 21, dm)
	assert.Equal(t, 10.0, ds)
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

func TestSolarElongations(t *testing.T) {
	SolarElongations()
}

func TestSolarEclipse(t *testing.T) {
	SolarEclipse()
}
