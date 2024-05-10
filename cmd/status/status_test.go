package status_test

import (
	"testing"
	"time"

	"github.com/TristanShz/flow/cmd/status"
	"github.com/TristanShz/flow/internal/domain/session"
	"github.com/TristanShz/flow/internal/infra"
	"github.com/TristanShz/flow/test"
	"github.com/matryer/is"
)

func TestStatusCommand(t *testing.T) {
	is := is.New(t)

	sessionRepository := &infra.InMemorySessionRepository{}
	dateProvider := infra.NewStubDateProvider()
	app := test.InitializeApp(sessionRepository, dateProvider)

	tt := []struct {
		givenNow      time.Time
		error         error
		name          string
		want          string
		args          []string
		givenSessions []session.Session
	}{
		{
			name:          "No current session",
			givenSessions: []session.Session{},
			want:          "No active flow session",
		},
		{
			name: "Current session",
			givenSessions: []session.Session{
				{
					Id:        "1",
					StartTime: time.Date(2024, time.April, 13, 17, 20, 0, 0, time.UTC),
					Project:   "Flow",
					Tags:      []string{"status"},
				},
			},
			givenNow: time.Date(2024, time.April, 13, 17, 30, 0, 0, time.UTC),
			want:     "You're in the flow for 10m0s on project Flow with tags: status",
		},
		{
			name: "Current session with multiple tags",
			givenSessions: []session.Session{
				{
					Id:        "1",
					StartTime: time.Date(2024, time.April, 13, 17, 20, 0, 0, time.UTC),
					Project:   "Flow",
					Tags:      []string{"status", "stop"},
				},
			},
			givenNow: time.Date(2024, time.April, 13, 18, 30, 0, 0, time.UTC),
			want:     "You're in the flow for 1h10m0s on project Flow with tags: status, stop",
		},
		{
			name: "Current session with no tags",
			givenSessions: []session.Session{
				{
					Id:        "1",
					StartTime: time.Date(2024, time.April, 13, 17, 20, 0, 0, time.UTC),
					Project:   "Flow",
				},
			},
			givenNow: time.Date(2024, time.April, 13, 17, 30, 0, 0, time.UTC),
			want:     "You're in the flow for 10m0s on project Flow",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			sessionRepository.Sessions = tc.givenSessions
			dateProvider.Now = tc.givenNow
			c := status.Command(app)

			got, err := test.ExecuteCmd(t, c, tc.args...)

			is.Equal(err, tc.error)

			if tc.error == nil {
				is.Equal(got, tc.want)
			}
		})
	}
}
