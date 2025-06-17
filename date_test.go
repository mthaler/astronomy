package astronomy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEaster(t *testing.T) {
	month, day := Easter(2009)
	assert.Equal(t, 4, month)
	assert.Equal(t, 12, day)
}
