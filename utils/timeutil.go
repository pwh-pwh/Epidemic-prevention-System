package utils

import (
	"time"
)

// FormatTime formats the given time using the given layout
func FormatTime(t time.Time, layout string) string {
	return t.Format(layout)
}

// ParseTime parses the given string using the given layout
func ParseTime(str string, layout string) (time.Time, error) {
	return time.Parse(layout, str)
}

// UnixTimestamp returns the Unix timestamp for the given time
func UnixTimestamp(t time.Time) int64 {
	return t.Unix()
}

// TimeFromUnixTimestamp returns the time corresponding to the given Unix timestamp
func TimeFromUnixTimestamp(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}

// DurationFromNow returns the duration between the current time and the given time
func DurationFromNow(t time.Time) time.Duration {
	return time.Until(t)
}

// IsAfterNow returns true if the given time is after the current time
func IsAfterNow(t time.Time) bool {
	return t.After(time.Now())
}

// IsBeforeNow returns true if the given time is before the current time
func IsBeforeNow(t time.Time) bool {
	return t.Before(time.Now())
}
