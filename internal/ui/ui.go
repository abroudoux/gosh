package ui

import (
	"fmt"
	"os/exec"
)

func printPwd() error {
	cmd := exec.Command("pwd")
	pwd, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("error executing command: %v", err)
	}

	fmt.Print(string(pwd))
	return nil
}

func PrintUi() error {
	err := printPwd()
	if err != nil {
		return err
	}

	fmt.Print("> ")
	return nil
}