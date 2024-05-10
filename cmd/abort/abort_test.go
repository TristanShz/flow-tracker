package abort_test

import (
	"testing"
	"time"

	"github.com/TristanShz/flow/cmd/abort"
	"github.com/TristanShz/flow/internal/domain/session"
	"github.com/TristanShz/flow/internal/infra"
	"github.com/TristanShz/flow/test"
	"github.com/matryer/is"
)

func TestAbortCommand(t *testing.T) {
	sessionRepository := &infra.InMemorySessionRepository{}
	dateProvider := infra.NewStubDateProvider()
	app := test.InitializeApp(sessionRepository, dateProvider)

	tt := []struct {
		name          string
		args          []string
		givenSessions []session.Session
		want          string
		error         error
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
			want: "Session aborted",
		},
		{
			name: "No active session",
			want: "no active session",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			is := is.New(t)

			sessionRepository.Sessions = tc.givenSessions

			c := abort.Command(app)

			got, err := test.ExecuteCmd(t, c, tc.args...)

			is.Equal(tc.error, err)

			if tc.error == nil {
				is.Equal(tc.want, got)
			}
		})
	}
}
