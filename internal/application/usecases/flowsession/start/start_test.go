package start_test

import (
	"testing"
	"time"

	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/start"
	"github.com/TristanSch1/flow/internal/domain/session"
	"github.com/TristanSch1/flow/internal/tests"
)

func TestStartFlowSession_Success(t *testing.T) {
	f := tests.GetSessionFixture(t)

	f.GivenNowIs(time.Date(2024, time.April, 13, 17, 20, 0, 0, time.UTC))

	command := start.Command{
		Project: "Flow",
		Tags:    []string{"start"},
	}

	f.WhenStartingFlowSession(command)

	f.ThenSessionShouldBeSaved(session.Session{
		StartTime: time.Date(2024, time.April, 13, 17, 20, 0, 0, time.UTC),
		Project:   "Flow",
		Tags:      []string{"start"},
	})
}

func TestStartFlowSession_AlreadyStarted(t *testing.T) {
	f := tests.GetSessionFixture(t)

	f.GivenSomeSessions([]session.Session{{
		StartTime: time.Date(2024, time.April, 13, 17, 20, 0, 0, time.UTC),
		Project:   "Flow",
		Tags:      []string{"start"},
	}})

	command := start.Command{
		Project: "Flow",
		Tags:    []string{"already_started"},
	}

	f.WhenStartingFlowSession(command)

	f.ThenErrorShouldBe(start.ErrSessionAlreadyStarted)
}
