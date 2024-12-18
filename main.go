package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
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

	cmd := exec.Command(input)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error executing command: %v", err)
	}

	return nil
}