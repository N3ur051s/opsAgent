package main

import (
	"os"

	"simpleagent/cmd/simpleagent/app"
)

func main() {
	if err := app.AgentCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
