package status

import (
	"log"
	"strings"

	app "github.com/TristanSch1/flow/internal/application/usecases"
	"github.com/TristanSch1/flow/utils"
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
				return err
			}

			logger.Printf(
				"You're in the flow for %v on project %v with tags: %v",
				utils.TimeColor(status.Duration.String()),
				utils.ProjectColor(status.Session.Project),
				utils.TagColor(strings.Join(status.Session.Tags, ", ")),
			)

			return nil
		},
	}
}
