package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	t, err := time.Parse("1/2/2006 15:04:05", date)

	if err != nil {
		panic(err)
	}

	return t
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	t, err := time.Parse("January 2, 2006 15:05:04", date)

	if err != nil {
		panic(err)
	}

	return t.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	t, err := time.Parse("Monday, January 2, 2006 15:05:04", date)

	if err != nil {
		panic(err)
	}

	hour := t.Hour()

	return 12 <= hour && hour < 18
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	schedule := Schedule(date)

	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", schedule.Weekday(), schedule.Month(), schedule.Day(), schedule.Year(), schedule.Hour(), schedule.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	now := time.Now()
	layout := "2006-1-2"
	value := fmt.Sprintf("%d-%d-%d", now.Year(), 9, 15)

	t, err := time.Parse(layout, value)

	if err != nil {
		panic(err)
	}

	return t
}
