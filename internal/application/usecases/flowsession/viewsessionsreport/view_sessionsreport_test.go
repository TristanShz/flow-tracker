package viewsessionsreport_test

import (
	"testing"
	"time"

	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/viewsessionsreport"
	"github.com/TristanSch1/flow/internal/domain/session"
	"github.com/TristanSch1/flow/internal/domain/sessionsreport"
	"github.com/TristanSch1/flow/internal/tests"
)

var sessionsForTest = []session.Session{{
	StartTime: time.Date(2024, time.April, 14, 10, 12, 0, 0, time.UTC),
	EndTime:   time.Date(2024, time.April, 14, 13, 12, 0, 0, time.UTC),
	Project:   "MyTodo",
	Tags:      []string{"add-todo"},
}, {
	StartTime: time.Date(2024, time.April, 14, 14, 12, 0, 0, time.UTC),
	EndTime:   time.Date(2024, time.April, 14, 15, 12, 0, 0, time.UTC),
	Project:   "Flow",
	Tags:      []string{"start-usecase"},
}, {
	StartTime: time.Date(2024, time.April, 14, 16, 24, 0, 0, time.UTC),
	EndTime:   time.Date(2024, time.April, 14, 18, 24, 30, 0, time.UTC),
	Project:   "Flow",
	Tags:      []string{"report-usecase"},
}}

func TestViewSessionsReport(t *testing.T) {
	f := tests.GetSessionFixture(t)

	tests := []struct {
		name           string
		command        viewsessionsreport.Command
		givenSessions  []session.Session
		want           sessionsreport.SessionsReport
		expectedFormat string
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
				}, {
					StartTime: time.Date(2024, time.April, 14, 14, 12, 0, 0, time.UTC),
					EndTime:   time.Date(2024, time.April, 14, 15, 12, 0, 0, time.UTC),
					Project:   "Flow",
					Tags:      []string{"start-usecase"},
				}, {
					StartTime: time.Date(2024, time.April, 14, 16, 24, 0, 0, time.UTC),
					EndTime:   time.Date(2024, time.April, 14, 18, 24, 30, 0, time.UTC),
					Project:   "Flow",
					Tags:      []string{"report-usecase"},
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
				}, {
					StartTime: time.Date(2024, time.April, 14, 16, 24, 0, 0, time.UTC),
					EndTime:   time.Date(2024, time.April, 14, 18, 24, 30, 0, time.UTC),
					Project:   "Flow",
					Tags:      []string{"report-usecase"},
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
				}, {
					StartTime: time.Date(2024, time.April, 14, 14, 12, 0, 0, time.UTC),
					EndTime:   time.Date(2024, time.April, 14, 15, 12, 0, 0, time.UTC),
					Project:   "Flow",
					Tags:      []string{"start-usecase"},
				}, {
					StartTime: time.Date(2024, time.April, 14, 16, 24, 0, 0, time.UTC),
					EndTime:   time.Date(2024, time.April, 14, 18, 24, 30, 0, time.UTC),
					Project:   "Flow",
					Tags:      []string{"report-usecase"},
				},
			}),
			expectedFormat: sessionsreport.FormatByDay,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f.GivenSomeSessions(tt.givenSessions)

			f.WhenUserSeesSessionsReport(tt.command)

			f.ThenUserShouldSeeSessionsReport(tt.want, tt.expectedFormat)
		})
	}
}
