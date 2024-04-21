package cmd

import (
	"fmt"
	"os"
	"strings"

	app "github.com/TristanSch1/flow/internal/application/usecases"
	"github.com/TristanSch1/flow/utils"
	"github.com/spf13/cobra"
)

func statusCmd(app *app.App) *cobra.Command {
	return &cobra.Command{
		Use:                   "status",
		Short:                 "Show the current flow session status",
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, _ []string) {
			status, err := app.FlowSessionStatusUseCase.Execute()
			if err != nil {
				fmt.Printf("%v", err)
				os.Exit(1)
			}

			fmt.Println("Current session:")
			fmt.Printf("  Start at: %v\n", utils.TimeColor(status.Session.StartTime.Format("2006-01-02 15:04:05")))
			fmt.Printf("  Project: %v\n", utils.ProjectColor(status.Session.Project))
			fmt.Printf("  Tags: [%v]\n", utils.TagColor(strings.Join(status.Session.Tags, ", ")))
			fmt.Printf("You're in the flow for %v", status.Duration.String())
		},
	}
}
