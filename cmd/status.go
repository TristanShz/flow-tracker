package cmd

import (
	"fmt"
	"os"

	app "github.com/TristanSch1/flow/internal/application/usecases"
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

			fmt.Printf("Current session :\n %v \n", status.Session.PrettyString())
			fmt.Printf("%v", status.StatusText)
		},
	}
}
