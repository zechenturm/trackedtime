package trackedtime

import "time"

type interval struct {
	StartTime time.Time
	StopTime  time.Time
	Comment   string
	Checked   bool
}

type CalWeek struct {
	Week int
	Year int
}

func (i interval) Duration() time.Duration {
	return i.StopTime.Sub(i.StartTime)
}

func (i interval) CalWeek() CalWeek {
	y, w := i.StartTime.ISOWeek()
	return CalWeek{w, y}
}

func AccumulateHours(intervals *[]interval) float64 {
	hours := 0.0
	for _, i := range *intervals {
		hours += i.Duration().Hours()
	}
	return hours
}
