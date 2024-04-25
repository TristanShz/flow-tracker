package stopsession_test

import (
	"testing"
	"time"

	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/stopsession"
	"github.com/TristanSch1/flow/internal/domain/session"
	"github.com/TristanSch1/flow/internal/tests"
)

func TestStopFlowSession_Success(t *testing.T) {
	f := tests.GetSessionFixture(t)

	f.GivenSomeSessions([]session.Session{{
		StartTime: time.Date(2024, time.April, 13, 17, 20, 0, 0, time.UTC),
		Project:   "Flow",
		Tags:      []string{"stop"},
	}})

	f.WhenStoppingFlowSession()

	f.ThenSessionShouldBeStopped()
}

func TestStopFlowSession_NoCurrentSession(t *testing.T) {
	f := tests.GetSessionFixture(t)

	f.WhenStoppingFlowSession()

	f.ThenErrorShouldBe(stopsession.ErrNoCurrentSession)
}
