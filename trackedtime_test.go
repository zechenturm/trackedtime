package trackedtime

import (
	"log"
	"testing"
	"time"
)

func TestCreate(t *testing.T) {
	_ = interval{}
}

func TestSetStartTime(t *testing.T) {
	i := interval{}
	tm, _ := time.Parse("15:04 2 Jan 2006", "16:10 11 Nov 2018")
	i.StartTime = tm
	if i.StartTime != tm {
		t.Fatalf("start time assignment failed: wanted %v, got %v", tm, i.StartTime)
	}
}

func TestStopTime(t *testing.T) {
	i := interval{}
	tm, _ := time.Parse("15:04 2 Jan 2006", "16:10 11 Nov 2018")
	i.StopTime = tm
	if i.StopTime != tm {
		t.Fatalf("start time assignment failed: wanted %v, got %v", tm, i.StopTime)
	}
}
func TestIntervalZero(t *testing.T) {
	i := interval{}
	tm, _ := time.Parse("15:04 2 Jan 2006", "16:10 11 Nov 2018")
	i.StartTime = tm
	i.StopTime = tm
	if i.Duration() != time.Duration(0) {
		log.Fatalf("Duration not 0, got %v", i.Duration())
	}
}

func TestIntervalNonZero(t *testing.T) {
	i := interval{}
	start, _ := time.Parse("15:04 2 Jan 2006", "16:10 11 Nov 2018")
	stop := start.Add(time.Hour)
	i.StartTime = start
	i.StopTime = stop
	if i.Duration() != time.Hour {
		log.Fatalf("Wrong Duration: want %v, got %v", stop.Sub(start), i.Duration())
	}
}
