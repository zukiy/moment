package moment

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMoment_GetWeekday(t *testing.T) {
	var testCases = []struct {
		weekday    time.Weekday
		in, expect string
	}{
		// Monday
		{time.Monday, "2018-01-01 15:23:54", "2018-01-01 00:00:00"},
		{time.Monday, "2018-01-02 00:00:00", "2018-01-01 00:00:00"},
		{time.Monday, "2018-01-03 12:00:00", "2018-01-01 00:00:00"},
		{time.Monday, "2018-01-04 12:00:00", "2018-01-01 00:00:00"},
		{time.Monday, "2018-01-05 12:00:00", "2018-01-01 00:00:00"},
		{time.Monday, "2018-01-06 12:00:00", "2018-01-01 00:00:00"},
		{time.Monday, "2018-01-07 12:00:00", "2018-01-01 00:00:00"},

		// Tuesday
		{time.Tuesday, "2018-01-01 15:23:54", "2018-01-02 00:00:00"},
		{time.Tuesday, "2018-01-02 00:00:00", "2018-01-02 00:00:00"},
		{time.Tuesday, "2018-01-03 12:00:00", "2018-01-02 00:00:00"},
		{time.Tuesday, "2018-01-04 12:00:00", "2018-01-02 00:00:00"},
		{time.Tuesday, "2018-01-05 12:00:00", "2018-01-02 00:00:00"},
		{time.Tuesday, "2018-01-06 12:00:00", "2018-01-02 00:00:00"},
		{time.Tuesday, "2018-01-07 12:00:00", "2018-01-02 00:00:00"},

		// Wednesday
		{time.Wednesday, "2018-01-01 15:23:54", "2018-01-03 00:00:00"},
		{time.Wednesday, "2018-01-02 00:00:00", "2018-01-03 00:00:00"},
		{time.Wednesday, "2018-01-03 12:00:00", "2018-01-03 00:00:00"},
		{time.Wednesday, "2018-01-04 12:00:00", "2018-01-03 00:00:00"},
		{time.Wednesday, "2018-01-05 12:00:00", "2018-01-03 00:00:00"},
		{time.Wednesday, "2018-01-06 12:00:00", "2018-01-03 00:00:00"},
		{time.Wednesday, "2018-01-07 12:00:00", "2018-01-03 00:00:00"},

		// Thursday
		{time.Thursday, "2018-01-01 15:23:54", "2018-01-04 00:00:00"},
		{time.Thursday, "2018-01-02 00:00:00", "2018-01-04 00:00:00"},
		{time.Thursday, "2018-01-03 12:00:00", "2018-01-04 00:00:00"},
		{time.Thursday, "2018-01-04 12:00:00", "2018-01-04 00:00:00"},
		{time.Thursday, "2018-01-05 12:00:00", "2018-01-04 00:00:00"},
		{time.Thursday, "2018-01-06 12:00:00", "2018-01-04 00:00:00"},
		{time.Thursday, "2018-01-07 12:00:00", "2018-01-04 00:00:00"},

		// Friday
		{time.Friday, "2018-01-01 15:23:54", "2018-01-05 00:00:00"},
		{time.Friday, "2018-01-02 00:00:00", "2018-01-05 00:00:00"},
		{time.Friday, "2018-01-03 12:00:00", "2018-01-05 00:00:00"},
		{time.Friday, "2018-01-04 12:00:00", "2018-01-05 00:00:00"},
		{time.Friday, "2018-01-05 12:00:00", "2018-01-05 00:00:00"},
		{time.Friday, "2018-01-06 12:00:00", "2018-01-05 00:00:00"},
		{time.Friday, "2018-01-07 12:00:00", "2018-01-05 00:00:00"},

		// Saturday
		{time.Saturday, "2018-01-01 15:23:54", "2018-01-06 00:00:00"},
		{time.Saturday, "2018-01-02 00:00:00", "2018-01-06 00:00:00"},
		{time.Saturday, "2018-01-03 12:00:00", "2018-01-06 00:00:00"},
		{time.Saturday, "2018-01-04 12:00:00", "2018-01-06 00:00:00"},
		{time.Saturday, "2018-01-05 12:00:00", "2018-01-06 00:00:00"},
		{time.Saturday, "2018-01-06 12:00:00", "2018-01-06 00:00:00"},
		{time.Saturday, "2018-01-07 12:00:00", "2018-01-06 00:00:00"},

		// Sunday
		{time.Sunday, "2018-01-01 15:23:54", "2018-01-07 00:00:00"},
		{time.Sunday, "2018-01-02 00:00:00", "2018-01-07 00:00:00"},
		{time.Sunday, "2018-01-03 12:00:00", "2018-01-07 00:00:00"},
		{time.Sunday, "2018-01-04 12:00:00", "2018-01-07 00:00:00"},
		{time.Sunday, "2018-01-05 12:00:00", "2018-01-07 00:00:00"},
		{time.Sunday, "2018-01-06 12:00:00", "2018-01-07 00:00:00"},
		{time.Sunday, "2018-01-07 12:00:00", "2018-01-07 00:00:00"},
	}

	for _, tc := range testCases {
		moment, _ := NewFromString(tc.in)
		result := moment.GetWeekday(tc.weekday).GetTime()
		assert.Equal(t, getTimeFromString(tc.expect), result)
	}

}

func TestMoment_GetBeginOf(t *testing.T) {
	var testCases = []struct {
		interval   TimeEntity
		in, expect string
	}{
		{CTimeEntityMinute, "2018-08-13 22:45:41", "2018-08-13 22:45:00"},
		{CTimeEntityHour, "2018-08-13 22:45:41", "2018-08-13 22:00:00"},
		{CTimeEntityDay, "2018-08-13 22:45:41", "2018-08-13 00:00:00"},
		{CTimeEntityMonth, "2018-08-13 22:45:41", "2018-08-01 00:00:00"},
		{CTimeEntityYear, "2018-08-13 22:45:41", "2018-01-01 00:00:00"},
	}

	for _, tc := range testCases {
		moment, _ := NewFromString(tc.in)
		result := moment.GetBeginOf(tc.interval).GetTime()
		assert.Equal(t, getTimeFromString(tc.expect), result)
	}
}

func TestMoment_GetEndOf(t *testing.T) {
	var testCases = []struct {
		interval   TimeEntity
		in, expect string
	}{
		{CTimeEntityMinute, "2018-08-13 22:45:41", "2018-08-13 22:45:59"},
		{CTimeEntityHour, "2018-08-13 22:45:41", "2018-08-13 22:59:59"},
		{CTimeEntityDay, "2018-08-13 22:45:41", "2018-08-13 23:59:59"},
		{CTimeEntityMonth, "2018-08-13 22:45:41", "2018-08-31 23:59:59"},
		{CTimeEntityYear, "2018-08-13 22:45:41", "2018-12-31 23:59:59"},
	}

	for _, tc := range testCases {
		moment, _ := NewFromString(tc.in)
		result := moment.GetEndOf(tc.interval).GetTime()
		assert.Equal(t, getTimeFromString(tc.expect), result)
	}
}

func getTimeFromString(s string) time.Time {
	t, _ := time.ParseInLocation(cDateTimeFormatMySQL, s, time.UTC)
	return t
}
