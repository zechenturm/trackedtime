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

func TestComment(t *testing.T) {
	i := interval{}
	i.Comment = "test"
	if i.Comment != "test" {
		t.Fatalf("Comment failed, got \"%v\"", i.Comment)
	}
}

func TestChecked(t *testing.T) {
	i := interval{}
	if i.Checked != false {
		t.Fatalf("checked is not false")
	}
	i.Checked = true
	if i.Checked != true {
		t.Fatalf("checked is not true")
	}
}

func TestCalendarWeek(t *testing.T) {
	i := interval{}
	type cwtime struct {
		Time string
		Week int
		Year int
	}
	times := []cwtime{
		{"16:10 30 Sep 2019", 40, 2019},
		{"8:20 23 Nov 2018", 47, 2018},
		{"11:40 2 Oct 2018", 40, 2018},
		{"12:00 2 Oct 2019", 40, 2019},
		{"16:30 1 Oct 2019", 40, 2019},
	}
	for _, tm := range times {
		cwt, err := time.Parse("15:04 2 Jan 2006", tm.Time)
		if err != nil {
			t.Fatalf("failed to parse time: %v", err)
		}
		i.StartTime = cwt
		cw := i.CalWeek()
		if cw.Week != tm.Week || cw.Year != tm.Year {
			t.Fatalf("wrong calendar week, got (%v, %v), wanted (%v, %v)", tm.Week, tm.Year, cw.Week, cw.Year)
		}
	}
}

func TestAccHours(t *testing.T) {
	type timeDur struct {
		Start    string
		Duration time.Duration
	}

	times := []timeDur{
		{"16:10 30 Sep 2019", 4 * time.Hour},
		{"8:20 23 Nov 2018", 8 * time.Hour},
		{"11:40 2 Oct 2018", 7*time.Hour + 30*time.Minute},
		{"12:00 2 Oct 2019", 6 * time.Hour},
		{"16:30 1 Oct 2019", 3 * time.Hour},
	}

	refHours := 0.0
	intervals := make([]interval, len(times))
	for _, tm := range times {
		tt, err := time.Parse("15:04 2 Jan 2006", tm.Start)
		if err != nil {
			t.Fatalf("failed to parse time: %v", err)
		}
		intervals = append(intervals, interval{
			StartTime: tt,
			StopTime:  tt.Add(tm.Duration),
		})
		refHours += tm.Duration.Hours()
	}

	hours := AccumulateHours(&intervals)
	floatEqual := func(a float64, b float64, tolerance float64) bool {
		return (a > b-tolerance) && (a < b+tolerance)
	}
	if !(floatEqual(hours, refHours, 0.0001)) {
		t.Fatalf("wrong time accumulated: got %v, wanted %v", hours, refHours)
	}
}
