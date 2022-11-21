package main

import (
	"os"

	"opsAgent/cmd/opsAgent/app"
)

func main() {
	if err := app.AgentCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
