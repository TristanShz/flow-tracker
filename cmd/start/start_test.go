package start_test

import (
	"errors"
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
	dateProvider := infra.NewStubDateProvider()
	app := test.InitializeApp(sessionRepository, dateProvider)

	tt := []struct {
		name          string
		want          string
		args          []string
		error         error
		givenSessions []session.Session
		givenNow      time.Time
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
		{
			name:  "First arg is a tag",
			args:  []string{"+add-todo"},
			error: errors.New("the first argument must be the project name"),
		},
		{
			name:  "Invalid tag",
			args:  []string{"my-todo", "add-todo"},
			error: errors.New("invalid tag add-todo (must start with '+')"),
		},
		{
			name:     "Valid command with project",
			args:     []string{"my-todo"},
			givenNow: time.Date(2024, time.April, 14, 10, 12, 0, 0, time.UTC),
			want:     "Starting flow session for the project my-todo at 10:12AM",
		},
		{
			name:     "Valid command with project and tags",
			args:     []string{"my-todo", "+add-todo", "+update-todo"},
			givenNow: time.Date(2024, time.April, 14, 10, 12, 0, 0, time.UTC),
			want:     "Starting flow session for the project my-todo [add-todo, update-todo] at 10:12AM",
		},
		{
			name:     "Session already started",
			args:     []string{"my-todo"},
			error:    errors.New("there is already a session in progress"),
			givenNow: time.Date(2024, time.April, 14, 10, 12, 0, 0, time.UTC),
			givenSessions: []session.Session{
				{
					Id:        "1",
					StartTime: time.Date(2024, time.April, 14, 10, 12, 0, 0, time.UTC),
					Project:   "my-todo",
					Tags:      []string{"add-todo"},
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			sessionRepository.Sessions = tc.givenSessions
			dateProvider.Now = tc.givenNow
			c := start.Command(app)

			got, err := test.ExecuteCmd(t, c, tc.args...)

			is.Equal(err, tc.error)

			if tc.error == nil {
				is.Equal(got, tc.want)
			}
		})
	}
}
