package shell

import (
	"bufio"
	"fmt"
	"os"

	"github.com/abroudoux/gosh/internal/cmds"
	"github.com/abroudoux/gosh/internal/history"
	"github.com/abroudoux/gosh/internal/ui"
)

type History = history.History

func StartShell(history History) {
	err := ui.PrintUi()
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
		err := cmds.ExecCommand(input, history)
		if err != nil {
			fmt.Println(err)
		}

		history.AddCommand(input)
		StartShell(history)
	}
}

func readInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return string(input), nil
}