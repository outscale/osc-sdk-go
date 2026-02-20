package iso8601

import "time"

type Format int

const (
	DateTime Format = iota
	DateYear
	DateMonth
	DateDay
)

type Time struct {
	Format Format
	time.Time
}

func (t Time) String() string {
	switch t.Format {
	case DateYear:
		return t.Time.In(time.UTC).Format("2006")
	case DateMonth:
		return t.Time.In(time.UTC).Format("2006-01")
	case DateDay:
		return t.Time.In(time.UTC).Format("2006-01-02")
	default:
		// Let's try to aim for a format that is RFC3339 and ISO8601 compatible
		return t.Time.In(time.UTC).Format("2006-01-02T15:04:05.999Z")
	}
}

func Now() Time {
	return Time{Format: DateTime, Time: time.Now()}
}

func Date(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) Time {
	return Time{Format: DateTime, Time: time.Date(year, month, day, hour, min, sec, nsec, time.UTC)}
}

func Day(year int, month time.Month, day int) Time {
	return Time{Format: DateDay, Time: time.Date(year, month, day, 0, 0, 0, 0, time.UTC)}
}

func Month(year int, month time.Month) Time {
	return Time{Format: DateMonth, Time: time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)}
}

func Year(year int) Time {
	return Time{Format: DateYear, Time: time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)}
}

func (t Time) Add(d time.Duration) Time {
	return Time{Format: t.Format, Time: t.Time.Add(d)}
}

func (t Time) AddDate(y, m, d int) Time {
	return Time{Format: t.Format, Time: t.Time.AddDate(y, m, d)}
}

// At returns the date time at a specific time of day.
func (t Time) At(hour, min, sec, nsec int) Time {
	return Time{Format: DateTime, Time: time.Date(t.Year(), t.Month(), t.Day(), hour, min, sec, nsec, time.UTC)}
}

// MonthStart returns the first day of the month
func (t Time) MonthStart() Time {
	return Time{Format: DateDay, Time: time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.UTC)}
}

// AtStart returns the first time of the month
func (t Time) AtStart() Time {
	return t.MonthStart().At(0, 0, 0, 0)
}

// MonthEnd returns the last day of the month
func (t Time) MonthEnd() Time {
	return t.MonthStart().AddDate(0, 1, -1)
}

// AtEnd returns the last time of the month with a millisecond precision
func (t Time) AtEnd() Time {
	return t.MonthStart().AddDate(0, 1, 0).At(0, 0, 0, int(-time.Millisecond))
}

// Equal checks that two Time structs are equal (same Format, both Time are equal)
func (t Time) Equal(t2 Time) bool {
	if t.Format != t2.Format {
		return false
	}
	return t.Time.Equal(t2.Time)
}

// ParseString parses a year/month/day or datetime string, expecting a time.UTC location.
func ParseString(s string) (Time, error) {
	return ParseInLocationString(s, time.UTC)
}

// ParseInLocationString parses a year/month/day or datetime string into a Time.
func ParseInLocationString(s string, loc *time.Location) (Time, error) {
	t, precision, err := parseInLocation([]byte(s), loc)
	if err != nil {
		return Time{}, err
	}
	var f Format
	switch precision {
	case year:
		f = DateYear
	case month:
		f = DateMonth
	case day:
		f = DateDay
	default:
		f = DateTime
	}
	return Time{Format: f, Time: t}, nil
}

// The following is copied from the Go standard library to implement date range validation logic
// equivalent to the behaviour of Go's time.Parse.

func isLeap(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

// daysInMonth is the number of days for non-leap years in each calendar month starting at 1
var daysInMonth = [...]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func daysIn(m time.Month, year int) int {
	if m == time.February && isLeap(year) {
		return 29
	}
	return daysInMonth[int(m)]
}
