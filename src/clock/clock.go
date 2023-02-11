package clock

import "fmt"

var minutesInDay = 24 * 60

type Clock struct {
	Hours   int
	Minutes int
}

func New(h, m int) Clock {
	total := (h*60 + m) % minutesInDay
	if total < 0 {
		total += minutesInDay
	}

	return Clock{
		Hours:   total / 60,
		Minutes: total % 60,
	}
}

func (c Clock) Add(m int) Clock {
	return New(c.Hours, c.Minutes+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(c.Hours, c.Minutes-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hours, c.Minutes)
}
