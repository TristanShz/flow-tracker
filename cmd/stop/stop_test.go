package stop_test

import (
	"testing"
	"time"

	"github.com/TristanSch1/flow/cmd/stop"
	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/stopsession"
	"github.com/TristanSch1/flow/internal/domain/session"
	"github.com/TristanSch1/flow/internal/infra"
	"github.com/TristanSch1/flow/test"
	"github.com/matryer/is"
)

func TestStopCommand(t *testing.T) {
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
			name:  "No sessions",
			args:  []string{},
			error: stopsession.ErrNoCurrentSession,
		},
		{
			name: "Session flowing",
			args: []string{},
			givenSessions: []session.Session{
				{
					Id:        "1",
					StartTime: time.Date(2024, time.April, 13, 17, 20, 0, 0, time.UTC),
					Project:   "Flow",
					Tags:      []string{"stop"},
				},
			},
			givenNow: time.Date(2024, time.April, 13, 17, 30, 0, 0, time.UTC),
			want:     "Flow session stopped, you were in the flow for 10m0s",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			sessionRepository.Sessions = tc.givenSessions
			dateProvider.Now = tc.givenNow
			c := stop.Command(app)

			got, err := test.ExecuteCmd(t, c, tc.args...)

			is.Equal(err, tc.error)

			if tc.error == nil {
				is.Equal(got, tc.want)
			}
		})
	}
}
