package edit

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"

	app "github.com/TristanSch1/flow/internal/application/usecases"
	"github.com/TristanSch1/flow/utils"
	"github.com/spf13/cobra"
)

func Command(app *app.App, sessionsPath string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit",
		Short: "Open the flow session in the default editor",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return nil
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

			fileName := strconv.FormatInt(session.StartTime.Unix(), 10) + ".json"
			filePath := filepath.Join(sessionsPath, fileName)

			logger.Printf("Opening %v...", filePath)
			var command *exec.Cmd
			switch os := runtime.GOOS; os {
			case "windows":
				command = exec.Command("notepad", filePath)
			case "darwin":
				command = exec.Command("nano", filePath)
			case "linux":
				command = exec.Command("nano", filePath)
			default:
				fmt.Printf("Impossible d'ouvrir le fichier sur ce système d'exploitation: %s\n", os)
				return nil
			}

			command.Stdin = os.Stdin
			command.Stdout = os.Stdout
			command.Stderr = os.Stderr

			err := command.Run()
			if err != nil {
				fmt.Printf("Erreur lors de l'ouverture du fichier: %v\n", err)
				return nil
			}

			return nil
		},
	}

	return cmd
}
