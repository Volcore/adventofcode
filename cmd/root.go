package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var cfgFile string

var RootCmd = &cobra.Command{
	Use:   "adventofcode",
	Short: "Advent of code entries.",
	Long:  "Run the different entries, dayXY. eg adventofcode day01",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
