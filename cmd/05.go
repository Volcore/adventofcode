package cmd

import (
	"fmt"
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
	PrintResult(input, count)
	return count
}

func test05a(input string, output int) {
	val := compute05a(input)
	if val != output {
		fmt.Println("Test failed, value should be", output, "but is", val)
	}
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
	PrintResult(input, count)
	return count
}

func test05b(input string, output int) {
	val := compute05b(input)
	if val != output {
		fmt.Println("Test failed, value should be", output, "but is", val)
	}
}

func run05(cmd *cobra.Command, args []string) {
	input := LoadData("data/05-input.txt")
	input = strings.Replace(input, "\n", " ", -1)
	test05a("-1", 1)
	test05a("1", 1)
	test05a("0", 2)
	test05a("0 3 0 1 -3", 5)
	compute05a(input)
	test05b("1 1 1", 3)
	test05b("0 3 0 1 -3", 10)
	compute05b(input)
}
