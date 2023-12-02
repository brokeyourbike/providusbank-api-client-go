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
	s := strings.Trim(string(b), `"`)
	if s == "" {
		return nil
	}

	formats := []string{
		"1/02/2006 3:04:05 PM",
		"2006-01-02 15:04:05",
		"2006-01-02T15:04:05.999-0700",
	}

	for _, f := range formats {
		parsed, err := time.Parse(f, s)
		if err == nil {
			t.Time = parsed
			return nil
		}
	}

	return fmt.Errorf("cannot parse time string %s", s)
}
