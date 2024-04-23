package timerange_test

import (
	"testing"
	"time"

	"github.com/TristanSch1/flow/pkg/timerange"
)

func TestTimeRange(t *testing.T) {
	tests := []struct {
		tr                timerange.TimeRange
		name              string
		wantIsZero        bool
		wantJustSince     bool
		wantJustUntil     bool
		wantSinceAndUntil bool
	}{
		{
			name: "Since is zero",
			tr: timerange.TimeRange{
				Since: time.Time{},
				Until: time.Date(2024, 4, 17, 20, 0, 0, 0, time.UTC),
			},
			wantIsZero:        false,
			wantJustSince:     false,
			wantJustUntil:     true,
			wantSinceAndUntil: false,
		},
		{
			name: "Until is zero",
			tr: timerange.TimeRange{
				Since: time.Date(2024, 4, 17, 20, 0, 0, 0, time.UTC),
				Until: time.Time{},
			},
			wantIsZero:        false,
			wantJustSince:     true,
			wantJustUntil:     false,
			wantSinceAndUntil: false,
		},
		{
			name: "Both are zero",
			tr: timerange.TimeRange{
				Since: time.Time{},
				Until: time.Time{},
			},
			wantIsZero:        true,
			wantJustSince:     false,
			wantJustUntil:     false,
			wantSinceAndUntil: false,
		},
		{
			name: "None are zero",
			tr: timerange.TimeRange{
				Since: time.Date(2024, 4, 17, 20, 0, 0, 0, time.UTC),
				Until: time.Date(2024, 4, 17, 20, 0, 0, 0, time.UTC),
			},
			wantIsZero:        false,
			wantJustSince:     false,
			wantJustUntil:     false,
			wantSinceAndUntil: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.IsZero(); got != tt.wantIsZero {
				t.Errorf("TimeRange.IsZero() = %v, want %v", got, tt.wantIsZero)
			}

			if got := tt.tr.JustSince(); got != tt.wantJustSince {
				t.Errorf("TimeRange.JustSince() = %v, want %v", got, tt.wantJustSince)
			}

			if got := tt.tr.JustUntil(); got != tt.wantJustUntil {
				t.Errorf("TimeRange.JustUntil() = %v, want %v", got, tt.wantJustUntil)
			}

			if got := tt.tr.SinceAndUntil(); got != tt.wantSinceAndUntil {
				t.Errorf("TimeRange.SinceAndUntil() = %v, want %v", got, tt.wantSinceAndUntil)
			}
		})
	}
}

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
