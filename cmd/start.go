package cmd

import (
	"errors"
	"fmt"

	app "github.com/TristanSch1/flow/internal/application/usecases"
	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/start"
	"github.com/spf13/cobra"
)

func startCmd(app *app.App) *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "Start flow session",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("requires at least one arg")
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			command := start.Command{
				Project: args[0],
			}

			err := app.StartFlowSessionUseCase.Execute(command)
			if err != nil {
				fmt.Printf("%v", err)
			}
		},
	}
}
