package booking

import "time"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	var t, _ = time.Parse("1/2/2006 15:04:05", date)

	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	var t, _ = time.Parse("January 2, 2006 15:04:05", date)
	return t.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	var t, _ = time.Parse("Monday, January 2, 2006 15:04:05", date)

	return t.Hour() >= 12 && t.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	var t, _ = time.Parse("1/2/2006 15:04:05", date)
	return t.Format("You have an appointment on Monday, January 2, 2006, at 15:04.")
	return Schedule(date).Format("You have an appointment on Monday")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	var t, _ = time.Parse("2006-01-02", "2020-09-15")
	return time.Date(time.Now().Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
}
