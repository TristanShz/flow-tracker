package status

import (
	"fmt"
	"log"
	"strings"

	app "github.com/TristanShz/flow/internal/application/usecases"
	"github.com/TristanShz/flow/internal/application/usecases/flowsession/sessionstatus"
	"github.com/TristanShz/flow/utils"
	"github.com/spf13/cobra"
)

func Command(app *app.App) *cobra.Command {
	return &cobra.Command{
		Use:                   "status",
		Short:                 "Show the current flow session status",
		DisableFlagsInUseLine: true,
		RunE: func(cmd *cobra.Command, _ []string) error {
			logger := log.New(cmd.OutOrStdout(), "", 0)
			status, err := app.FlowSessionStatusUseCase.Execute()
			if err != nil {
				if err == sessionstatus.ErrNoCurrentSession {
					logger.Printf("No active flow session")
					return nil
				}
				return err
			}

			msg := fmt.Sprintf(
				"You're in the flow for %v on project %v",
				utils.TimeColor(status.Duration.String()),
				utils.ProjectColor(status.Session.Project),
			)

			if len(status.Session.Tags) > 0 {
				msg += fmt.Sprintf(" with tags: %v", utils.TagColor(strings.Join(status.Session.Tags, ", ")))
			}

			logger.Println(msg)

			return nil
		},
	}
}
