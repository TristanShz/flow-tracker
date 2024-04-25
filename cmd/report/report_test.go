package report_test

import (
	"errors"
	"testing"
	"time"

	"github.com/TristanSch1/flow/cmd/report"
	"github.com/TristanSch1/flow/internal/domain/session"
	"github.com/TristanSch1/flow/internal/infra"
	"github.com/TristanSch1/flow/test"
	is "github.com/matryer/is"
)

func TestReportCommand(t *testing.T) {
	is := is.New(t)

	sessionRepository := &infra.InMemorySessionRepository{}
	dateProvider := infra.NewStubDateProvider()
	sessionRepository.Sessions = []session.Session{
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
	}
	app := test.InitializeApp(sessionRepository, dateProvider)

	tt := []struct {
		error error
		name  string
		want  string
		args  []string
	}{
		{
			name:  "No args",
			args:  []string{"report"},
			want:  "Sessions Report\n\nSun, 14 Apr 2024 :\n    From 10:12:00 to 13:10:00 2h58m0s MyTodo [add-todo]\n    From 14:12:00 to 15:12:00 1h0m0s Flow [start-usecase]",
			error: nil,
		},
		{
			name:  "Invalid format flag",
			args:  []string{"report", "--format", "invalid"},
			want:  "",
			error: errors.New("invalid format flag. possible values: by-day, by-project"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			c := report.ReportCmd(app)

			got, err := test.ExecuteCmd(t, c, tc.args...)

			is.Equal(tc.error, err)

			if tc.error == nil {
				is.Equal(tc.want, got)
			}
		})
	}
}
