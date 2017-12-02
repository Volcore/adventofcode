package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"strings"
)

var day01cmd = &cobra.Command{
	Use: "day01",
	Run: run01,
}

func init() {
	RootCmd.AddCommand(day01cmd)
}

func compute01(input string) int {
	fmt.Println("Input is:\n", input)
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

func test01(input string, output int) {
	val := compute01(input)
	if val != output {
		fmt.Println("Test failed, value should be", output, "but is", val)
	}
}

func run01(cmd *cobra.Command, args []string) {
	test01("1122", 3)
	test01("1111", 4)
	test01("1234", 0)
	test01("91212129", 9)
	b, err := ioutil.ReadFile("data/01-input.txt")
	if err == nil {
		input := strings.Trim(string(b), "\n\r \t")
		compute01(input)
	}
}
