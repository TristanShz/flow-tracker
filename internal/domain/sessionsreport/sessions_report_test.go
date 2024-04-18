package sessionsreport_test

import (
	"testing"
	"time"

	"github.com/TristanSch1/flow/internal/domain/session"
	"github.com/TristanSch1/flow/internal/domain/sessionsreport"
)

func TestSessionsReport_TotalDuration(t *testing.T) {
	tests := []struct {
		name string
		e    sessionsreport.SessionsReport
		want time.Duration
	}{
		{
			name: "test",
			e: sessionsreport.SessionsReport{
				Sessions: []session.Session{
					{
						StartTime: time.Date(2020, 1, 1, 8, 0, 0, 0, time.UTC),
						EndTime:   time.Date(2020, 1, 1, 10, 0, 0, 0, time.UTC),
						Project:   "my-todo",
						Tags:      []string{"add-todo"},
					},
					{
						StartTime: time.Date(2020, 1, 1, 12, 0, 0, 0, time.UTC),
						EndTime:   time.Date(2020, 1, 1, 12, 50, 0, 0, time.UTC),
						Project:   "my-todo",
						Tags:      []string{"add-todo"},
					},
				},
			},
			want: 2*time.Hour + 50*time.Minute,
		},
		{
			name: "One session without EndTime",
			e: sessionsreport.SessionsReport{
				Sessions: []session.Session{
					{
						StartTime: time.Date(2020, 1, 1, 8, 0, 0, 0, time.UTC),
						Project:   "my-todo",
						Tags:      []string{"add-todo"},
					},
				},
			},
			want: 0,
		},
		{
			name: "No sessions",
			e: sessionsreport.SessionsReport{
				Sessions: []session.Session{},
			},
			want: 0,
		},
		{
			name: "One with EndTime and one without",
			e: sessionsreport.SessionsReport{
				Sessions: []session.Session{
					{
						StartTime: time.Date(2020, 1, 1, 8, 0, 0, 0, time.UTC),
						EndTime:   time.Date(2020, 1, 1, 10, 0, 0, 0, time.UTC),
						Project:   "my-todo",
					},
					{
						StartTime: time.Date(2020, 1, 1, 12, 0, 0, 0, time.UTC),
						Project:   "my-todo",
					},
				},
			},
			want: 2 * time.Hour,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.TotalDuration(); got != tt.want {
				t.Errorf("Entry.Duration() = %v, want %v", got, tt.want)
			}
		})
	}
}
