package viewsessionsreport_test

import (
	"testing"
	"time"

	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/viewsessionsreport"
	"github.com/TristanSch1/flow/internal/domain/session"
	"github.com/TristanSch1/flow/internal/domain/sessionsreport"
	"github.com/TristanSch1/flow/internal/tests"
)

var sessionsForTest = []session.Session{
	{
		Id:        "1",
		StartTime: time.Date(2024, time.April, 14, 10, 12, 0, 0, time.UTC),
		EndTime:   time.Date(2024, time.April, 14, 13, 12, 0, 0, time.UTC),
		Project:   "MyTodo",
		Tags:      []string{"add-todo"},
	},
	{
		Id:        "2",
		StartTime: time.Date(2024, time.April, 14, 14, 12, 0, 0, time.UTC),
		EndTime:   time.Date(2024, time.April, 14, 15, 12, 0, 0, time.UTC),
		Project:   "Flow",
		Tags:      []string{"start-usecase"},
	},
	{
		Id:        "3",
		StartTime: time.Date(2024, time.April, 14, 16, 24, 0, 0, time.UTC),
		EndTime:   time.Date(2024, time.April, 14, 18, 24, 30, 0, time.UTC),
		Project:   "Flow",
		Tags:      []string{"report-usecase"},
	},
	{
		Id:        "4",
		StartTime: time.Date(2024, time.April, 15, 8, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2024, time.April, 15, 12, 0, 0, 0, time.UTC),
		Project:   "MyTodo",
		Tags:      []string{"edit-todo"},
	},
	{
		Id:        "5",
		StartTime: time.Date(2024, time.April, 15, 14, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2024, time.April, 15, 16, 0, 0, 0, time.UTC),
		Project:   "Flow",
		Tags:      []string{"start-usecase"},
	},
	{
		Id:        "6",
		StartTime: time.Date(2024, time.April, 16, 10, 12, 0, 0, time.UTC),
		EndTime:   time.Date(2024, time.April, 16, 13, 12, 0, 0, time.UTC),
		Project:   "MyTodo",
		Tags:      []string{"delete-todo"},
	},
	{
		Id:        "7",
		StartTime: time.Date(2024, time.April, 16, 14, 12, 0, 0, time.UTC),
		EndTime:   time.Date(2024, time.April, 16, 15, 12, 0, 0, time.UTC),
		Project:   "Pomodoro",
		Tags:      []string{"start-pomodoro"},
	},
	{
		Id:        "8",
		StartTime: time.Date(2024, time.April, 18, 16, 24, 0, 0, time.UTC),
		EndTime:   time.Date(2024, time.April, 18, 18, 24, 30, 0, time.UTC),
		Project:   "Pomodoro",
		Tags:      []string{"report-pomodoro"},
	},
	{
		Id:        "9",
		StartTime: time.Date(2024, time.April, 20, 8, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2024, time.April, 20, 12, 0, 0, 0, time.UTC),
		Project:   "MyTodo",
		Tags:      []string{"deploy"},
	},
	{
		Id:        "10",
		StartTime: time.Date(2024, time.April, 20, 14, 0, 0, 0, time.UTC),
		Project:   "Pomodoro",
		Tags:      []string{"pause-pomodoro"},
	},
}

func TestViewSessionsReport(t *testing.T) {
	testsTable := []struct {
		command        viewsessionsreport.Command
		name           string
		expectedFormat string
		givenSessions  []session.Session
		want           sessionsreport.SessionsReport
	}{
		{
			name:           "No arguments given",
			command:        viewsessionsreport.Command{},
			givenSessions:  sessionsForTest,
			want:           sessionsreport.NewSessionsReport(sessionsForTest),
			expectedFormat: sessionsreport.FormatByDay,
		},
		{
			name:           "No arguments and no sessions",
			command:        viewsessionsreport.Command{},
			givenSessions:  []session.Session{},
			want:           sessionsreport.NewSessionsReport([]session.Session{}),
			expectedFormat: sessionsreport.FormatByDay,
		},
		{
			name: "Project argument given",
			command: viewsessionsreport.Command{
				Project: "Flow",
			},
			givenSessions: sessionsForTest,
			want: sessionsreport.NewSessionsReport([]session.Session{
				{
					Id:        "2",
					StartTime: time.Date(2024, time.April, 14, 14, 12, 0, 0, time.UTC),
					EndTime:   time.Date(2024, time.April, 14, 15, 12, 0, 0, time.UTC),
					Project:   "Flow",
					Tags:      []string{"start-usecase"},
				},
				{
					Id:        "3",
					StartTime: time.Date(2024, time.April, 14, 16, 24, 0, 0, time.UTC),
					EndTime:   time.Date(2024, time.April, 14, 18, 24, 30, 0, time.UTC),
					Project:   "Flow",
					Tags:      []string{"report-usecase"},
				},
				{
					Id:        "5",
					StartTime: time.Date(2024, time.April, 15, 14, 0, 0, 0, time.UTC),
					EndTime:   time.Date(2024, time.April, 15, 16, 0, 0, 0, time.UTC),
					Project:   "Flow",
					Tags:      []string{"start-usecase"},
				},
			}),
			expectedFormat: sessionsreport.FormatByDay,
		},
		{
			name: "Format by day",
			command: viewsessionsreport.Command{
				Format: sessionsreport.FormatByDay,
			},
			givenSessions:  sessionsForTest,
			want:           sessionsreport.NewSessionsReport(sessionsForTest),
			expectedFormat: sessionsreport.FormatByDay,
		},
		{
			name: "View sessions of a given day",
			command: viewsessionsreport.Command{
				Since: time.Date(2024, time.April, 14, 0, 0, 0, 0, time.UTC),
				Until: time.Date(2024, time.April, 14, 23, 59, 59, 0, time.UTC),
			},
			givenSessions: sessionsForTest,
			want: sessionsreport.NewSessionsReport([]session.Session{
				{
					Id:        "1",
					StartTime: time.Date(2024, time.April, 14, 10, 12, 0, 0, time.UTC),
					EndTime:   time.Date(2024, time.April, 14, 13, 12, 0, 0, time.UTC),
					Project:   "MyTodo",
					Tags:      []string{"add-todo"},
				},
				{
					Id:        "2",
					StartTime: time.Date(2024, time.April, 14, 14, 12, 0, 0, time.UTC),
					EndTime:   time.Date(2024, time.April, 14, 15, 12, 0, 0, time.UTC),
					Project:   "Flow",
					Tags:      []string{"start-usecase"},
				},
				{
					Id:        "3",
					StartTime: time.Date(2024, time.April, 14, 16, 24, 0, 0, time.UTC),
					EndTime:   time.Date(2024, time.April, 14, 18, 24, 30, 0, time.UTC),
					Project:   "Flow",
					Tags:      []string{"report-usecase"},
				},
			}),
			expectedFormat: sessionsreport.FormatByDay,
		},
		{
			name: "View sessions since a given day",
			command: viewsessionsreport.Command{
				Since: time.Date(2024, time.April, 18, 0, 0, 0, 0, time.UTC),
			},
			givenSessions: sessionsForTest,
			want: sessionsreport.NewSessionsReport([]session.Session{
				{
					Id:        "8",
					StartTime: time.Date(2024, time.April, 18, 16, 24, 0, 0, time.UTC),
					EndTime:   time.Date(2024, time.April, 18, 18, 24, 30, 0, time.UTC),
					Project:   "Pomodoro",
					Tags:      []string{"report-pomodoro"},
				},
				{
					Id:        "9",
					StartTime: time.Date(2024, time.April, 20, 8, 0, 0, 0, time.UTC),
					EndTime:   time.Date(2024, time.April, 20, 12, 0, 0, 0, time.UTC),
					Project:   "MyTodo",
					Tags:      []string{"deploy"},
				},
				{
					Id:        "10",
					StartTime: time.Date(2024, time.April, 20, 14, 0, 0, 0, time.UTC),
					Project:   "Pomodoro",
					Tags:      []string{"pause-pomodoro"},
				},
			}),
			expectedFormat: sessionsreport.FormatByDay,
		},
		{
			name: "View sessions until a given day",
			command: viewsessionsreport.Command{
				Until: time.Date(2024, time.April, 15, 0, 0, 0, 0, time.UTC),
			},
			givenSessions: sessionsForTest,
			want: sessionsreport.NewSessionsReport([]session.Session{
				{
					Id:        "1",
					StartTime: time.Date(2024, time.April, 14, 10, 12, 0, 0, time.UTC),
					EndTime:   time.Date(2024, time.April, 14, 13, 12, 0, 0, time.UTC),
					Project:   "MyTodo",
					Tags:      []string{"add-todo"},
				},
				{
					Id:        "2",
					StartTime: time.Date(2024, time.April, 14, 14, 12, 0, 0, time.UTC),
					EndTime:   time.Date(2024, time.April, 14, 15, 12, 0, 0, time.UTC),
					Project:   "Flow",
					Tags:      []string{"start-usecase"},
				},
				{
					Id:        "3",
					StartTime: time.Date(2024, time.April, 14, 16, 24, 0, 0, time.UTC),
					EndTime:   time.Date(2024, time.April, 14, 18, 24, 30, 0, time.UTC),
					Project:   "Flow",
					Tags:      []string{"report-usecase"},
				},
			}),
			expectedFormat: sessionsreport.FormatByDay,
		},
	}

	for _, tt := range testsTable {
		t.Run(tt.name, func(t *testing.T) {
			f := tests.GetSessionFixture(t)

			f.GivenSomeSessions(tt.givenSessions)

			f.WhenUserSeesSessionsReport(tt.command)

			f.ThenUserShouldSeeSessionsReport(tt.want, tt.expectedFormat)
		})
	}
}
