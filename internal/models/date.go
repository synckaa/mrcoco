package models

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

type Date struct {
	time.Time
}

func (d *Date) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), "\"")
	if s == "null" || s == "" {
		return nil
	}

	_LAYOUTS := []string{
		"2006-01-02",
		"2006-01-02T15:04:05Z07:00",
		"2006-01-02T15:04:05",
	}

	for _, layout := range _LAYOUTS {
		t, err := time.Parse(layout, s)
		if err == nil {
			d.Time = t
			return nil
		}
	}

	return nil
}

func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.Format("2006-01-02") + `"`), nil
}

func (d Date) Value() (driver.Value, error) {
	if d.Time.IsZero() {
		return nil, nil
	}
	return d.Time.Format("2006-01-02"), nil
}

func (d *Date) Scan(value interface{}) error {
	if value == nil {
		d.Time = time.Time{}
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		d.Time = v
	case []byte:
		t, err := time.Parse("2006-01-02", string(v))
		if err != nil {
			t, err = time.Parse("2006-01-02T15:04:05Z07:00", string(v))
		}
		if err != nil {
			return fmt.Errorf("cannot parse date: %v", err)
		}
		d.Time = t
	case string:
		t, err := time.Parse("2006-01-02", v)
		if err != nil {
			t, err = time.Parse("2006-01-02T15:04:05Z07:00", v)
		}
		if err != nil {
			return fmt.Errorf("cannot parse date: %v", err)
		}
		d.Time = t
	}
	return nil
}
