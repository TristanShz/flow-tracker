package entry

import (
	"testing"
	"time"
)

func TestEntry_Duration(t *testing.T) {
	tests := []struct {
		name string
		e    Entry
		want time.Duration
	}{
		{
			name: "test",
			e: Entry{
				StartTime: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				EndTime:   time.Date(2020, 1, 1, 0, 0, 1, 0, time.UTC),
			},
			want: time.Second,
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
