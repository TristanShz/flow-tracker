package report

import (
	"errors"
	"fmt"
	"log"
	"time"

	app "github.com/TristanSch1/flow/internal/application/usecases"
	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/viewsessionsreport"
	"github.com/TristanSch1/flow/internal/domain/sessionsreport"
	"github.com/TristanSch1/flow/internal/infra/presenter"
	"github.com/TristanSch1/flow/pkg/timerange"
	"github.com/spf13/cobra"
)

func isFormatFlagValid(flag string) bool {
	return flag == sessionsreport.FormatByDay || flag == sessionsreport.FormatByProject
}

func parseTimeFlag(flag string) (time.Time, error) {
	parsedTime, err := time.Parse("2006-01-02", flag)

	if err == nil {
		return parsedTime, nil
	}

	return time.Time{}, fmt.Errorf("%v is not a valid time format", flag)
}

func parseSinceFlag(cmd *cobra.Command) (time.Time, error) {
	sinceFlag, _ := cmd.Flags().GetString("since")
	if sinceFlag != "" {
		since, err := parseTimeFlag(sinceFlag)
		if err != nil {
			return time.Time{}, err
		}
		return since, nil
	}
	return time.Time{}, nil
}

func parseUntilFlag(cmd *cobra.Command) (time.Time, error) {
	untilFlag, _ := cmd.Flags().GetString("until")
	fmt.Println(untilFlag)
	if untilFlag != "" {
		until, err := parseTimeFlag(untilFlag)
		if err != nil {
			return time.Time{}, err
		}
		return until, nil
	}
	return time.Time{}, nil
}

func Command(app *app.App) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "report",
		Short: "Report",
		RunE: func(cmd *cobra.Command, args []string) error {
			logger := log.New(cmd.OutOrStdout(), "", 0)

			presenter := presenter.SessionsReportCLIPresenter{Logger: logger}

			formatFlag, _ := cmd.Flags().GetString("format")

			if formatFlag != "" && !isFormatFlagValid(formatFlag) {
				return errors.New("invalid format flag. possible values: by-day, by-project")
			}

			projectFlag, _ := cmd.Flags().GetString("project")
			command := viewsessionsreport.Command{
				Project: projectFlag,
				Format:  formatFlag,
			}

			dayFlag, _ := cmd.Flags().GetBool("day")
			if dayFlag {
				timeRange := timerange.NewDayTimeRange(app.DateProvider.GetNow())

				command.Since = timeRange.Since
				command.Until = timeRange.Until
			}

			weekFlag, _ := cmd.Flags().GetBool("week")
			if weekFlag {
				timeRange := timerange.NewWeekTimeRange(app.DateProvider.GetNow())

				command.Since = timeRange.Since
				command.Until = timeRange.Until
			}

			sinceFlag, sinceFlagErr := parseSinceFlag(cmd)
			if sinceFlagErr != nil {
				return sinceFlagErr
			}

			if !sinceFlag.IsZero() {
				command.Since = sinceFlag
			}

			untilFlag, untilFlagErr := parseUntilFlag(cmd)
			if untilFlagErr != nil {
				return untilFlagErr
			}

			if !untilFlag.IsZero() {
				command.Until = untilFlag
			}

			fmt.Println(command)
			err := app.ViewSessionsReportUseCase.Execute(command, presenter)
			if err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().StringP("project", "p", "", "get a report for all flow sessions of given project")
	cmd.Flags().StringP("format", "f", "", "Specify the format of the report. Possible values: by-day, by-project, total-duration")
	cmd.Flags().StringP("since", "s", "", "Specify the start date of the report")
	cmd.Flags().StringP("until", "u", "", "Specify the end date of the report")
	cmd.Flags().BoolP("day", "d", false, "Get a report for all flow sessions of the day")
	cmd.Flags().BoolP("week", "w", false, "Get a report for all flow sessions of the week")

	return cmd
}

func Execute(cmd *cobra.Command) error {
	return cmd.Execute()
}
