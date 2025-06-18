package astronomy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecinalDegreest(t *testing.T) {
	assert.Equal(t, 182.52416666666667, DecimalDegrees(182, 31, 27))
}

func TestDecimalDegreesToDegreeHourMinute(t *testing.T) {
	d, m, s := DecimalDegreesToDegreeHourMinute(182.52416666666667)
	assert.Equal(t, 182, d)
	assert.Equal(t, 31, m)
	assert.Equal(t, 27, s)
}
