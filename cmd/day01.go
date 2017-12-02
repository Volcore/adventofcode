package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var day01class = &cobra.Command{
	Use: "day01",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("asdf")
	},
}
