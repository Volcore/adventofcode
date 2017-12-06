package cmd

import (
	"github.com/spf13/cobra"
)

var day01cmd = &cobra.Command{
	Use: "day01",
	Run: run01,
}

func init() {
	RootCmd.AddCommand(day01cmd)
}

func compute01(input string) int {
	last := input[len(input)-1]
	count := 0
	for i := 0; i < len(input); i++ {
		val := input[i]
		if last == val {
			count += int(val - '0')
		}
		last = val
	}
	return count
}

func compute01b(input string) int {
	l := len(input)
	count := 0
	for i := 0; i < len(input); i++ {
		val1 := input[i]
		val2 := input[(i+l/2)%l]
		if val1 == val2 {
			count += int(val1 - '0')
		}
	}
	return count
}

func run01(cmd *cobra.Command, args []string) {
	input := LoadData("data/01-input.txt")
	Test(compute01, "1122", 3)
	Test(compute01, "1111", 4)
	Test(compute01, "1234", 0)
	Test(compute01, "91212129", 9)
	PrintResult(input, compute01(input))
	Test(compute01b, "1212", 6)
	Test(compute01b, "1221", 0)
	Test(compute01b, "123425", 4)
	Test(compute01b, "123123", 12)
	Test(compute01b, "12131415", 4)
	PrintResult(input, compute01b(input))
}
