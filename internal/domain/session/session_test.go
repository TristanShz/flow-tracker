package session

import (
	"testing"
	"time"
)

func TestSession_Duration(t *testing.T) {
	tests := []struct {
		name string
		e    Session
		want time.Duration
	}{
		{
			name: "test",
			e: Session{
				StartTime: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				EndTime:   time.Date(2020, 1, 1, 0, 0, 1, 0, time.UTC),
			},
			want: time.Second,
		},
		{
			name: "without end time",
			e: Session{
				StartTime: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Duration(); got != tt.want {
				t.Errorf("Entry.Duration() = %v, want %v", got, tt.want)
			}
		})
	}
}
