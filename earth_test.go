package astronomy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrecession(t *testing.T) {
	Precession(9, 10, 43.0, 14, 23, 25.0)
}

func TestNutation(t *testing.T) {
	DP, De := Nutation(1988, 9, 1.0)
	assert.Equal(t, 5.49291620725019, DP)
	assert.Equal(t, 9.241559684661622, De)
}
