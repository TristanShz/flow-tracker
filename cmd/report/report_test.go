package report_test

import (
	"bytes"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/TristanSch1/flow/cmd/report"
	"github.com/TristanSch1/flow/internal/application"
	app "github.com/TristanSch1/flow/internal/application/usecases"
	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/start"
	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/status"
	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/stop"
	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/viewsessionsreport"
	"github.com/TristanSch1/flow/internal/application/usecases/project/list"
	"github.com/TristanSch1/flow/internal/domain/session"
	"github.com/TristanSch1/flow/internal/infra"
	is "github.com/matryer/is"
	"github.com/spf13/cobra"
)

func execute(t *testing.T, c *cobra.Command, args ...string) (string, error) {
	t.Helper()

	buf := new(bytes.Buffer)
	c.SetOut(buf)
	c.SetErr(buf)
	c.SetArgs(args)

	err := c.Execute()
	return strings.TrimSpace(buf.String()), err
}

func initializeApp(sessionRepository application.SessionRepository) *app.App {
	dateProvider := &infra.StubDateProvider{}
	idProvider := &infra.StubIDProvider{}

	startFlowSessionUseCase := start.NewStartFlowSessionUseCase(sessionRepository, dateProvider, idProvider)
	stopFlowSessionUseCase := stop.NewStopSessionUseCase(sessionRepository, dateProvider)
	flowSessionStatusUseCase := status.NewFlowSessionStatusUseCase(sessionRepository, dateProvider)

	viewSessionsReportUseCase := viewsessionsreport.NewViewSessionsReportUseCase(sessionRepository)

	listProjectsUseCase := list.NewListProjectsUseCase(sessionRepository)

	return app.NewApp(
		startFlowSessionUseCase,
		stopFlowSessionUseCase,
		flowSessionStatusUseCase,
		listProjectsUseCase,
		viewSessionsReportUseCase,
	)
}

func TestReportCommand(t *testing.T) {
	is := is.New(t)

	sessionRepository := &infra.InMemorySessionRepository{}
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
	app := initializeApp(sessionRepository)

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

			got, err := execute(t, c, tc.args...)

			is.Equal(tc.error, err)

			if tc.error == nil {
				is.Equal(tc.want, got)
			}
		})
	}
}
