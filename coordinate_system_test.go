package astronomy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecinalDegreest(t *testing.T) {
	assert.Equal(t, 182.52416666666667, DecimalDegrees(182, 31, 27))
}

func TestDecimalDegreesToDegreeHourMinute(t *testing.T) {
	h, m, s := DecimalDegreesToDegreeHourMinute(182.52416666666667)
	assert.Equal(t, 182, h)
	assert.Equal(t, 31, m)
	assert.Equal(t, 27, s)
}

func TestRightAscension(t *testing.T) {
	h, m, s := RightAscension(9, 36, 10)
	assert.Equal(t, 144, h)
	assert.Equal(t, 2, m)
	assert.Equal(t, 33, s)
}
