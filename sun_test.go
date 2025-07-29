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
	SunPosition(2003, 7, 27, 0, 0, 0)
}

func TestSunriseSunset(t *testing.T) {
	SunriseSunset()
}
