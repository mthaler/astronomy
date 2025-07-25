package astronomy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalizeTime(t *testing.T) {
	assert.Equal(t, 12.0, normalizeTime(12))
	assert.Equal(t, 12.0, normalizeTime(36))
}
