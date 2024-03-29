package service

import "time"

const (
	saturday = 6
	sunday   = 0
)

func isWeekend(now time.Time) bool {
	switch now.Weekday() {
	case saturday, sunday:
		return true
	default:
		return false
	}
}
