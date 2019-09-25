package trackedtime

import "time"

type interval struct {
	StartTime time.Time
	StopTime  time.Time
	Comment   string
}

func (i interval) Duration() time.Duration {
	return i.StopTime.Sub(i.StartTime)
}
