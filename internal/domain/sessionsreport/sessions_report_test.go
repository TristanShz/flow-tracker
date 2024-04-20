package sessionsreport_test

import (
	"reflect"
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

func TestSessionsReport_Formats(t *testing.T) {
	tests := []struct {
		wantByDays     []sessionsreport.DayReport
		wantByProjects []sessionsreport.ProjectReport
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
			wantByDays: []sessionsreport.DayReport{
				{
					Day: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
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
			},
			wantByProjects: []sessionsreport.ProjectReport{
				{
					Project: "my-todo",
					DurationByTag: map[string]time.Duration{
						"add-todo": 2*time.Hour + 50*time.Minute,
					},
					TotalDuration: 2*time.Hour + 50*time.Minute,
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
			wantByDays: []sessionsreport.DayReport{
				{
					Day: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
					Sessions: []session.Session{
						{
							StartTime: time.Date(2020, 1, 1, 8, 0, 0, 0, time.UTC),
							Project:   "my-todo",
							Tags:      []string{"add-todo"},
						},
					},
				},
			},
			wantByProjects: []sessionsreport.ProjectReport{
				{
					Project: "my-todo",
					DurationByTag: map[string]time.Duration{
						"add-todo": 0,
					},
				},
			},
		},
		{
			name:           "No sessions",
			e:              sessionsreport.NewSessionsReport([]session.Session{}),
			wantByDays:     []sessionsreport.DayReport{},
			wantByProjects: []sessionsreport.ProjectReport{},
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
					StartTime: time.Date(2020, 1, 2, 10, 0, 0, 0, time.UTC),
					EndTime:   time.Date(2020, 1, 2, 12, 0, 0, 0, time.UTC),
					Project:   "flow",
				},
			}),
			wantByDays: []sessionsreport.DayReport{
				{
					Day: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
					Sessions: []session.Session{
						{
							Id:        "1",
							StartTime: time.Date(2020, 1, 1, 8, 0, 0, 0, time.UTC),
							EndTime:   time.Date(2020, 1, 1, 10, 0, 0, 0, time.UTC),
							Project:   "my-todo",
						},
					},
				},
				{
					Day: time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC),
					Sessions: []session.Session{
						{
							Id:        "2",
							StartTime: time.Date(2020, 1, 2, 10, 0, 0, 0, time.UTC),
							EndTime:   time.Date(2020, 1, 2, 12, 0, 0, 0, time.UTC),
							Project:   "flow",
						},
					},
				},
			},
			wantByProjects: []sessionsreport.ProjectReport{
				{
					Project:       "my-todo",
					DurationByTag: map[string]time.Duration{},
					TotalDuration: 2 * time.Hour,
				}, {
					Project:       "flow",
					DurationByTag: map[string]time.Duration{},
					TotalDuration: 2 * time.Hour,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.GetByDayReport(); !reflect.DeepEqual(got, tt.wantByDays) {
				t.Errorf("Entry.SplitSessionsByDays() = %v, want %v", got, tt.wantByDays)
			}

			if got := tt.e.GetByProjectReport(); !reflect.DeepEqual(got, tt.wantByProjects) {
				t.Errorf("Entry.SplitSessionsByProjects() = %v, want %v", got, tt.wantByProjects)
			}
		})
	}
}
