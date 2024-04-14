package cmd

import (
	"fmt"
	"os"

	app "github.com/TristanSch1/flow/internal/application/usecases"
	"github.com/spf13/cobra"
)

func stopCmd(app *app.App) *cobra.Command {
	return &cobra.Command{
		Use:                   "stop",
		Short:                 "Stop flow session",
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, _ []string) {
			err := app.StopFlowSessionUseCase.Execute()
			if err != nil {
				fmt.Printf("%v", err)
				os.Exit(1)
			}

			fmt.Printf("Flow session stopped")
		},
	}
}
