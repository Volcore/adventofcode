package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"strings"
)

var day01cmd = &cobra.Command{
	Use: "day01",
	Run: run,
}

func init() {
	RootCmd.AddCommand(day01cmd)
}

func compute(input string) int {
	fmt.Println("Input is: ", input)
	last := input[len(input)-1]
	count := 0
	for i := 0; i < len(input); i++ {
		val := input[i]
		if last == val {
			count += int(val - '0')
		}
		last = val
	}
	fmt.Println("Count is: ", count)
	return count
}

func test(input string, output int) {
	val := compute(input)
	if val != output {
		fmt.Println("Test failed, value should be", output, "but is", val)
	}
}

func run(cmd *cobra.Command, args []string) {
	test("1122", 3)
	test("1111", 4)
	test("1234", 0)
	test("91212129", 9)
	b, err := ioutil.ReadFile("data/01-input.txt")
	if err == nil {
		input := strings.Trim(string(b), "\n\r \t")
		compute(input)
	}

}
