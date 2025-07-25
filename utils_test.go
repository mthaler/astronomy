package astronomy

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalizeTime(t *testing.T) {
	assert.Equal(t, math.Pi, normalizeTime(180*math.Pi/180.0))
}
