package timerange_test

import (
	"testing"
	"time"

	"github.com/TristanSch1/flow/pkg/timerange"
)

func TestTimeRange_NewDayTimeRange(t *testing.T) {
	day := time.Date(2024, 4, 17, 19, 0, 0, 0, time.UTC)
	expected := timerange.TimeRange{
		Since: time.Date(2024, 4, 17, 0, 0, 0, 0, time.UTC),
		Until: time.Date(2024, 4, 18, 0, 0, 0, 0, time.UTC).Add(-time.Second),
	}
	got := timerange.NewDayTimeRange(day)
	if got != expected {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}

func TestTimeRange_NewWeekTimeRange(t *testing.T) {
	day := time.Date(2024, 4, 17, 19, 0, 0, 0, time.UTC)
	expected := timerange.TimeRange{
		Since: time.Date(2024, 4, 15, 0, 0, 0, 0, time.UTC),
		Until: time.Date(2024, 4, 22, 0, 0, 0, 0, time.UTC).Add(-time.Second),
	}
	got := timerange.NewWeekTimeRange(day)
	if got != expected {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}

func TestTimeRange_NewMonthTimeRange(t *testing.T) {
	day := time.Date(2024, 4, 17, 19, 0, 0, 0, time.UTC)
	expected := timerange.TimeRange{
		Since: time.Date(2024, 4, 1, 0, 0, 0, 0, time.UTC),
		Until: time.Date(2024, 5, 1, 0, 0, 0, 0, time.UTC).Add(-time.Second),
	}
	got := timerange.NewMonthTimeRange(day)
	if got != expected {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}
