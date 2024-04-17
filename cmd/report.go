package cmd

import (
	"fmt"
	"os"

	app "github.com/TristanSch1/flow/internal/application/usecases"
	"github.com/spf13/cobra"
)

func reportCmd(app *app.App) *cobra.Command {
	command := &cobra.Command{
		Use:   "report",
		Short: "Report",
		Run: func(cmd *cobra.Command, args []string) {
			projectFlag, _ := cmd.Flags().GetString("project")

			if projectFlag != "" {
				report, err := app.ViewSessionsReportUseCase.Execute()
				if err != nil {
					fmt.Printf("%v", err)
					os.Exit(1)
				}

				fmt.Println(report.TotalDuration())
			} else {
				report, err := app.ViewSessionsReportUseCase.Execute()
				if err != nil {
					fmt.Printf("%v", err)
					os.Exit(1)
				}

				fmt.Println(report.TotalDuration())
			}
		},
	}

	command.PersistentFlags().String("project", "", "get a report for all flow sessions of given project")

	return command
}
