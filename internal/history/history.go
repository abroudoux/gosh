package history

import (
	"fmt"
)

type History struct {
	commands []string
}

func InitHistory() History {
	return History{}
}

func (history *History) AddCommand(command string) {
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

func (history *History) PrintHistory() {
	if (len(history.commands) == 0) {
		fmt.Println("No commands in history")
		return
	}

	for _, command := range history.commands {
		fmt.Println(command)
	}
}