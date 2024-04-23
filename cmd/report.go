package cmd

import (
	"fmt"
	"os"
	"time"

	app "github.com/TristanSch1/flow/internal/application/usecases"
	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/viewsessionsreport"
	"github.com/TristanSch1/flow/internal/domain/sessionsreport"
	"github.com/TristanSch1/flow/internal/infra/presenter"
	"github.com/spf13/cobra"
)

func checkFormatFlag(flag string) bool {
	return flag == sessionsreport.FormatByDay || flag == sessionsreport.FormatByProject
}

func reportCmd(app *app.App) *cobra.Command {
	command := &cobra.Command{
		Use:   "report",
		Short: "Report",
		Run: func(cmd *cobra.Command, args []string) {
			projectFlag, _ := cmd.Flags().GetString("project")
			formatFlag, _ := cmd.Flags().GetString("format")
			dayFlag, _ := cmd.Flags().GetBool("day")
			weekFlag, _ := cmd.Flags().GetBool("week")

			if formatFlag != "" && !checkFormatFlag(formatFlag) {
				fmt.Printf("Invalid format flag. Possible values: by-day, by-project")
				os.Exit(1)
			}

			presenter := presenter.SessionsReportCLIPresenter{}

			command := viewsessionsreport.Command{
				Project: projectFlag,
				Format:  formatFlag,
			}

			if dayFlag {
				now := time.Now()

				command.Since = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
				command.Until = time.Now()
			}

			if weekFlag {
				now := time.Now()
				weekStart := now.AddDate(0, 0, -int(now.Weekday()))
				weekEnd := weekStart.AddDate(0, 0, 6)
				command.Since = weekStart
				command.Until = weekEnd
			}

			err := app.ViewSessionsReportUseCase.Execute(command, presenter)
			if err != nil {
				fmt.Printf("%v", err)
				os.Exit(1)
			}
		},
	}

	command.Flags().StringP("project", "p", "", "get a report for all flow sessions of given project")
	command.Flags().StringP("format", "f", "by-day", "Specify the format of the report. Possible values: by-day, by-project, total-duration")
	command.Flags().BoolP("day", "d", false, "Get a report for all flow sessions of the day")
	command.Flags().BoolP("week", "w", false, "Get a report for all flow sessions of the week")

	return command
}
