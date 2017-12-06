package cmd

import (
	"github.com/spf13/cobra"
	"strings"
)

var day05cmd = &cobra.Command{
	Use: "day05",
	Run: run05,
}

func init() {
	RootCmd.AddCommand(day05cmd)
}

func compute05a(input string) int {
	memory := ParseIntArray(input)
	count := 0
	idx := 0
	for {
		if idx < 0 || idx >= len(memory) {
			break
		}
		count++
		val := memory[idx]
		memory[idx]++
		idx += val
	}
	return count
}

func compute05b(input string) int {
	memory := ParseIntArray(input)
	count := 0
	idx := 0
	for {
		if idx < 0 || idx >= len(memory) {
			break
		}
		count++
		val := memory[idx]
		if val >= 3 {
			memory[idx]--
		} else {
			memory[idx]++
		}
		idx += val
	}
	return count
}

func run05(cmd *cobra.Command, args []string) {
	input := LoadData("data/05-input.txt")
	input = strings.Replace(input, "\n", " ", -1)
	Test(compute05a, "-1", 1)
	Test(compute05a, "1", 1)
	Test(compute05a, "0", 2)
	Test(compute05a, "0 3 0 1 -3", 5)
	PrintResult(input, compute05a(input))
	Test(compute05b, "1 1 1", 3)
	Test(compute05b, "0 3 0 1 -3", 10)
	PrintResult(input, compute05b(input))
}
