package start

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	app "github.com/TristanSch1/flow/internal/application/usecases"
	startsession "github.com/TristanSch1/flow/internal/application/usecases/flowsession/start"
	"github.com/TristanSch1/flow/utils"
	"github.com/spf13/cobra"
)

func isTag(arg string) bool {
	return strings.HasPrefix(arg, "+")
}

func Command(app *app.App) *cobra.Command {
	return &cobra.Command{
		Use:                   "start [project] [+tag1 +tag2...]",
		Example:               "start my-todo +add-todo +update-todo",
		Short:                 "Start flow session",
		DisableFlagsInUseLine: true,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return nil
			}

			if len(args) > 0 && strings.HasPrefix(args[0], "+") {
				return errors.New("the first argument must be the project name")
			}

			for _, arg := range args[1:] {
				if !isTag(arg) {
					return fmt.Errorf("invalid tag %v (must start with '+')", arg)
				}
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			logger := log.New(cmd.OutOrStdout(), "", 0)
			// no args -> show list of existing projects
			if len(args) == 0 {
				projects, err := app.ListProjectsUseCase.Execute()
				if err != nil {
					return err
				}
				msg := "Please provide a project name"

				if len(projects) > 0 {
					msg += ", existing projects: "

					for i, project := range projects {
						msg += utils.ProjectColor(project)
						if i < len(projects)-1 {
							msg += ", "
						}
					}
				}

				logger.Println(msg)
				return nil
			}

			tags := []string{}

			for _, tag := range args[1:] {
				tagWithoutPrefix, _ := strings.CutPrefix(tag, "+")
				tags = append(tags, tagWithoutPrefix)
			}
			command := startsession.Command{
				Project: args[0],
				Tags:    tags,
			}

			err := app.StartFlowSessionUseCase.Execute(command)
			if err != nil {
				return err
			}

			text := fmt.Sprintf("Starting flow session for the project %v", utils.ProjectColor(command.Project))

			if len(command.Tags) > 0 {
				text += fmt.Sprintf(" [%v]", utils.TagColor(strings.Join(command.Tags, ", ")))
			}

			text += fmt.Sprintf(" at %v", utils.TimeColor(app.DateProvider.GetNow().Format(time.Kitchen)))

			logger.Println(text)

			return nil
		},
	}
}

func Execute(cmd *cobra.Command) error {
	return cmd.Execute()
}
