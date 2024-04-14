package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	app "github.com/TristanSch1/flow/internal/application/usecases"
	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/start"
	"github.com/TristanSch1/flow/utils"
	"github.com/spf13/cobra"
)

func isTag(arg string) bool {
	return strings.HasPrefix(arg, "+")
}

func startCmd(app *app.App) *cobra.Command {
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
		Run: func(cmd *cobra.Command, args []string) {
			// no args -> show list of existing projects
			if len(args) == 0 {
				projects, err := app.ListProjectsUseCase.Execute()
				if err != nil {
					fmt.Printf("%v", err)
					os.Exit(1)
				}
				fmt.Printf("Existing projects : %v", strings.Join(projects, ", "))
				os.Exit(0)
			}

			tags := []string{}

			for _, tag := range args[1:] {
				tagWithoutPrefix, _ := strings.CutPrefix(tag, "+")
				tags = append(tags, tagWithoutPrefix)
			}
			command := start.Command{
				Project: args[0],
				Tags:    tags,
			}

			err := app.StartFlowSessionUseCase.Execute(command)
			if err != nil {
				fmt.Printf("%v", err)
				os.Exit(1)
			}

			text := fmt.Sprintf("Starting flow session for the project %v", utils.PurpleText(command.Project))

			if len(command.Tags) > 0 {
				text += fmt.Sprintf(" [%v]", utils.YellowText(strings.Join(command.Tags, ", ")))
			}

			text += fmt.Sprintf(" at %v", utils.GreenText(time.Now().Format(time.Kitchen)))

			fmt.Println(text)
		},
	}
}
