package cmd

import (
	"fmt"
	"os"

	app "github.com/TristanSch1/flow/internal/application/usecases"
	"github.com/spf13/cobra"
)

func reportCmd(app *app.App) *cobra.Command {
	return &cobra.Command{
		Use:   "report",
		Short: "Report",
		Run: func(cmd *cobra.Command, args []string) {
			report, err := app.AllSessionsReportUseCase.Execute()
			if err != nil {
				fmt.Printf("%v", err)
				os.Exit(1)
			}

			fmt.Println(report.PrettyPrint())
		},
	}
}
