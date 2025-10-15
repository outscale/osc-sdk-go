package iso8601

import (
	"fmt"
	"time"
)

type Time struct {
	time.Time
}

func (t *Time) MarshalJSON() ([]byte, error) {
	// Let's try to aim for a format that is RFC3339 and ISO8601 compatible
	s := fmt.Sprintf("\"%s\"", t.Time.Format("2006-01-02T15:04:05Z"))
	return []byte(s), nil
}

func (t *Time) UnmarshalJSON(data []byte) error {
	// try to be stupid
	if err := t.Time.UnmarshalJSON(data); err == nil {
		return nil
	}

	// Ignore null
	if string(data) == "null" {
		return nil
	}

	// Remove quotes
	if len(data) > 0 && data[0] == '"' && data[len(data)-1] == '"' {
		data = data[1 : len(data)-1]
	} else {
		return ErrNotString
	}

	tmp, err := ParseInLocation(data, time.UTC)
	if err != nil {
		return err
	}

	t.Time = tmp
	return nil
}
