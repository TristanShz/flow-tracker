package startflowsession_test

import (
	"testing"
	"time"

	startflowsession "github.com/TristanSch1/flow/internal/application/usecases/start_flow_session"
	"github.com/TristanSch1/flow/internal/domain/session"
	"github.com/TristanSch1/flow/internal/tests"
)

func TestStartFlowSession_Success(t *testing.T) {
	f := tests.GetSessionFixture(t)

	f.GivenPredefinedStartTime(time.Date(2024, time.April, 13, 17, 20, 0, 0, time.UTC))

	command := startflowsession.Command{
		Project: "Flow",
		Tags:    []string{"start"},
	}

	f.WhenStartingFlowSession(command)

	f.ThenSessionWithGivenStartTimeShouldBeSaved(time.Date(2024, time.April, 13, 17, 20, 0, 0, time.UTC))
}

func TestStartFlowSession_AlreadyStarted(t *testing.T) {
	f := tests.GetSessionFixture(t)

	f.GivenSomeSessions([]session.Session{{
		StartTime: time.Date(2024, time.April, 13, 17, 20, 0, 0, time.UTC),
		Project:   "Flow",
		Tags:      []string{"start"},
	}})

	command := startflowsession.Command{
		Project: "Flow",
		Tags:    []string{"already_started"},
	}

	f.WhenStartingFlowSession(command)

	f.ThenErrorShouldBe(startflowsession.ErrSessionAlreadyStarted)
}
