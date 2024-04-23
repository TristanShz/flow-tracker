package timerange

import (
	"time"
)

type TimeRange struct {
	Since time.Time
	Until time.Time
}

func (t TimeRange) IsZero() bool {
	return t.Since.IsZero() && t.Until.IsZero()
}

func (t TimeRange) JustSince() bool {
	return !t.Since.IsZero() && t.Until.IsZero()
}

func (t TimeRange) JustUntil() bool {
	return t.Since.IsZero() && !t.Until.IsZero()
}

func (t TimeRange) SinceAndUntil() bool {
	return !t.Since.IsZero() && !t.Until.IsZero()
}

func NewDayTimeRange(day time.Time) TimeRange {
	startOfDay := time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, time.UTC)
	endOfDay := startOfDay.AddDate(0, 0, 1).Add(-time.Second)
	return TimeRange{
		Since: startOfDay,
		Until: endOfDay,
	}
}

func NewWeekTimeRange(day time.Time) TimeRange {
	weekDay := int(day.Weekday())
	weekStart := day.AddDate(0, 0, -(weekDay - 1))
	weekEnd := weekStart.AddDate(0, 0, 7).Add(-time.Second)
	return TimeRange{
		Since: time.Date(weekStart.Year(), weekStart.Month(), weekStart.Day(), 0, 0, 0, 0, time.UTC),
		Until: time.Date(weekEnd.Year(), weekEnd.Month(), weekEnd.Day(), 0, 0, 0, 0, time.UTC).Add(-time.Second),
	}
}

func NewMonthTimeRange(day time.Time) TimeRange {
	monthStart := time.Date(day.Year(), day.Month(), 1, 0, 0, 0, 0, time.UTC)
	monthEnd := monthStart.AddDate(0, 1, 0).Add(-time.Second)
	return TimeRange{
		Since: monthStart,
		Until: monthEnd,
	}
}
