package astronomy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDegreest(t *testing.T) {
	assert.Equal(t, 182.52416666666667, Degrees(182, 31, 27))
}
