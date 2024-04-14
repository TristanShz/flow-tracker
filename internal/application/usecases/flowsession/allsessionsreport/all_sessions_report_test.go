package allsessionsreport_test

import (
	"testing"
	"time"

	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/allsessionsreport"
	"github.com/TristanSch1/flow/internal/domain/session"
	"github.com/TristanSch1/flow/internal/tests"
)

func TestFlowSessionsReport_Globsl_Success(t *testing.T) {
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

	f.WhenUserSeesFlowSessionsReport(allsessionsreport.Command{})

	f.ThenUserShouldSeeReport(
		allsessionsreport.AllSessionsReport{
			Projects: map[string]time.Duration{
				"MyTodo": 3*time.Hour + 0*time.Minute + 0*time.Second,
				"Flow":   3*time.Hour + 0*time.Minute + 30*time.Second,
			},
			Total:            6*time.Hour + 0*time.Minute + 30*time.Second,
			NumberOfSessions: 3,
		},
	)
}
