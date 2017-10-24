package clock

import "fmt"

const testVersion = 4

// describe clock data structure
type Clock struct {
	hour   int
	minute int
}

func New(hour, minute int) (result Clock) {
	result.hour = hour
	result.minute = minute
	result = result.fix_minutes()
	result = result.fix_hours()
	return
}

func (c Clock) String() string {
	hs := fmt.Sprintf("%02d", c.hour)
	ms := fmt.Sprintf("%02d", c.minute)
	return hs + ":" + ms
}

func (c Clock) Add(minutes int) (result Clock) {
	c.minute = c.minute + minutes
	result = c
	result = result.fix_minutes()
	result = result.fix_hours()
	return
}

// not very optimize, but works
func (c Clock) fix_minutes() (result Clock) {
	for {
		switch {
		// 60 atau lebih
		case c.minute > 59:
			c.hour++
			c.minute = c.minute - 60
		// kurang dari 0
		case c.minute < 0:
			c.hour--
			c.minute = c.minute + 60
		default:
			result = c
			return
		}
	}
}

func (c Clock) fix_hours() (result Clock) {
	for {
		switch {
		// 60 atau lebih
		case c.hour > 23:
			c.hour = c.hour - 24
		// kurang dari 0
		case c.hour < 0:
			c.hour = c.hour + 24
		default:
			result = c
			return
		}
	}
}
