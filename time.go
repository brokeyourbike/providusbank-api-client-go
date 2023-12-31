package providusbank

import (
	"fmt"
	"strings"
	"time"
)

type Time struct {
	time.Time
}

func (t *Time) UnmarshalJSON(b []byte) error {
	loc, err := time.LoadLocation("Africa/Lagos")
	if err != nil {
		return fmt.Errorf("cannot find timezeone: %w", err)
	}

	s := strings.Trim(string(b), `"`)
	if s == "" {
		return nil
	}

	formats := []string{
		"1/02/2006 3:04:05 PM",
		"2006-01-02 15:04:05",
		"2006-01-02T15:04:05.999-0700",
		"02-Jan-06",
	}

	for _, f := range formats {
		parsed, err := time.ParseInLocation(f, s, loc)
		if err == nil {
			t.Time = parsed
			return nil
		}
	}

	return fmt.Errorf("cannot parse time string %s", s)
}
