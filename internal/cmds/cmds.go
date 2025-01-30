package cmds

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/abroudoux/gosh/internal/history"
)

type History = history.History

func ExecCommand(input string, history History) error {
	input = strings.TrimSuffix(input, "\n")
	args := strings.Split(input, " ")

	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return fmt.Errorf("cd: missing argument")
		}
		return os.Chdir(args[1])
	case "history":
		history.PrintHistory()
		return nil
	case "exit":
		os.Exit(0)
	}

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error executing command: %v", err)
	}

	return nil
}