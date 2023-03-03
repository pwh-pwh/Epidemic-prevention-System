package models

import (
	"database/sql/driver"
	"errors"
	"time"
)

const (
	timeFormat = "2006-01-02 15:04:05"
	timezone   = "Asia/Shanghai"
)

type LocalTime time.Time

func (t LocalTime) MarshalJSON() ([]byte, error) {
	if time.Time(t).IsZero() {
		return []byte("null"), nil
	}
	b := make([]byte, 0, len(timeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormat)
	b = append(b, '"')
	return b, nil
}

func (t *LocalTime) UnmarshalJSON(data []byte) (err error) {
	if string(data) == "null" {
		*t = LocalTime{}
		return nil
	}
	now, err := time.ParseInLocation(`"`+timeFormat+`"`, string(data), time.Local)
	*t = LocalTime(now)
	return
}

func (t LocalTime) String() string {
	return time.Time(t).Format(timeFormat)
}

func (t LocalTime) Local() time.Time {
	location, _ := time.LoadLocation(timezone)
	return time.Time(t).In(location)
}

func (t LocalTime) Value() (driver.Value, error) {
	var ti = time.Time(t)
	if ti.IsZero() {
		return nil, nil
	}
	return ti, nil
}

func (t *LocalTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = LocalTime(value)
		return nil
	}
	return errors.New("can not convert to LocalTime")
}
