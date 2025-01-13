package models

import "time"

type TimeOnly time.Time

// unmarshalls time in the format "HH:MM"
func (t *TimeOnly) UnmarshalJSON(b []byte) error {
	tm, err := time.Parse(`"15:04"`, string(b))
	if err != nil {
		return err
	}
	*t = TimeOnly(tm)
	return nil
}

// marshalls time in the format "HH:MM"
func (t TimeOnly) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(t).Format(`"15:04"`)), nil
}

// returns the time.Time value
func (t TimeOnly) Time() time.Time {
	return time.Time(t)
}
