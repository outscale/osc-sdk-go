package iso8601

import "encoding/json"

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.String() + `"`), nil
}

func (t *Time) UnmarshalJSON(data []byte) error {
	// try to be stupid
	if err := t.Time.UnmarshalJSON(data); err == nil {
		t.Format = DateTime
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

	tmp, err := ParseString(string(data))
	if err != nil {
		return err
	}

	*t = tmp
	return nil
}

var (
	_ json.Marshaler   = Time{}
	_ json.Unmarshaler = (*Time)(nil)
)
