package wrappers

import "time"

type CustomTime struct {
	time.Time
}

func (ct *CustomTime) UnmarshalParam(src string) error {
	t, err := time.Parse(time.RFC3339, src)

	if err != nil {
		return err
	}

	ct.Time = t

	return nil
}
