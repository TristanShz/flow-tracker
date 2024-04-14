package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	app "github.com/TristanSch1/flow/internal/application/usecases"
	"github.com/TristanSch1/flow/internal/application/usecases/flowsession/start"
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
				fmt.Printf("%v", strings.Join(projects, ", "))
				os.Exit(0)
			}
			command := start.Command{
				Project: args[0],
				Tags:    args[1:],
			}

			err := app.StartFlowSessionUseCase.Execute(command)
			if err != nil {
				fmt.Printf("%v", err)
				os.Exit(1)
			}

			fmt.Printf("Starting flow session for the project %v at %v", command.Project, time.Now().Format(time.Kitchen))
		},
	}
}
