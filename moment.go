package moment

import "time"

// Moment assistant to work with time
type Moment struct {
	t time.Time
}

// String stringify moment
func (m *Moment) String() string {
	return m.t.Format(cDateTimeFormatMySQL)
}

// New return moment instance
func New() *Moment {
	return NewFromTime(time.Now())
}

// NewFromTime create new moment instance from time and return it
func NewFromTime(t time.Time) *Moment {
	return &Moment{
		t: t,
	}
}

// NewFromString create new moment instance from string and return it
func NewFromString(s string) (*Moment, error) {
	var err error

	t, err := time.ParseInLocation(cDateTimeFormatMySQL, s, time.UTC)
	if err != nil {
		return nil, err
	}

	return NewFromTime(t), err
}

// GetTime return moment time
func (m *Moment) GetTime() time.Time {
	return m.t
}

// GetWeekday return time by given weekday
func (m *Moment) GetWeekday(weekDay time.Weekday) *Moment {
	var getWeekDay func(t time.Time) time.Time

	getWeekDay = func(t time.Time) time.Time {
		if t.Weekday() == weekDay {
			return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
		}

		direction := -1
		if weekDay == time.Sunday || t.Weekday() < weekDay && t.Weekday() != time.Sunday {
			direction = 1
		}

		return getWeekDay(time.Date(t.Year(), t.Month(), t.Day()+direction, 0, 0, 0, 0, t.Location()))
	}

	m.t = getWeekDay(m.t)
	return m
}

// GetBeginOf time of begin interval
func (m *Moment) GetBeginOf(interval TimeEntity) *Moment {
	switch interval {
	case CTimeEntityMinute:
		m.t = time.Date(m.t.Year(), m.t.Month(), m.t.Day(), m.t.Hour(), m.t.Minute(), 0, 0, m.t.Location())
	case CTimeEntityHour:
		m.t = time.Date(m.t.Year(), m.t.Month(), m.t.Day(), m.t.Hour(), 0, 0, 0, m.t.Location())
	case CTimeEntityDay:
		m.t = time.Date(m.t.Year(), m.t.Month(), m.t.Day(), 0, 0, 0, 0, m.t.Location())
	case CTimeEntityMonth:
		m.t = time.Date(m.t.Year(), m.t.Month(), 1, 0, 0, 0, 0, m.t.Location())
	case CTimeEntityYear:
		m.t = time.Date(m.t.Year(), 1, 1, 0, 0, 0, 0, m.t.Location())
	}

	return m
}

// GetEndOf time of end interval
func (m *Moment) GetEndOf(interval TimeEntity) *Moment {
	switch interval {
	case CTimeEntityMinute:
		m.t = time.Date(m.t.Year(), m.t.Month(), m.t.Day(), m.t.Hour(), m.t.Minute()+1, 0, 0, m.t.Location())
	case CTimeEntityHour:
		m.t = time.Date(m.t.Year(), m.t.Month(), m.t.Day(), m.t.Hour()+1, 0, 0, 0, m.t.Location())
	case CTimeEntityDay:
		m.t = time.Date(m.t.Year(), m.t.Month(), m.t.Day()+1, 0, 0, 0, 0, m.t.Location())
	case CTimeEntityMonth:
		m.t = time.Date(m.t.Year(), m.t.Month()+1, 1, 0, 0, 0, 0, m.t.Location())
	case CTimeEntityYear:
		m.t = time.Date(m.t.Year()+1, 1, 1, 0, 0, 0, 0, m.t.Location())
	}

	m.t = m.t.Add(-time.Second)
	return m
}
