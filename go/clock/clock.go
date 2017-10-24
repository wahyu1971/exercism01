package clock

import "fmt"

const testVersion = 4

// You can find more details and hints in the test file.

// describe clock data structure
type Clock struct {
	hour   int
	minute int
}

func New(hour, minute int) (result Clock) {
	result.hour = hour
	result.minute = minute
	result.Add(0) // fix hour and minutes
	return
}

func (c Clock) String() string {
	hs := fmt.Sprintf("%02d", c.hour)
	ms := fmt.Sprintf("%02d", c.minute)
	return hs + ":" + ms
}

func (c Clock) Add(minutes int) (result Clock) {
	tmp := c.minute + minutes
	result = c
	for {
		if tmp >= 60 {
			result.hour++
			tmp = tmp - 60
			continue
		}
		result.minute = tmp
		break
	}
	for {
		if result.hour >= 24 {
			result.hour = result.hour - 24
			continue
		}
		break
	}
	return
}

// Remember to delete all of the stub comments.
// They are just noise, and reviewers will complain.
