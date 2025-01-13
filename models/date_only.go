package models

import (
	"time"
)

type DateOnly time.Time

// unmarshals date in the format "YYYY-MM-DD"
func (d *DateOnly) UnmarshalJSON(b []byte) error {
	t, err := time.Parse(`"2006-01-02"`, string(b))
	if err != nil {
		return err
	}
	*d = DateOnly(t)
	return nil
}

// marshals date in the format "YYYY-MM-DD"
func (d DateOnly) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(d).Format(`"2006-01-02"`)), nil
}

// returns the time.Time value of the date
func (d DateOnly) Time() time.Time {
	return time.Time(d)
}
