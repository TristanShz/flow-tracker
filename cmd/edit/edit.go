package edit

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	app "github.com/TristanSch1/flow/internal/application/usecases"
	"github.com/TristanSch1/flow/internal/infra/filesystem"
	"github.com/TristanSch1/flow/utils"
	"github.com/spf13/cobra"
)

func getOpenCommand(filePath string) *exec.Cmd {
	var command *exec.Cmd
	switch os := runtime.GOOS; os {
	case "windows":
		command = exec.Command("notepad", filePath)
	case "darwin":
		command = exec.Command("nano", filePath)
	case "linux":
		command = exec.Command("nano", filePath)
	default:
		fmt.Printf("Unsupported OS: %v\n", os)
		return nil
	}

	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	return command
}

func Command(app *app.App, sessionsPath string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit",
		Short: "Open the flow session in the default editor",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("missing session ID")
			}

			if len(args) == 1 {
				if utils.IsIDValid(args[0]) {
					return nil
				} else {
					return fmt.Errorf("invalid ID %v", args[0])
				}
			}

			return fmt.Errorf("too many arguments")
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			logger := log.New(cmd.OutOrStdout(), "", 0)

			session := app.SessionRepository.FindById(args[0])

			if session == nil {
				logger.Println("Session not found")
				return nil
			}

			sessionFilename := filesystem.SessionFilename{
				Id:        session.Id,
				Project:   session.Project,
				StartTime: session.StartTime,
			}

			filePath := filepath.Join(sessionsPath, sessionFilename.String())

			command := getOpenCommand(filePath)

			err := command.Run()
			if err != nil {
				fmt.Printf("Error whilte opening the file: %v\n", err)
				return nil
			}

			return nil
		},
	}

	return cmd
}
