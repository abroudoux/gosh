package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	history := initHistory()

	startShell(history)
}

func startShell(history History) {
	err := printUi()
	if err != nil {
		fmt.Println(err)
		return
	}

	input, err := readInput()
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		err := execCommand(input, history)
		if err != nil {
			fmt.Println(err)
		}

		history.addCommand(input)
		startShell(history)
	}
}

func readInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("error reading input: %v", err)
	}

	return string(input), nil
}

func execCommand(input string, history History) error {
	input = strings.TrimSuffix(input, "\n")
	args := strings.Split(input, " ")

	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return fmt.Errorf("cd: missing argument")
		}
		return os.Chdir(args[1])
	case "history":
		history.printHistory()
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

func printPwd() error {
	cmd := exec.Command("pwd")
	pwd, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("error executing command: %v", err)
	}

	fmt.Print(string(pwd))
	return nil
}

func printUi() error {
	err := printPwd()
	if err != nil {
		return err
	}

	fmt.Print("> ")
	return nil
}

type History struct {
	commands []string
}

func initHistory() History {
	history := History{}
	return history
}

func (history *History) addCommand(command string) {
	lastCommand := history.getLastCommand()

	if (command != lastCommand) {	
		history.commands = append(history.commands, command)
	}
}

func (history *History) getLastCommand() string {
	if (len(history.commands) == 0) {
		return ""
	}

	return history.commands[len(history.commands)-1]
}

func (history *History) printHistory() {
	if (len(history.commands) == 0) {
		fmt.Println("No commands in history")
		return
	}

	for _, command := range history.commands {
		fmt.Println(command)
	}
}