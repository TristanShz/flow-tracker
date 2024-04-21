package status_test

import (
	"testing"
	"time"

	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/status"
	"github.com/TristanSch1/flow/internal/domain/session"
	"github.com/TristanSch1/flow/internal/tests"
)

func TestFlowSessionStatus_Success(t *testing.T) {
	f := tests.GetSessionFixture(t)

	f.GivenSomeSessions([]session.Session{{
		StartTime: time.Date(2024, time.April, 14, 11, 26, 0, 0, time.UTC),
		Project:   "Flow",
		Tags:      []string{"status"},
	}})

	f.GivenNowIs(time.Date(2024, time.April, 14, 12, 26, 0, 0, time.UTC))

	f.WhenUserSeesTheCurrentSessionStatus()

	f.ThenUserShouldSee(session.Session{
		StartTime: time.Date(2024, time.April, 14, 11, 26, 0, 0, time.UTC),
		Project:   "Flow",
		Tags:      []string{"status"},
	}, 1*time.Hour)
}

func TestFlowSessionStatus_NoCurrentSession(t *testing.T) {
	f := tests.GetSessionFixture(t)

	f.GivenSomeSessions([]session.Session{{
		StartTime: time.Date(2024, time.April, 14, 11, 26, 0, 0, time.UTC),
		EndTime:   time.Date(2024, time.April, 15, 11, 32, 0, 0, time.UTC),
		Project:   "Flow",
		Tags:      []string{"status"},
	}})

	f.WhenUserSeesTheCurrentSessionStatus()

	f.ThenErrorShouldBe(status.ErrNoCurrentSession)
}
