package astronomy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlanetCoordinates(t *testing.T) {
	ah, am, as, dd, dm, ds := PlanetCoordinates(2003, 11, 22.0, 11.8579, 337.917132, 14.6633, 0.048907, 5.20278, 1.3035, 100.595)
	assert.Equal(t, 11, ah)
	assert.Equal(t, 11, am)
	assert.Equal(t, 14.0, as)
	assert.Equal(t, 6, dd)
	assert.Equal(t, 21, dm)
	assert.Equal(t, 25.0, ds)
}

func TestPlanetPerturbation(t *testing.T) {
	PlanetPerturbation(2003, 11, 22.0, 0.048907, 497.809764, 14.6633, 5.20278)
}
