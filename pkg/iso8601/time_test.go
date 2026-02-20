package iso8601_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/outscale/osc-sdk-go/v3/pkg/iso8601"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	tcs = []struct {
		json string
		str  string
		t    iso8601.Time
		err  bool
	}{
		{
			json: `"2026"`,
			str:  "2026",
			t:    iso8601.Time{Format: iso8601.DateYear, Time: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)},
		},
		{
			json: `"2026-02"`,
			str:  "2026-02",
			t:    iso8601.Time{Format: iso8601.DateMonth, Time: time.Date(2026, 2, 1, 0, 0, 0, 0, time.UTC)},
		},
		{
			json: `"2026-09-02"`,
			str:  "2026-09-02",
			t:    iso8601.Time{Format: iso8601.DateDay, Time: time.Date(2026, 9, 2, 0, 0, 0, 0, time.UTC)},
		},
		{
			json: `"2026-06-04T09:32:44Z"`,
			str:  "2026-06-04T09:32:44Z",
			t:    iso8601.Time{Format: iso8601.DateTime, Time: time.Date(2026, 6, 4, 9, 32, 44, 0, time.UTC)},
		},
		{
			json: `"2026-06-04T09:32:44+0000"`,
			str:  "2026-06-04T09:32:44Z",
			t:    iso8601.Time{Format: iso8601.DateTime, Time: time.Date(2026, 6, 4, 9, 32, 44, 0, time.UTC)},
		},
		{
			json: `"2026-06-04T09:32:44+0100"`,
			str:  "2026-06-04T08:32:44Z",
			t:    iso8601.Time{Format: iso8601.DateTime, Time: time.Date(2026, 6, 4, 8, 32, 44, 0, time.UTC)},
		},
		{
			json: `"2026-06-04T09:32:44.642Z"`,
			str:  "2026-06-04T09:32:44.642Z",
			t:    iso8601.Time{Format: iso8601.DateTime, Time: time.Date(2026, 6, 4, 9, 32, 44, int(642*time.Millisecond), time.UTC)},
		},
	}
)

func TestParseString(t *testing.T) {
	for _, tc := range tcs {
		ts, err := iso8601.ParseString(tc.str)
		if tc.err {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}
		assert.Truef(t, tc.t.Equal(ts), "source: %s, parsed: %s", tc.str, ts.String())
	}
}

func TestString(t *testing.T) {
	for _, tc := range tcs {
		str := tc.t.String()
		assert.Equal(t, tc.str, str)
	}
}

func TestUnmarshal(t *testing.T) {
	for _, tc := range tcs {
		var ts iso8601.Time
		err := json.Unmarshal([]byte(tc.json), &ts)
		require.NoError(t, err)
		assert.Truef(t, tc.t.Equal(ts), "source: %s, unmarshaled: %s", tc.json, ts.String())
	}
}

func TestMarshal(t *testing.T) {
	for _, tc := range tcs {
		buf, err := json.Marshal(tc.t)
		require.NoError(t, err)
		var ts iso8601.Time
		err = json.Unmarshal(buf, &ts)
		require.NoError(t, err)
		assert.Truef(t, tc.t.Equal(ts), "source: %s, json %s, unmarshaled: %s", tc.t.String(), string(buf), ts.String())
	}
}

func TestAt(t *testing.T) {
	d := iso8601.Day(2026, 2, 15)
	assert.Equal(t, "2026-02-15", d.String())
	at := d.At(23, 42, 15, int(999*time.Millisecond))
	assert.Equal(t, "2026-02-15T23:42:15.999Z", at.String())
}

func TestStart(t *testing.T) {
	t.Run("MonthStart/MonthEnd work with standard years", func(t *testing.T) {
		d := iso8601.Day(2026, 2, 15)
		assert.Equal(t, "2026-02-15", d.String())
		start := d.MonthStart()
		assert.Equal(t, "2026-02-01", start.String())
		end := d.MonthEnd()
		assert.Equal(t, "2026-02-28", end.String())
	})
	t.Run("MonthStart/MonthEnd work with leap years", func(t *testing.T) {
		d := iso8601.Day(2028, 2, 15)
		assert.Equal(t, "2028-02-15", d.String())
		start := d.MonthStart()
		assert.Equal(t, "2028-02-01", start.String())
		end := d.MonthEnd()
		assert.Equal(t, "2028-02-29", end.String())
	})
	t.Run("AtStart/AtEnd work with standard years", func(t *testing.T) {
		d := iso8601.Day(2026, 2, 15)
		assert.Equal(t, "2026-02-15", d.String())
		start := d.AtStart()
		assert.Equal(t, "2026-02-01T00:00:00Z", start.String())
		end := d.AtEnd()
		assert.Equal(t, "2026-02-28T23:59:59.999Z", end.String())
	})
	t.Run("AtStart/AtEnd work with leap years", func(t *testing.T) {
		d := iso8601.Day(2028, 2, 15)
		assert.Equal(t, "2028-02-15", d.String())
		start := d.AtStart()
		assert.Equal(t, "2028-02-01T00:00:00Z", start.String())
		end := d.AtEnd()
		assert.Equal(t, "2028-02-29T23:59:59.999Z", end.String())
	})
}
