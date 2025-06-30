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
	h, m, s, H, M, S := EquatorialToHorizontal(5, 51, 44, 23, 13, 10, 52)
	assert.Equal(t, 19, h)
	assert.Equal(t, 20, m)
	assert.Equal(t, 3.642807769706735, s)
	assert.Equal(t, 283, H)
	assert.Equal(t, 16, M)
	assert.Equal(t, 15.69816218970118, S)
}
