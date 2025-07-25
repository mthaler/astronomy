package astronomy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalizeTime(t *testing.T) {
	assert.Equal(t, 12.0, normalizeTime(12))
	assert.Equal(t, 12.0, normalizeTime(36))
	assert.Equal(t, 12.0, normalizeTime(-12))
}

func TestIsInteger(t *testing.T) {
	assert.Equal(t, false, isInteger(12.5))
	assert.Equal(t, true, isInteger(12.999999999))
}
