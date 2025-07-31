package astronomy

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestEclipticToEquatorial(t *testing.T) {
	d, m, s := EclipticToEquatorial(139, 41, 10.0, 4, 52, 31.0, 2009, 7, 6.0)
	assert.Equal(t, 23, d)
	assert.Equal(t, 26, m)
	assert.Equal(t, 17.0, math.Round(s))
}

func TestMeanObliquity(t *testing.T) {
	d, m, s := MeanObliquity(2009, 7, 6.0)
	assert.Equal(t, 23, d)
	assert.Equal(t, 26, m)
	assert.Equal(t, 17.0, math.Round(s))
}

func TestEquatorialToEcliptic(t *testing.T) {
	dl, hl, ml, db, hb, mb := EquatorialToEcliptic(9, 34, 53.32, 19, 32, 6.01)
	assert.Equal(t, 139, dl)
	assert.Equal(t, 41, hl)
	assert.Equal(t, 9.98120330850952, ml)
	assert.Equal(t, 4, db)
	assert.Equal(t, 52, hb)
	assert.Equal(t, 30.992980767144616, mb)
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

func TestGalacticToEquatorial(t *testing.T) {
	ah, am, as, dh, dm, ds := GalacticToEquatorial(232, 14, 52.0, 51, 7, 20.0)
	assert.Equal(t, 10, ah)
	assert.Equal(t, 21, am)
	assert.Equal(t, 0.0, math.Round(as))
	assert.Equal(t, 10, dh)
	assert.Equal(t, 3, dm)
	assert.Equal(t, 11.0, math.Round(ds))
}
