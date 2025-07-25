package astronomy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAngleBetween(t *testing.T) {
	AngleBetween(5, 13, 31.7, -8, 13, 30.0, 6, 44, 13.4, -16, 41, 11.0)
}

func TestRisingSetting(t *testing.T) {
	RisingSetting(23, 39, 20.0, 21, 42, 0.0, 2010, 8, 24)
}

func TestPrecession(t *testing.T) {
	Precession(9, 10, 43.0, 14, 23, 25.0)
}

func TestNutation(t *testing.T) {
	DP, De := Nutation(1988, 9, 1.0)
	assert.Equal(t, 5.49291620725019, DP)
	assert.Equal(t, 9.241559684661622, De)
}

func TestAberration(t *testing.T) {
	Aberration(352, 37, 10.1)
}
