package projectsessionsreport_test

import (
	"testing"
	"time"

	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/projectsessionsreport"
	"github.com/TristanSch1/flow/internal/domain/session"
	"github.com/TristanSch1/flow/internal/tests"
)

func TestProjectSessionsReport_Success(t *testing.T) {
	f := tests.GetSessionFixture(t)

	f.GivenSomeSessions([]session.Session{{
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
	}})

	f.WhenUserSeesProjectSessionsReport(projectsessionsreport.Command{Project: "Flow"})

	f.ThenUserShouldSeeProjectSessionsReport(
		projectsessionsreport.ProjectSessionReport{
			Total:            3*time.Hour + 0*time.Minute + 30*time.Second,
			NumberOfSessions: 2,
		},
	)
}

func TestProjectSessionsReport_NoSessions(t *testing.T) {
	f := tests.GetSessionFixture(t)

	f.WhenUserSeesProjectSessionsReport(projectsessionsreport.Command{Project: "Flow"})

	f.ThenUserShouldSeeProjectSessionsReport(
		projectsessionsreport.ProjectSessionReport{
			Total:            0,
			NumberOfSessions: 0,
		},
	)
}
