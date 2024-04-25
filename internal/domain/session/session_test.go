package session_test

import (
	"testing"
	"time"

	"github.com/TristanSch1/flow/internal/domain/session"
)

func TestSession_Duration(t *testing.T) {
	tt := []struct {
		name string
		e    session.Session
		want time.Duration
	}{
		{
			name: "test",
			e: session.Session{
				StartTime: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				EndTime:   time.Date(2020, 1, 1, 0, 0, 1, 0, time.UTC),
			},
			want: time.Second,
		},
		{
			name: "without end time",
			e: session.Session{
				StartTime: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: 0,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.e.Duration(); got != tc.want {
				t.Errorf("Entry.Duration() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestSession_Status(t *testing.T) {
	tt := []struct {
		name string
		want string
		e    session.Session
	}{
		{
			name: "Session with end time",
			e: session.Session{
				Id:        "1",
				StartTime: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				EndTime:   time.Date(2020, 1, 1, 0, 0, 1, 0, time.UTC),
			},
			want: session.EndedStatus,
		},
		{
			name: "Session without end time",
			e: session.Session{
				Id:        "1",
				StartTime: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: session.FlowingStatus,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.e.Status(); got != tc.want {
				t.Errorf("Entry.Status() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestSession_GetFormattedEndTime(t *testing.T) {
	tt := []struct {
		name string
		want string
		e    session.Session
	}{
		{
			name: "Session with end time",
			e: session.Session{
				Id:        "1",
				StartTime: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				EndTime:   time.Date(2020, 1, 1, 0, 0, 1, 0, time.UTC),
			},
			want: "2020-01-01 00:00:01",
		},
		{
			name: "Session without end time",
			e: session.Session{
				Id:        "1",
				StartTime: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: "/",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.e.GetFormattedEndTime(); got != tc.want {
				t.Errorf("Entry.GetFormattedEndTime() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestSession_GetFormattedStartTime(t *testing.T) {
	session := session.Session{
		StartTime: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	want := "2020-01-01 00:00:00"

	if got := session.GetFormattedStartTime(); got != want {
		t.Errorf("Entry.GetFormattedEndTime() = %v, want %v", got, want)
	}
}

func TestSession_HasTag(t *testing.T) {
	tt := []struct {
		name string
		e    session.Session
		want bool
	}{
		{
			name: "Session with one tag",
			e: session.Session{
				Id:        "1",
				StartTime: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				Project:   "my-todo",
				Tags:      []string{"tag"},
			},
			want: true,
		},
		{
			name: "Session with multiple tags",
			e: session.Session{
				Id:        "1",
				StartTime: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				Project:   "my-todo",
				Tags:      []string{"add-todo", "remove-todo", "tag"},
			},
			want: true,
		},
		{
			name: "Session without tag",
			e: session.Session{
				Id:        "1",
				StartTime: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				Project:   "my-todo",
			},
			want: false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.e.HasTag("tag"); got != tc.want {
				t.Errorf("Entry.HasTag() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestSession_Equals(t *testing.T) {
	tt := []struct {
		name  string
		e     session.Session
		given session.Session
		want  bool
	}{
		{
			name: "Same id",
			e: session.Session{
				Id:        "1",
				StartTime: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				Project:   "my-todo",
			},
			given: session.Session{
				Id:        "1",
				StartTime: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				Project:   "my-todo",
			},
			want: true,
		},
		{
			name: "Different id",
			e: session.Session{
				Id:        "1",
				StartTime: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				Project:   "my-todo",
			},
			given: session.Session{
				Id:        "2",
				StartTime: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				Project:   "my-todo",
			},
			want: false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.e.Equals(tc.given); got != tc.want {
				t.Errorf("Entry.Equals() = %v, want %v", got, tc.want)
			}
		})
	}
}
