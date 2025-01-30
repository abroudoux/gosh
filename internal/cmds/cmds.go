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
	case "help":
		printHelp()
		return nil
	case "exit", "quit":
		os.Exit(0)
	}

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}


func printHelp() {
	fmt.Println("cd <directory>: change working directory")
	fmt.Println("history: print command history")
	fmt.Println("help: print this help message")
	fmt.Println("exit or quit: exit the shell")
}