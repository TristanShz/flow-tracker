package start_test

import (
	"errors"
	"testing"
	"time"

	"github.com/TristanShz/flow/cmd/start"
	"github.com/TristanShz/flow/internal/domain/session"
	"github.com/TristanShz/flow/internal/infra"
	"github.com/TristanShz/flow/test"
	"github.com/matryer/is"
)

func TestStartCommand(t *testing.T) {
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
			want:  "Please provide a project name, existing projects: MyTodo, Flow",
			error: nil,
			givenSessions: []session.Session{
				{
					Id:        "1",
					StartTime: time.Date(2024, time.April, 14, 10, 12, 0, 0, time.UTC),
					EndTime:   time.Date(2024, time.April, 14, 13, 10, 0, 0, time.UTC),
					Project:   "MyTodo",
					Tags:      []string{"add-todo"},
				},
				{
					Id:        "2",
					StartTime: time.Date(2024, time.April, 14, 14, 12, 0, 0, time.UTC),
					EndTime:   time.Date(2024, time.April, 14, 16, 10, 0, 0, time.UTC),
					Project:   "Flow",
					Tags:      []string{"start-session"},
				},
			},
		},
		{
			name: "No args and existig project with active session",
			args: []string{"Flow"},
			want: "There is already a session in progress",
			givenSessions: []session.Session{
				{
					Id:        "1",
					StartTime: time.Date(2024, time.April, 14, 10, 12, 0, 0, time.UTC),
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
			want:     "There is already a session in progress",
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
