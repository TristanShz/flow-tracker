package sessionsreport_test

import (
	"reflect"
	"slices"
	"testing"
	"time"

	"github.com/TristanSch1/flow/internal/domain/session"
	"github.com/TristanSch1/flow/internal/domain/sessionsreport"
)

var sessionsReportTest = sessionsreport.NewSessionsReport([]session.Session{
	{
		Id:        "1",
		StartTime: time.Date(2020, 1, 1, 8, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2020, 1, 1, 10, 0, 0, 0, time.UTC),
		Project:   "my-todo",
		Tags:      []string{"add-todo"},
	},
	{
		Id:        "2",
		StartTime: time.Date(2020, 1, 1, 12, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2020, 1, 1, 13, 0, 0, 0, time.UTC),
		Project:   "my-todo",
		Tags:      []string{"add-todo"},
	},
	{
		Id:        "3",
		StartTime: time.Date(2020, 1, 2, 7, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2020, 1, 2, 9, 0, 0, 0, time.UTC),
		Project:   "my-todo",
		Tags:      []string{"add-todo"},
	},
	{
		Id:        "4",
		StartTime: time.Date(2020, 1, 2, 12, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2020, 1, 2, 18, 0, 0, 0, time.UTC),
		Project:   "my-todo",
		Tags:      []string{"add-todo", "remove-todo"},
	},
	{
		Id:        "5",
		StartTime: time.Date(2020, 1, 2, 20, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2020, 1, 2, 21, 0, 0, 0, time.UTC),
		Project:   "flow",
		Tags:      []string{"start-usecase"},
	},
	{
		Id:        "6",
		StartTime: time.Date(2020, 1, 3, 8, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2020, 1, 3, 10, 0, 0, 0, time.UTC),
		Project:   "flow",
		Tags:      []string{"start-usecase"},
	},
	{
		Id:        "7",
		StartTime: time.Date(2020, 1, 3, 12, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2020, 1, 3, 13, 0, 0, 0, time.UTC),
		Project:   "flow",
		Tags:      []string{"start-usecase", "stop-usecase"},
	},
})

func TestSessionsReport_TotalDuration(t *testing.T) {
	tests := []struct {
		name string
		e    sessionsreport.SessionsReport
		want time.Duration
	}{
		{
			name: "test",
			e: sessionsreport.NewSessionsReport([]session.Session{
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
			}),
			want: 2*time.Hour + 50*time.Minute,
		},
		{
			name: "One session without EndTime",
			e: sessionsreport.NewSessionsReport([]session.Session{
				{
					StartTime: time.Date(2020, 1, 1, 8, 0, 0, 0, time.UTC),
					Project:   "my-todo",
					Tags:      []string{"add-todo"},
				},
			}),
			want: 0,
		},
		{
			name: "No sessions",
			e:    sessionsreport.NewSessionsReport([]session.Session{}),
			want: 0,
		},
		{
			name: "One with EndTime and one without",
			e: sessionsreport.NewSessionsReport([]session.Session{
				{
					StartTime: time.Date(2020, 1, 1, 8, 0, 0, 0, time.UTC),
					EndTime:   time.Date(2020, 1, 1, 10, 0, 0, 0, time.UTC),
					Project:   "my-todo",
				},
				{
					StartTime: time.Date(2020, 1, 1, 12, 0, 0, 0, time.UTC),
					Project:   "my-todo",
				},
			}),
			want: 2 * time.Hour,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.TotalDuration(); got != tt.want {
				t.Errorf("Entry.TotalDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSessionsReport_SplitBy(t *testing.T) {
	tests := []struct {
		wantByDays     map[time.Time][]session.Session
		wantByProjects map[string][]session.Session
		name           string
		e              sessionsreport.SessionsReport
	}{
		{
			name: "Two sessions on the same day",
			e: sessionsreport.NewSessionsReport([]session.Session{
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
			}),
			wantByDays: map[time.Time][]session.Session{
				time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC): {
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
			wantByProjects: map[string][]session.Session{
				"my-todo": {
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
		},
		{
			name: "One session without EndTime",
			e: sessionsreport.NewSessionsReport([]session.Session{
				{
					StartTime: time.Date(2020, 1, 1, 8, 0, 0, 0, time.UTC),
					Project:   "my-todo",
					Tags:      []string{"add-todo"},
				},
			}),
			wantByDays: map[time.Time][]session.Session{
				time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC): {
					{
						StartTime: time.Date(2020, 1, 1, 8, 0, 0, 0, time.UTC),
						Project:   "my-todo",
						Tags:      []string{"add-todo"},
					},
				},
			},
			wantByProjects: map[string][]session.Session{
				"my-todo": {
					{
						StartTime: time.Date(2020, 1, 1, 8, 0, 0, 0, time.UTC),
						Project:   "my-todo",
						Tags:      []string{"add-todo"},
					},
				},
			},
		},
		{
			name:           "No sessions",
			e:              sessionsreport.NewSessionsReport([]session.Session{}),
			wantByDays:     map[time.Time][]session.Session{},
			wantByProjects: map[string][]session.Session{},
		},
		{
			name: "Two sessions on different days for different projects",
			e: sessionsreport.NewSessionsReport([]session.Session{
				{
					Id:        "1",
					StartTime: time.Date(2020, 1, 1, 8, 0, 0, 0, time.UTC),
					EndTime:   time.Date(2020, 1, 1, 10, 0, 0, 0, time.UTC),
					Project:   "my-todo",
				},
				{
					Id:        "2",
					StartTime: time.Date(2020, 1, 2, 12, 0, 0, 0, time.UTC),
					EndTime:   time.Date(2020, 1, 2, 10, 0, 0, 0, time.UTC),
					Project:   "flow",
				},
			}),
			wantByDays: map[time.Time][]session.Session{
				time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC): {
					{
						Id:        "1",
						StartTime: time.Date(2020, 1, 1, 8, 0, 0, 0, time.UTC),
						EndTime:   time.Date(2020, 1, 1, 10, 0, 0, 0, time.UTC),
						Project:   "my-todo",
					},
				},
				time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC): {
					{
						Id:        "2",
						StartTime: time.Date(2020, 1, 2, 12, 0, 0, 0, time.UTC),
						EndTime:   time.Date(2020, 1, 2, 10, 0, 0, 0, time.UTC),
						Project:   "flow",
					},
				},
			},
			wantByProjects: map[string][]session.Session{
				"my-todo": {
					{
						Id:        "1",
						StartTime: time.Date(2020, 1, 1, 8, 0, 0, 0, time.UTC),
						EndTime:   time.Date(2020, 1, 1, 10, 0, 0, 0, time.UTC),
						Project:   "my-todo",
					},
				},
				"flow": {
					{
						Id:        "2",
						StartTime: time.Date(2020, 1, 2, 12, 0, 0, 0, time.UTC),
						EndTime:   time.Date(2020, 1, 2, 10, 0, 0, 0, time.UTC),
						Project:   "flow",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.SplitSessionsByDay(); !reflect.DeepEqual(got, tt.wantByDays) {
				t.Errorf("Entry.SplitSessionsByDays() = %v, want %v", got, tt.wantByDays)
			}

			if got := tt.e.SplitSessionsByProject(); !reflect.DeepEqual(got, tt.wantByProjects) {
				t.Errorf("Entry.SplitSessionsByProjects() = %v, want %v", got, tt.wantByProjects)
			}
		})
	}
}

func TestSessionsReport_FindUniqueTags(t *testing.T) {
	tests := []struct {
		name string
		e    sessionsreport.SessionsReport
		want []string
	}{
		{
			name: "test",
			e:    sessionsReportTest,
			want: []string{"add-todo", "remove-todo", "start-usecase", "stop-usecase"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.FindUniqueTags(tt.e.Sessions)
			slices.Sort(got)
			slices.Sort(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Entry.FindUniqueTags() = %v, want %v", got, tt.want)
			}
		})
	}
}
