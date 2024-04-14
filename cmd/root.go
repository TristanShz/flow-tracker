package cmd

import (
	"fmt"
	"os"

	app "github.com/TristanSch1/flow/internal/application/usecases"
	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/allsessionsreport"
	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/projectsessionsreport"
	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/start"
	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/status"
	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/stop"
	"github.com/TristanSch1/flow/internal/application/usecases/project/list"
	"github.com/TristanSch1/flow/internal/infra"
	"github.com/TristanSch1/flow/internal/infra/filesystem"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "flow",
	Short: "Flow is a tool to manage your time tracking",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, world!")
	},
}

func initializeApp() *app.App {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return nil
	}

	sessionRepository := &filesystem.FileSystemSessionRepository{
		FlowFolderPath: homePath,
	}
	dateProvider := &infra.RealDateProvider{}

	startFlowSessionUseCase := start.NewStartFlowSessionUseCase(sessionRepository, dateProvider)
	stopFlowSessionUseCase := stop.NewStopSessionUseCase(sessionRepository, dateProvider)
	flowSessionStatusUseCase := status.NewFlowSessionStatusUseCase(sessionRepository, dateProvider)

	allSessionsReportUseCase := allsessionsreport.NewFlowSessionsReportUseCase(sessionRepository)
	projectSessionsReportUseCase := projectsessionsreport.NewProjectSessionsReportUseCase(sessionRepository)

	listProjectsUseCase := list.NewListProjectsUseCase(sessionRepository)

	return app.NewApp(
		startFlowSessionUseCase,
		stopFlowSessionUseCase,
		flowSessionStatusUseCase,
		listProjectsUseCase,
		allSessionsReportUseCase,
		projectSessionsReportUseCase,
	)
}

func Execute() {
	app := initializeApp()

	rootCmd.AddCommand(startCmd(app))
	rootCmd.AddCommand(stopCmd(app))
	rootCmd.AddCommand(statusCmd(app))
	rootCmd.AddCommand(reportCmd(app))

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
