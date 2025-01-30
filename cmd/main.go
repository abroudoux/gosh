package main

import (
	"github.com/abroudoux/gosh/internal/history"
	"github.com/abroudoux/gosh/internal/shell"
)

func main() {
	history := history.InitHistory()
	shell.StartShell(history)
}