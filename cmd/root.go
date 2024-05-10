package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/TristanShz/flow/cmd/abort"
	"github.com/TristanShz/flow/cmd/edit"
	"github.com/TristanShz/flow/cmd/report"
	"github.com/TristanShz/flow/cmd/start"
	"github.com/TristanShz/flow/cmd/status"
	"github.com/TristanShz/flow/cmd/stop"
	app "github.com/TristanShz/flow/internal/application/usecases"
	abortsession "github.com/TristanShz/flow/internal/application/usecases/flowsession/abort"
	"github.com/TristanShz/flow/internal/application/usecases/flowsession/sessionstatus"
	startsession "github.com/TristanShz/flow/internal/application/usecases/flowsession/start"
	"github.com/TristanShz/flow/internal/application/usecases/flowsession/stopsession"
	"github.com/TristanShz/flow/internal/application/usecases/flowsession/viewsessionsreport"
	"github.com/TristanShz/flow/internal/application/usecases/project/list"
	"github.com/TristanShz/flow/internal/infra"
	"github.com/TristanShz/flow/internal/infra/filesystem"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "flow",
	Short: "Flow is a tool to manage your time tracking",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, world!")
	},
}

func initializeApp(path string) *app.App {
	sessionRepository := filesystem.NewFileSystemSessionRepository(path)

	dateProvider := &infra.RealDateProvider{}
	idProvider := &infra.RealIDProvider{}

	startFlowSessionUseCase := startsession.NewStartFlowSessionUseCase(&sessionRepository, dateProvider, idProvider)
	stopFlowSessionUseCase := stopsession.NewStopSessionUseCase(&sessionRepository, dateProvider)
	abortFlowSessionUseCase := abortsession.NewAbortFlowSessionUseCase(&sessionRepository)
	flowSessionStatusUseCase := sessionstatus.NewFlowSessionStatusUseCase(&sessionRepository, dateProvider)

	viewSessionsReportUseCase := viewsessionsreport.NewViewSessionsReportUseCase(&sessionRepository)

	listProjectsUseCase := list.NewListProjectsUseCase(&sessionRepository)

	return app.NewApp(
		&sessionRepository,
		dateProvider,
		startFlowSessionUseCase,
		stopFlowSessionUseCase,
		abortFlowSessionUseCase,
		flowSessionStatusUseCase,
		listProjectsUseCase,
		viewSessionsReportUseCase,
	)
}

func Execute() {
	homePath, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	sessionsPath := filepath.Join(homePath, ".flow")

	app := initializeApp(sessionsPath)

	rootCmd.AddCommand(start.Command(app))
	rootCmd.AddCommand(stop.Command(app))
	rootCmd.AddCommand(status.Command(app))
	rootCmd.AddCommand(report.Command(app))
	rootCmd.AddCommand(edit.Command(app, sessionsPath))
	rootCmd.AddCommand(abort.Command(app))

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
