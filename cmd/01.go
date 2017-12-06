package cmd

import (
	"fmt"
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
	PrintResult(input, count)
	return count
}

func test01(input string, output int) {
	val := compute01(input)
	if val != output {
		fmt.Println("Test failed, value should be", output, "but is", val)
	}
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
	PrintResult(input, count)
	return count
}

func test01b(input string, output int) {
	val := compute01b(input)
	if val != output {
		fmt.Println("Test failed, value should be", output, "but is", val)
	}
}

func run01(cmd *cobra.Command, args []string) {
	input := LoadData("data/01-input.txt")
	test01("1122", 3)
	test01("1111", 4)
	test01("1234", 0)
	test01("91212129", 9)
	compute01(input)
	test01b("1212", 6)
	test01b("1221", 0)
	test01b("123425", 4)
	test01b("123123", 12)
	test01b("12131415", 4)
	compute01b(input)
}
