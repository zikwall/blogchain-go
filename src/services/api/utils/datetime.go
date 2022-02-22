package utils

import "time"

const DateTimeFormat = "2006-01-02 15:04:05"
const DateFormat = "2006-01-02"

func Date(t time.Time) string {
	return format(t, DateFormat)
}

func Datetime(t time.Time) string {
	return format(t, DateTimeFormat)
}

func format(t time.Time, format string) string {
	location, _ := time.LoadLocation("UTC")
	return t.In(location).Format(format)
}
