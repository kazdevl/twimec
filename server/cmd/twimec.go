package main

import (
	"github.com/kazdevl/twimec/cmd/set"
	"github.com/kazdevl/twimec/cmd/start"
	"github.com/kazdevl/twimec/cmd/stop"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "twimec"}
	rootCmd.AddCommand(
		start.NewCmd(),
		set.NewCmd(),
		stop.NewCmd(),
	)
	cobra.CheckErr(rootCmd.Execute())
}
