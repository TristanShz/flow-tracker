package abortsession_test

import (
	"testing"
	"time"

	abortsession "github.com/TristanShz/flow/internal/application/usecases/flowsession/abort"
	"github.com/TristanShz/flow/internal/domain/session"
	"github.com/TristanShz/flow/internal/tests"
)

func TestAbort(t *testing.T) {
	tt := []struct {
		error         error
		name          string
		givenSessions []session.Session
	}{
		{
			name: "Active session",
			givenSessions: []session.Session{
				{
					Id:        "1",
					StartTime: time.Date(2024, time.April, 13, 17, 20, 0, 0, time.UTC),
					Project:   "Flow",
				},
			},
		},
		{
			name: "No active session",
			givenSessions: []session.Session{
				{
					Id:        "1",
					StartTime: time.Date(2024, time.April, 13, 17, 20, 0, 0, time.UTC),
					EndTime:   time.Date(2024, time.April, 13, 18, 20, 0, 0, time.UTC),
					Project:   "Flow",
				},
			},
			error: abortsession.ErrNoActiveSession,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			f := tests.GetSessionFixture(t)

			f.GivenSomeSessions(tc.givenSessions)

			f.WhenAbortingFlowSession()

			f.ThenNoSessionShouldBeActive()
		})
	}
}
