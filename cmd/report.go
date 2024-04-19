package cmd

import (
	"fmt"
	"os"

	app "github.com/TristanSch1/flow/internal/application/usecases"
	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/viewsessionsreport"
	"github.com/TristanSch1/flow/internal/infra/presenter"
	"github.com/spf13/cobra"
)

func reportCmd(app *app.App) *cobra.Command {
	command := &cobra.Command{
		Use:   "report",
		Short: "Report",
		Run: func(cmd *cobra.Command, args []string) {
			projectFlag, _ := cmd.Flags().GetString("project")
			presenter := presenter.SessionsReportCLIPresenter{}

			if projectFlag != "" {
				err := app.ViewSessionsReportUseCase.Execute(viewsessionsreport.Command{}, presenter)
				if err != nil {
					fmt.Printf("%v", err)
					os.Exit(1)
				}
			} else {
				err := app.ViewSessionsReportUseCase.Execute(viewsessionsreport.Command{}, presenter)
				if err != nil {
					fmt.Printf("%v", err)
					os.Exit(1)
				}
			}
		},
	}

	command.Flags().StringP("project", "p", "", "get a report for all flow sessions of given project")
	command.Flags().StringP("format", "f", "by-day", "Specify the format of the report. Possible values: by-day, by-project, total-duration")

	return command
}
