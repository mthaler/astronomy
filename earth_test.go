package astronomy

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrecession(t *testing.T) {
	ah, am, as, dd, dm, ds := Precession(9, 10, 43.0, 14, 23, 25.0)
	assert.Equal(t, 9, ah)
	assert.Equal(t, 12, am)
	assert.Equal(t, 20, int(math.Round(as)))
	assert.Equal(t, 14, dd)
	assert.Equal(t, 16, dm)
	assert.Equal(t, 8, int(math.Round(ds)))
}

func TestNutation(t *testing.T) {
	DP, De := Nutation(1988, 9, 1.0)
	assert.Equal(t, 5.49291620725019, DP)
	assert.Equal(t, 9.241559684661622, De)
}

func TestAberration(t *testing.T) {
	ld, lm, ls, bd, bm, bs := Aberration(352, 37, 10.1, -1, 32, 56.4)
	assert.Equal(t, 352, ld)
	assert.Equal(t, 37, lm)
	assert.Equal(t, 30.45212807281132, ls)
	assert.Equal(t, -1, bd)
	assert.Equal(t, 32, bm)
	assert.Equal(t, 56.3, bs)
}

func TestRefraction(t *testing.T) {
	Refraction(5, 51, 44.0, 23, 13, 10.0)
}
