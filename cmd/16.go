package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

var day16cmd = &cobra.Command{
	Use: "day16",
	Run: run16,
}

func init() {
	RootCmd.AddCommand(day16cmd)
}

func spin16(input string, width int) string {
	l := len(input)
	return input[l-width:] + input[:l-width]
}

func exchange16(input string, idx1 int, idx2 int) string {
	copy := []rune(input)
	copy[idx1] = rune(input[idx2])
	copy[idx2] = rune(input[idx1])
	return string(copy)
}

func partner16(input string, r1 rune, r2 rune) string {
	idx1 := strings.Index(input, string(r1))
	idx2 := strings.Index(input, string(r2))
	copy := []rune(input)
	copy[idx1] = rune(input[idx2])
	copy[idx2] = rune(input[idx1])
	return string(copy)
}

func create16(input string) string {
	// Create the initial order
	size, _ := strconv.Atoi(input)
	order := ""
	for i := 0; i < size; i++ {
		order += string('a' + i)
	}
	return order
}

func dance16(order string, dance []string) string {
	for _, op := range dance {
		switch op[0] {
		case 's':
			{
				width, _ := strconv.Atoi(op[1:])
				order = spin16(order, width)
				break
			}
		case 'x':
			{
				indices := strings.Split(op[1:], "/")
				idx1, _ := strconv.Atoi(indices[0])
				idx2, _ := strconv.Atoi(indices[1])
				order = exchange16(order, idx1, idx2)
				break
			}
		case 'p':
			{
				order = partner16(order, rune(op[1]), rune(op[3]))
				break
			}
		}
	}
	return order
}

func compute16a(input string) string {
	parts := strings.Split(input, ",")
	order := create16(parts[0])
	return dance16(order, parts[1:])
}

func compute16b(input string) string {
	parts := strings.Split(input, ",")
	initial := create16(parts[0])
	order := initial
	memory := make(map[string]int)
	for i := 0; i < 1000000000; i++ {
		if val, ok := memory[order]; ok {
			// Found a loop
			fmt.Println("Found a loop for ", order, "original index ", val, "and now again at start of ", i)
			loopSize := i - val
			reps := (1000000000 - val) / loopSize
			finalOccurence := loopSize*reps + val
			fmt.Println("loop size", loopSize, "reps", reps, "final", finalOccurence)
			for j := finalOccurence; j < 1000000000; j++ {
				order = dance16(order, parts[1:])
			}
			return order
		} else {
			memory[order] = i
		}
		order = dance16(order, parts[1:])
	}
	return order
}

func run16(cmd *cobra.Command, args []string) {
	input := LoadData("data/16-input.txt")
	input = "16," + input
	TestS(compute16a, "5,s1,x3/4,pe/b", "baedc")
	PrintResultS(input, compute16a(input))
	TestS(compute16b, "5,s1,x3/4,pe/b", "abcde")
	PrintResultS(input, compute16b(input))
}
