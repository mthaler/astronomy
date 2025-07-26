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
	jd = JulianDay(2013, 7, 0.942361)
	assert.Equal(t, 2456474.442361, jd)
}

func TestJulianDayToGreenwichCalendarDate(t *testing.T) {
	y, m, d := JulianDayToGreenwichCalendarDate(2455002.25)
	assert.Equal(t, 2009, y)
	assert.Equal(t, 6, m)
	assert.Equal(t, 19.75, d)
}

func TestWeekday(t *testing.T) {
	assert.Equal(t, "Friday", Weekday(2009, 6, 19.75))
}

func TestDecimalHour(t *testing.T) {
	assert.Equal(t, 18.524166666666666, DecimalHour(18, 31, 27))
	assert.Equal(t, 14.614352777777778, DecimalHour(14, 36, 51.67))
}

func TestDecimalHourToHourMinuteSecond(t *testing.T) {
	h, m, s := DecimalHourToHourMinuteSecond(18.524167)
	S := int(s)
	assert.Equal(t, 18, h)
	assert.Equal(t, 31, m)
	assert.Equal(t, 27, S)
	h, m, s = DecimalHourToHourMinuteSecond(10.349995)
	assert.Equal(t, 10, h)
	assert.Equal(t, 21, m)
	assert.Equal(t, 0, S)
}

func TestLocalTimeToUT(t *testing.T) {
	y, m, d, tt := LocalTimeToUT(2013, 7, 1, 3, 37, 0.0, 1, 4)
	assert.Equal(t, 2013, y)
	assert.Equal(t, 6, m)
	assert.Equal(t, 30, d)
	assert.Equal(t, 22.61666665971279, tt)

}

func TestGST(t *testing.T) {
	h, m, s := GST(1980, 4, 22, 14, 36, 51.67)
	assert.Equal(t, 4, h)
	assert.Equal(t, 40, m)
	assert.Equal(t, 5.229576759185761, s)
}

func TestDays(t *testing.T) {
	Days(7, 27)
}

func TestIsLeapYear(t *testing.T) {

}
