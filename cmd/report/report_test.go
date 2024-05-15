package report_test

import (
	"errors"
	"testing"
	"time"

	"github.com/TristanShz/flow/cmd/report"
	"github.com/TristanShz/flow/internal/domain/session"
	"github.com/TristanShz/flow/internal/infra"
	"github.com/TristanShz/flow/test"
	is "github.com/matryer/is"
)

func TestReportCommand(t *testing.T) {
	is := is.New(t)

	sessionRepository := &infra.InMemorySessionRepository{}
	dateProvider := infra.NewStubDateProvider()
	app := test.InitializeApp(sessionRepository, dateProvider)

	tt := []struct {
		error         error
		name          string
		want          string
		givenSessions []session.Session
		givenNow      time.Time
		args          []string
	}{
		{
			name: "No args",
			args: []string{},
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
					EndTime:   time.Date(2024, time.April, 14, 15, 12, 0, 0, time.UTC),
					Project:   "Flow",
					Tags:      []string{"start-usecase"},
				},
			},
			want: "Sessions Report\n\nSun, 14 Apr 2024 - 3h58m0s\n    1 10:12:00 to 13:10:00 2h58m0s MyTodo [add-todo]\n    2 14:12:00 to 15:12:00 1h0m0s Flow [start-usecase]",
		},
		{
			name:  "Invalid format flag",
			args:  []string{"--format", "invalid"},
			error: errors.New("invalid format flag. possible values: by-day, by-project"),
		},
		{
			name: "By day",
			args: []string{"--format", "by-day"},
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
					EndTime:   time.Date(2024, time.April, 14, 15, 12, 0, 0, time.UTC),
					Project:   "Flow",
					Tags:      []string{"start-usecase"},
				},
			},
			want: "Sessions Report\n\nSun, 14 Apr 2024 - 3h58m0s\n    1 10:12:00 to 13:10:00 2h58m0s MyTodo [add-todo]\n    2 14:12:00 to 15:12:00 1h0m0s Flow [start-usecase]",
		},
		{
			name: "By project",
			args: []string{"--format", "by-project"},
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
					EndTime:   time.Date(2024, time.April, 14, 15, 12, 0, 0, time.UTC),
					Project:   "Flow",
					Tags:      []string{"start-usecase"},
				},
			},
			want: "Sessions Report\n\nMyTodo - 2h58m0s\n    [add-todo] -> 2h58m0s\n\nFlow - 1h0m0s\n    [start-usecase] -> 1h0m0s",
		},
		{
			name: "Sessions of project",
			args: []string{"--project", "MyTodo"},
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
					EndTime:   time.Date(2024, time.April, 14, 15, 12, 0, 0, time.UTC),
					Project:   "Flow",
					Tags:      []string{"start-usecase"},
				},
				{
					Id:        "3",
					StartTime: time.Date(2024, time.April, 15, 16, 12, 0, 0, time.UTC),
					EndTime:   time.Date(2024, time.April, 15, 17, 12, 0, 0, time.UTC),
					Project:   "Flow",
					Tags:      []string{"start-usecase"},
				},
			},
			want: "Sessions Report\n\nSun, 14 Apr 2024 - 2h58m0s\n    1 10:12:00 to 13:10:00 2h58m0s MyTodo [add-todo]",
		},
		{
			name: "Sessions of project but project does not exist",
			args: []string{"--project", "unknown-project"},
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
					EndTime:   time.Date(2024, time.April, 14, 15, 12, 0, 0, time.UTC),
					Project:   "Flow",
					Tags:      []string{"start-usecase"},
				},
				{
					Id:        "3",
					StartTime: time.Date(2024, time.April, 15, 16, 12, 0, 0, time.UTC),
					EndTime:   time.Date(2024, time.April, 15, 17, 12, 0, 0, time.UTC),
					Project:   "Flow",
					Tags:      []string{"start-usecase"},
				},
			},
			want: "No sessions found",
		},
		{
			name:     "Sessions of the day",
			args:     []string{"--day"},
			givenNow: time.Date(2024, time.April, 14, 18, 0, 0, 0, time.UTC),
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
					EndTime:   time.Date(2024, time.April, 14, 15, 12, 0, 0, time.UTC),
					Project:   "Flow",
					Tags:      []string{"start-usecase"},
				},
				{
					Id:        "3",
					StartTime: time.Date(2024, time.April, 15, 16, 12, 0, 0, time.UTC),
					EndTime:   time.Date(2024, time.April, 15, 17, 12, 0, 0, time.UTC),
					Project:   "Flow",
					Tags:      []string{"start-usecase"},
				},
			},
			want: "Sessions Report\n\nSun, 14 Apr 2024 - 3h58m0s\n    1 10:12:00 to 13:10:00 2h58m0s MyTodo [add-todo]\n    2 14:12:00 to 15:12:00 1h0m0s Flow [start-usecase]",
		},
		{
			name:     "Sessions of the week",
			args:     []string{"--week"},
			givenNow: time.Date(2024, time.April, 16, 18, 0, 0, 0, time.UTC),
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
					EndTime:   time.Date(2024, time.April, 14, 15, 12, 0, 0, time.UTC),
					Project:   "Flow",
					Tags:      []string{"start-usecase"},
				},
				{
					Id:        "3",
					StartTime: time.Date(2024, time.April, 15, 16, 12, 0, 0, time.UTC),
					EndTime:   time.Date(2024, time.April, 15, 17, 12, 0, 0, time.UTC),
					Project:   "Flow",
					Tags:      []string{"start-usecase"},
				},
			},
			want: "Sessions Report\n\nMon, 15 Apr 2024 - 1h0m0s\n    3 16:12:00 to 17:12:00 1h0m0s Flow [start-usecase]",
		},
		{
			name: "Since flag",
			args: []string{"--since", "2024-04-15"},
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
					EndTime:   time.Date(2024, time.April, 14, 15, 12, 0, 0, time.UTC),
					Project:   "Flow",
					Tags:      []string{"start-usecase"},
				},
				{
					Id:        "3",
					StartTime: time.Date(2024, time.April, 15, 16, 12, 0, 0, time.UTC),
					EndTime:   time.Date(2024, time.April, 15, 17, 12, 0, 0, time.UTC),
					Project:   "Flow",
					Tags:      []string{"start-usecase"},
				},
			},
			want: "Sessions Report\n\nMon, 15 Apr 2024 - 1h0m0s\n    3 16:12:00 to 17:12:00 1h0m0s Flow [start-usecase]",
		},
		{
			name:  "Invalid since flag",
			args:  []string{"--since", "2024-04-15T"},
			error: errors.New("2024-04-15T is not a valid time format"),
		},
		{
			name: "Until flag",
			args: []string{"--until", "2024-04-15"},
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
					EndTime:   time.Date(2024, time.April, 14, 15, 12, 0, 0, time.UTC),
					Project:   "Flow",
					Tags:      []string{"start-usecase"},
				},
				{
					Id:        "3",
					StartTime: time.Date(2024, time.April, 15, 16, 12, 0, 0, time.UTC),
					EndTime:   time.Date(2024, time.April, 15, 17, 12, 0, 0, time.UTC),
					Project:   "Flow",
					Tags:      []string{"start-usecase"},
				},
			},
			want: "Sessions Report\n\nSun, 14 Apr 2024 - 3h58m0s\n    1 10:12:00 to 13:10:00 2h58m0s MyTodo [add-todo]\n    2 14:12:00 to 15:12:00 1h0m0s Flow [start-usecase]",
		},
		{
			name:  "Invalid until flag",
			args:  []string{"--until", "224-04-15"},
			error: errors.New("224-04-15 is not a valid time format"),
		},
		{
			name: "Since and Until flag",
			args: []string{"--since", "2024-04-14", "--until", "2024-04-16"},
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
					EndTime:   time.Date(2024, time.April, 14, 15, 12, 0, 0, time.UTC),
					Project:   "Flow",
					Tags:      []string{"start-usecase"},
				},
				{
					Id:        "3",
					StartTime: time.Date(2024, time.April, 15, 16, 12, 0, 0, time.UTC),
					EndTime:   time.Date(2024, time.April, 15, 17, 12, 0, 0, time.UTC),
					Project:   "Flow",
					Tags:      []string{"start-usecase"},
				},
				{
					Id:        "4",
					StartTime: time.Date(2024, time.April, 16, 16, 12, 0, 0, time.UTC),
					EndTime:   time.Date(2024, time.April, 16, 17, 12, 0, 0, time.UTC),
					Project:   "Flow",
					Tags:      []string{"start-usecase"},
				},
			},
			want: "Sessions Report\n\nSun, 14 Apr 2024 - 3h58m0s\n    1 10:12:00 to 13:10:00 2h58m0s MyTodo [add-todo]\n    2 14:12:00 to 15:12:00 1h0m0s Flow [start-usecase]\n\nMon, 15 Apr 2024 - 1h0m0s\n    3 16:12:00 to 17:12:00 1h0m0s Flow [start-usecase]",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			sessionRepository.Sessions = tc.givenSessions
			dateProvider.Now = tc.givenNow
			c := report.Command(app)

			got, err := test.ExecuteCmd(t, c, tc.args...)

			is.Equal(tc.error, err)

			if tc.error == nil {
				is.Equal(tc.want, got)
			}
		})
	}
}
