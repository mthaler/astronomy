package astronomy

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecinalDegreest(t *testing.T) {
	assert.Equal(t, 182.52416666666667, DecimalDegrees(182, 31, 27))
}

func TestDecimalDegreesToDegreeHourMinute(t *testing.T) {
	h, m, s := DecimalDegreesToDegreeHourMinute(182.52416666666667)
	S := int(s)
	assert.Equal(t, 182, h)
	assert.Equal(t, 31, m)
	assert.Equal(t, 27, S)
}

func TestRightAscension(t *testing.T) {
	h, m, s := RightAscension(9, 36, 10.2)
	S := math.Round(s)
	assert.Equal(t, 144, h)
	assert.Equal(t, 2, m)
	assert.Equal(t, 33.0, S)
}

func TestEquatorialToHorizontal(t *testing.T) {
	h, m, s, H, M, S := EquatorialToHorizontal(5, 51, 44, 23, 13, 10, 52.0)
	assert.Equal(t, 19, h)
	assert.Equal(t, 20, m)
	assert.Equal(t, 3.642807769706735, s)
	assert.Equal(t, 283, H)
	assert.Equal(t, 16, M)
	assert.Equal(t, 15.69816218970118, S)
}

func TestHorizontalToEquatorial(t *testing.T) {
	h, m, s, h2, m2, s2 := HorizontalToEquatorial(19, 20, 3.64, 283, 16, 15.7, 52.0)
	assert.Equal(t, 5, h)
	assert.Equal(t, 51, m)
	assert.Equal(t, 44.0, math.Round(s))
	assert.Equal(t, 23, h2)
	assert.Equal(t, 13, m2)
	assert.Equal(t, 10.0, math.Round(s2))
}

func TestMeanObliquity(t *testing.T) {
	d, m, s := MeanObliquity(2009, 7, 6.0)
	assert.Equal(t, 23, d)
	assert.Equal(t, 26, m)
	assert.Equal(t, 17.0, math.Round(s))
}

func TestEclipticToEquatorial(t *testing.T) {
	EclipticToEquatorial(139, 41, 10.0, 4, 52, 31.0, 2009, 7, 6.0)
}

func TestEquatorialToEcliptic(t *testing.T) {
	EquatorialToEcliptic(9, 34, 53.32)

}

func TestEquatorialToGalactic(t *testing.T) {
	ld, lm, ls, bd, bm, bs := EquatorialToGalactic(10, 21, 0, 10, 3, 11.0)
	assert.Equal(t, 232, ld)
	assert.Equal(t, 14, lm)
	assert.Equal(t, 52.0, math.Round(ls))
	assert.Equal(t, 51, bd)
	assert.Equal(t, 7, bm)
	assert.Equal(t, 20.0, math.Round(bs))
}
