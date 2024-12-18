package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	printPwd()
	fmt.Print("> ")
	input, err := readInput()
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		err := execCommand(input)
		if err != nil {
			fmt.Println(err)
		}

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