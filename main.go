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
		main()
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

func execCommand(input string) error {
	input = strings.TrimSuffix(input, "\n")
	args := strings.Split(input, " ")

	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return fmt.Errorf("cd: missing argument")
		}
		return os.Chdir(args[1])
	case "up":
		println("up")
		os.Exit(0)
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
	return history.commands[len(history.commands)-1]
}
