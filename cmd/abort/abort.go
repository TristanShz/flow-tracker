package abort

import (
	"log"

	app "github.com/TristanShz/flow/internal/application/usecases"
	"github.com/spf13/cobra"
)

func Command(app *app.App) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "abort",
		Short: "Abort the current session",
		RunE: func(cmd *cobra.Command, args []string) error {
			logger := log.New(cmd.OutOrStdout(), "", 0)
			err := app.AbortFlowSessionUseCase.Execute()
			if err != nil {
				logger.Println(err)
				return nil
			}

			logger.Println("Session aborted")

			return nil
		},
	}

	return cmd
}
