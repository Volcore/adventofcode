package cmd

import (
	"github.com/spf13/cobra"
	"strings"
)

var day06cmd = &cobra.Command{
	Use: "day06",
	Run: run06,
}

func init() {
	RootCmd.AddCommand(day06cmd)
}

func compute06a(input string) int {
	return compute06(input, false)
}

func compute06b(input string) int {
	return compute06(input, true)
}

func compute06(input string, loopSize bool) int {
	banks := ParseIntArray(input)
	iteration := 0
	memory := make(map[string]int)
	for {
		// Check if we've seen this configuration before
		s := PrintIntArray(banks)
		if loopStart, ok := memory[s]; ok {
			// Done, we're at a configuration we've seen before
			if loopSize {
				// Second part solution
				return iteration - loopStart
			} else {
				// first part solution
				return iteration
			}
			break
		}
		memory[s] = iteration
		iteration++
		// Find largest bank
		big_idx := 0
		for idx, _ := range banks {
			if banks[idx] > banks[big_idx] {
				big_idx = idx
			}
		}
		// Redistribute
		bankCount := len(banks)
		size := banks[big_idx]
		banks[big_idx] = 0
		idx := big_idx
		for size > 0 {
			idx = (idx + 1) % bankCount
			banks[idx] += 1
			size--
		}
	}
	return 0
}

func run06(cmd *cobra.Command, args []string) {
	input := LoadData("data/06-input.txt")
	input = strings.Replace(input, "\n", " ", -1)
	input = strings.Replace(input, "\t", " ", -1)
	Test(compute06a, "0 2 7 0", 5)
	PrintResult(input, compute06a(input))
	Test(compute06b, "0 2 7 0", 4)
	PrintResult(input, compute06b(input))
}
