package stop

import (
	"log"

	app "github.com/TristanSch1/flow/internal/application/usecases"
	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/stopsession"
	"github.com/TristanSch1/flow/utils"
	"github.com/spf13/cobra"
)

func Command(app *app.App) *cobra.Command {
	return &cobra.Command{
		Use:                   "stop",
		Short:                 "Stop flow session",
		DisableFlagsInUseLine: true,
		RunE: func(cmd *cobra.Command, _ []string) error {
			logger := log.New(cmd.OutOrStdout(), "", 0)
			duration, err := app.StopFlowSessionUseCase.Execute()
			if err != nil {
				if err == stopsession.ErrNoCurrentSession {
					logger.Println("No flow session to stop.")
					return nil
				}
				return err
			}

			logger.Printf("Flow session stopped, you were in the flow for %v", utils.TimeColor(duration.String()))
			return nil
		},
	}
}
