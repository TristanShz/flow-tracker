package status_test

import (
	"testing"
	"time"

	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/status"
	"github.com/TristanSch1/flow/internal/domain/session"
	"github.com/TristanSch1/flow/internal/tests"
)

func TestFlowSessionStatus(t *testing.T) {
	testsTable := []struct {
		name             string
		givenSessions    []session.Session
		givenNow         time.Time
		expectedSession  session.Session
		expectedDuration time.Duration
		expectedError    error
	}{
		{
			name: "Session is current",
			givenSessions: []session.Session{
				{
					Id:        "1",
					StartTime: time.Date(2024, time.April, 14, 11, 26, 0, 0, time.UTC),
					Project:   "Flow",
					Tags:      []string{"status"},
				},
			},
			givenNow: time.Date(2024, time.April, 14, 12, 26, 0, 0, time.UTC),
			expectedSession: session.Session{
				Id:        "1",
				StartTime: time.Date(2024, time.April, 14, 11, 26, 0, 0, time.UTC),
				Project:   "Flow",
				Tags:      []string{"status"},
			},
			expectedDuration: 1 * time.Hour,
			expectedError:    nil,
		},
		{
			name:             "No sessions",
			givenSessions:    []session.Session{},
			givenNow:         time.Date(2024, time.April, 14, 12, 26, 0, 0, time.UTC),
			expectedSession:  session.Session{},
			expectedDuration: 0,
			expectedError:    status.ErrNoCurrentSession,
		},
		{
			name: "No current session",
			givenSessions: []session.Session{
				{
					Id:        "1",
					StartTime: time.Date(2024, time.April, 14, 11, 26, 0, 0, time.UTC),
					EndTime:   time.Date(2024, time.April, 14, 12, 26, 0, 0, time.UTC),
					Project:   "Flow",
					Tags:      []string{"status"},
				},
			},
			givenNow:         time.Date(2024, time.April, 14, 13, 26, 0, 0, time.UTC),
			expectedSession:  session.Session{},
			expectedDuration: 0,
			expectedError:    status.ErrNoCurrentSession,
		},
	}

	for _, tt := range testsTable {
		t.Run(tt.name, func(t *testing.T) {
			f := tests.GetSessionFixture(t)
			f.GivenSomeSessions(tt.givenSessions)
			f.GivenNowIs(tt.givenNow)
			f.WhenUserSeesTheCurrentSessionStatus()
			if tt.expectedError != nil {
				f.ThenErrorShouldBe(tt.expectedError)
			} else {
				f.ThenUserShouldSee(tt.expectedSession, 1*time.Hour)
			}
		})
	}
}
