package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/TristanSch1/flow/cmd/report"
	"github.com/TristanSch1/flow/cmd/start"
	app "github.com/TristanSch1/flow/internal/application/usecases"
	startsession "github.com/TristanSch1/flow/internal/application/usecases/flowsession/start"
	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/status"
	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/stop"
	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/viewsessionsreport"
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
		log.Fatal(err)
	}

	sessionRepository := filesystem.NewFileSystemSessionRepository(homePath)
	if err != nil {
		log.Fatal(err)
	}
	dateProvider := &infra.RealDateProvider{}
	idProvider := &infra.RealIDProvider{}

	startFlowSessionUseCase := startsession.NewStartFlowSessionUseCase(&sessionRepository, dateProvider, idProvider)
	stopFlowSessionUseCase := stop.NewStopSessionUseCase(&sessionRepository, dateProvider)
	flowSessionStatusUseCase := status.NewFlowSessionStatusUseCase(&sessionRepository, dateProvider)

	viewSessionsReportUseCase := viewsessionsreport.NewViewSessionsReportUseCase(&sessionRepository)

	listProjectsUseCase := list.NewListProjectsUseCase(&sessionRepository)

	return app.NewApp(
		startFlowSessionUseCase,
		stopFlowSessionUseCase,
		flowSessionStatusUseCase,
		listProjectsUseCase,
		viewSessionsReportUseCase,
	)
}

func Execute() {
	app := initializeApp()

	rootCmd.AddCommand(start.Command(app))
	rootCmd.AddCommand(stopCmd(app))
	rootCmd.AddCommand(statusCmd(app))
	rootCmd.AddCommand(report.ReportCmd(app))

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
