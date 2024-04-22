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
		StartTime: time.Date(2024, time.April, 14, 10, 12, 0, 0, time.UTC),
		EndTime:   time.Date(2024, time.April, 14, 13, 12, 0, 0, time.UTC),
		Project:   "MyTodo",
		Tags:      []string{"add-todo"},
	},
	{
		StartTime: time.Date(2024, time.April, 14, 14, 12, 0, 0, time.UTC),
		EndTime:   time.Date(2024, time.April, 14, 15, 12, 0, 0, time.UTC),
		Project:   "Flow",
		Tags:      []string{"start-usecase"},
	},
	{
		StartTime: time.Date(2024, time.April, 14, 16, 24, 0, 0, time.UTC),
		EndTime:   time.Date(2024, time.April, 14, 18, 24, 30, 0, time.UTC),
		Project:   "Flow",
		Tags:      []string{"report-usecase"},
	},
	{
		StartTime: time.Date(2024, time.April, 15, 8, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2024, time.April, 15, 12, 0, 0, 0, time.UTC),
		Project:   "MyTodo",
		Tags:      []string{"edit-todo"},
	},
	{
		StartTime: time.Date(2024, time.April, 15, 14, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2024, time.April, 15, 16, 0, 0, 0, time.UTC),
		Project:   "Flow",
		Tags:      []string{"start-usecase"},
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
			name:          "No arguments given",
			command:       viewsessionsreport.Command{},
			givenSessions: sessionsForTest,
			want: sessionsreport.NewSessionsReport([]session.Session{
				{
					StartTime: time.Date(2024, time.April, 14, 10, 12, 0, 0, time.UTC),
					EndTime:   time.Date(2024, time.April, 14, 13, 12, 0, 0, time.UTC),
					Project:   "MyTodo",
					Tags:      []string{"add-todo"},
				},
				{
					StartTime: time.Date(2024, time.April, 14, 14, 12, 0, 0, time.UTC),
					EndTime:   time.Date(2024, time.April, 14, 15, 12, 0, 0, time.UTC),
					Project:   "Flow",
					Tags:      []string{"start-usecase"},
				},
				{
					StartTime: time.Date(2024, time.April, 14, 16, 24, 0, 0, time.UTC),
					EndTime:   time.Date(2024, time.April, 14, 18, 24, 30, 0, time.UTC),
					Project:   "Flow",
					Tags:      []string{"report-usecase"},
				},
				{
					StartTime: time.Date(2024, time.April, 15, 8, 0, 0, 0, time.UTC),
					EndTime:   time.Date(2024, time.April, 15, 12, 0, 0, 0, time.UTC),
					Project:   "MyTodo",
					Tags:      []string{"edit-todo"},
				},
				{
					StartTime: time.Date(2024, time.April, 15, 14, 0, 0, 0, time.UTC),
					EndTime:   time.Date(2024, time.April, 15, 16, 0, 0, 0, time.UTC),
					Project:   "Flow",
					Tags:      []string{"start-usecase"},
				},
			}),
			expectedFormat: sessionsreport.FormatByProject,
		},
		{
			name:           "No arguments and no sessions",
			command:        viewsessionsreport.Command{},
			givenSessions:  []session.Session{},
			want:           sessionsreport.NewSessionsReport([]session.Session{}),
			expectedFormat: sessionsreport.FormatByProject,
		},
		{
			name: "Project argument given",
			command: viewsessionsreport.Command{
				Project: "Flow",
			},
			givenSessions: sessionsForTest,
			want: sessionsreport.NewSessionsReport([]session.Session{
				{
					StartTime: time.Date(2024, time.April, 14, 14, 12, 0, 0, time.UTC),
					EndTime:   time.Date(2024, time.April, 14, 15, 12, 0, 0, time.UTC),
					Project:   "Flow",
					Tags:      []string{"start-usecase"},
				},
				{
					StartTime: time.Date(2024, time.April, 14, 16, 24, 0, 0, time.UTC),
					EndTime:   time.Date(2024, time.April, 14, 18, 24, 30, 0, time.UTC),
					Project:   "Flow",
					Tags:      []string{"report-usecase"},
				},
				{
					StartTime: time.Date(2024, time.April, 15, 14, 0, 0, 0, time.UTC),
					EndTime:   time.Date(2024, time.April, 15, 16, 0, 0, 0, time.UTC),
					Project:   "Flow",
					Tags:      []string{"start-usecase"},
				},
			}),
			expectedFormat: sessionsreport.FormatByProject,
		},
		{
			name: "Format by day",
			command: viewsessionsreport.Command{
				Format: sessionsreport.FormatByDay,
			},
			givenSessions: sessionsForTest,
			want: sessionsreport.NewSessionsReport([]session.Session{
				{
					StartTime: time.Date(2024, time.April, 14, 10, 12, 0, 0, time.UTC),
					EndTime:   time.Date(2024, time.April, 14, 13, 12, 0, 0, time.UTC),
					Project:   "MyTodo",
					Tags:      []string{"add-todo"},
				},
				{
					StartTime: time.Date(2024, time.April, 14, 14, 12, 0, 0, time.UTC),
					EndTime:   time.Date(2024, time.April, 14, 15, 12, 0, 0, time.UTC),
					Project:   "Flow",
					Tags:      []string{"start-usecase"},
				},
				{
					StartTime: time.Date(2024, time.April, 14, 16, 24, 0, 0, time.UTC),
					EndTime:   time.Date(2024, time.April, 14, 18, 24, 30, 0, time.UTC),
					Project:   "Flow",
					Tags:      []string{"report-usecase"},
				},
				{
					StartTime: time.Date(2024, time.April, 15, 8, 0, 0, 0, time.UTC),
					EndTime:   time.Date(2024, time.April, 15, 12, 0, 0, 0, time.UTC),
					Project:   "MyTodo",
					Tags:      []string{"edit-todo"},
				},
				{
					StartTime: time.Date(2024, time.April, 15, 14, 0, 0, 0, time.UTC),
					EndTime:   time.Date(2024, time.April, 15, 16, 0, 0, 0, time.UTC),
					Project:   "Flow",
					Tags:      []string{"start-usecase"},
				},
			}),
			expectedFormat: sessionsreport.FormatByDay,
		},
		{
			name: "View sessions of a given day",
			command: viewsessionsreport.Command{
				From: time.Date(2024, time.April, 14, 0, 0, 0, 0, time.UTC),
				To:   time.Date(2024, time.April, 14, 23, 59, 59, 0, time.UTC),
			},
			givenSessions: sessionsForTest,
			want: sessionsreport.NewSessionsReport([]session.Session{
				{
					StartTime: time.Date(2024, time.April, 14, 10, 12, 0, 0, time.UTC),
					EndTime:   time.Date(2024, time.April, 14, 13, 12, 0, 0, time.UTC),
					Project:   "MyTodo",
					Tags:      []string{"add-todo"},
				},
				{
					StartTime: time.Date(2024, time.April, 14, 14, 12, 0, 0, time.UTC),
					EndTime:   time.Date(2024, time.April, 14, 15, 12, 0, 0, time.UTC),
					Project:   "Flow",
					Tags:      []string{"start-usecase"},
				},
				{
					StartTime: time.Date(2024, time.April, 14, 16, 24, 0, 0, time.UTC),
					EndTime:   time.Date(2024, time.April, 14, 18, 24, 30, 0, time.UTC),
					Project:   "Flow",
					Tags:      []string{"report-usecase"},
				},
			}),
			expectedFormat: sessionsreport.FormatByProject,
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
