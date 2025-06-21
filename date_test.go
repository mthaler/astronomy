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

func TestJulianDay(t *testing.T) {
	jd := JulianDay(2009, 6, 19.75)
	assert.Equal(t, 2455002.25, jd)
}

func TestDate(t *testing.T) {
	y, m, d := Date(2455002.25)
	assert.Equal(t, 2009, y)
	assert.Equal(t, 6, m)
	assert.Equal(t, 19.75, d)
}

func TestWeekday(t *testing.T) {
	assert.Equal(t, "Friday", Weekday(2009, 6, 19.75))
}

func TestDecimalHour(t *testing.T) {
	assert.Equal(t, 18.524166666666666, DecimalHour(18, 31, 27))
}

func TestDecimalHourToHourMinuteSecond(t *testing.T) {
	h, m, s := DecimalHourToHourMinuteSecond(18.524167)
	S := int(s)
	assert.Equal(t, 18, h)
	assert.Equal(t, 31, m)
	assert.Equal(t, 27, S)
}

func TestGST(t *testing.T) {
	h, m, s := GST(1980, 4, 22, 14, 36, 51.67)
	assert.Equal(t, 4, h)
	assert.Equal(t, 40, m)
	assert.Equal(t, 5.23, s)
}
