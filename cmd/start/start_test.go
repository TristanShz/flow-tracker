package start_test

import (
	"testing"
	"time"

	"github.com/TristanSch1/flow/cmd/start"
	"github.com/TristanSch1/flow/internal/domain/session"
	"github.com/TristanSch1/flow/internal/infra"
	"github.com/TristanSch1/flow/test"
	"github.com/matryer/is"
)

func TestStarTCommand(t *testing.T) {
	is := is.New(t)

	sessionRepository := &infra.InMemorySessionRepository{}
	app := test.InitializeApp(sessionRepository)

	tt := []struct {
		name          string
		want          string
		args          []string
		error         error
		givenSessions []session.Session
	}{
		{
			name:          "No args and no existing project",
			args:          []string{},
			want:          "Please provide a project name",
			error:         nil,
			givenSessions: []session.Session{},
		},
		{
			name:  "No args and existing project",
			args:  []string{},
			want:  "Please provide a project name, existing projects: MyTodo",
			error: nil,
			givenSessions: []session.Session{
				{
					Id:        "1",
					StartTime: time.Date(2024, time.April, 14, 10, 12, 0, 0, time.UTC),
					EndTime:   time.Date(2024, time.April, 14, 13, 10, 0, 0, time.UTC),
					Project:   "MyTodo",
					Tags:      []string{"add-todo"},
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			sessionRepository.Sessions = tc.givenSessions
			c := start.Command(app)

			got, err := test.ExecuteCmd(t, c, tc.args...)

			is.Equal(err, tc.error)

			if tc.error == nil {
				is.Equal(got, tc.want)
			}
		})
	}
}
