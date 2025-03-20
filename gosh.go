package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/charmbracelet/log"
)

func main() {
	StartGosh()
}

func StartGosh() {
	err := PrintUi()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	input, err := ReadInput()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	for {
		err := ExecCommand(input)
		log.Error(err)
		StartGosh()
	}
}

func PrintPwd() error {
	cmd := exec.Command("pwd")
	pwd, err := cmd.Output()
	if err != nil {
		return err
	}

	fmt.Print(string(pwd))
	return nil
}

func PrintUi() error {
	err := PrintPwd()
	if err != nil {
		return err
	}

	fmt.Print("> ")
	return nil
}

func ReadInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return string(input), nil
}

func ExecCommand(input string) error {
	input = strings.TrimSuffix(input, "\n")
	args := strings.Split(input, " ")

	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return fmt.Errorf("cd: missing argument")
		}
		return os.Chdir(args[1])
	case "help":
		PrintHelpMessage()
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

func PrintHelpMessage() {
	fmt.Println("cd <directory>: change working directory")
	fmt.Println("help: print this help message")
	fmt.Println("exit or quit: exit the shell")
}
