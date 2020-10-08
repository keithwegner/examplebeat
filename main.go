package main

import (
	"os"

	"github.com/keithwegner/examplebeat/cmd"

	// Make sure all your modules and metricsets are linked in this file
	_ "github.com/keithwegner/examplebeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
