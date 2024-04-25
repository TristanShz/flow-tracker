package test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/TristanSch1/flow/internal/application"
	app "github.com/TristanSch1/flow/internal/application/usecases"
	startsession "github.com/TristanSch1/flow/internal/application/usecases/flowsession/start"
	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/status"
	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/stop"
	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/viewsessionsreport"
	"github.com/TristanSch1/flow/internal/application/usecases/project/list"
	"github.com/TristanSch1/flow/internal/infra"
	"github.com/spf13/cobra"
)

func ExecuteCmd(t *testing.T, c *cobra.Command, args ...string) (string, error) {
	t.Helper()

	buf := new(bytes.Buffer)
	c.SetOut(buf)
	c.SetErr(buf)
	c.SetArgs(args)

	err := c.Execute()
	return strings.TrimSpace(buf.String()), err
}

func InitializeApp(sessionRepository application.SessionRepository) *app.App {
	dateProvider := &infra.StubDateProvider{}
	idProvider := &infra.StubIDProvider{}

	startFlowSessionUseCase := startsession.NewStartFlowSessionUseCase(sessionRepository, dateProvider, idProvider)
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
