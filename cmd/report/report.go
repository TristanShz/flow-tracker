package report

import (
	"errors"
	"log"
	"time"

	app "github.com/TristanSch1/flow/internal/application/usecases"
	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/viewsessionsreport"
	"github.com/TristanSch1/flow/internal/domain/sessionsreport"
	"github.com/TristanSch1/flow/internal/infra/presenter"
	"github.com/TristanSch1/flow/pkg/timerange"
	"github.com/spf13/cobra"
)

func checkFormatFlag(flag string) bool {
	return flag == sessionsreport.FormatByDay || flag == sessionsreport.FormatByProject
}

func Command(app *app.App) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "report",
		Short: "Report",
		RunE: func(cmd *cobra.Command, args []string) error {
			logger := log.New(cmd.OutOrStdout(), "", 0)

			projectFlag, _ := cmd.Flags().GetString("project")
			formatFlag, _ := cmd.Flags().GetString("format")
			dayFlag, _ := cmd.Flags().GetBool("day")
			weekFlag, _ := cmd.Flags().GetBool("week")

			if formatFlag != "" && !checkFormatFlag(formatFlag) {
				return errors.New("invalid format flag. possible values: by-day, by-project")
			}

			presenter := presenter.SessionsReportCLIPresenter{Logger: logger}

			command := viewsessionsreport.Command{
				Project: projectFlag,
				Format:  formatFlag,
			}

			if dayFlag {
				timeRange := timerange.NewDayTimeRange(time.Now())

				command.Since = timeRange.Since
				command.Until = timeRange.Until
			}

			if weekFlag {
				timeRange := timerange.NewWeekTimeRange(time.Now())

				command.Since = timeRange.Since
				command.Until = timeRange.Until
			}

			err := app.ViewSessionsReportUseCase.Execute(command, presenter)
			if err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().StringP("project", "p", "", "get a report for all flow sessions of given project")
	cmd.Flags().StringP("format", "f", "by-day", "Specify the format of the report. Possible values: by-day, by-project, total-duration")
	cmd.Flags().StringP("since", "s", "", "Specify the start date of the report")
	cmd.Flags().StringP("until", "u", "", "Specify the end date of the report")
	cmd.Flags().BoolP("day", "d", false, "Get a report for all flow sessions of the day")
	cmd.Flags().BoolP("week", "w", false, "Get a report for all flow sessions of the week")

	return cmd
}

func Execute(cmd *cobra.Command) error {
	return cmd.Execute()
}
