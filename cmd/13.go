package cmd

import (
	"bufio"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

var day13cmd = &cobra.Command{
	Use: "day13",
	Run: run13,
}

func init() {
	RootCmd.AddCommand(day13cmd)
}

func parse13(input string) map[int]int {
	m := make(map[int]int)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ":")
		left, _ := strconv.Atoi(parts[0])
		right, _ := strconv.Atoi(strings.Trim(parts[1], " "))
		m[left] = right
	}
	return m
}

func compute13a(input string) int {
	m := parse13(input)
	count := 0
	for key, val := range m {
		period := 2 * (val - 1)
		if (key % period) == 0 {
			count += key * val
		}
	}
	return count
}

func compute13b(input string) int {
	m := parse13(input)
	i := 0
	for {
		count := 0
		for key, val := range m {
			period := 2 * (val - 1)
			if ((key + i) % period) == 0 {
				count += 1
			}
		}
		if count == 0 {
			return i
		}
		i += 1
		if i > 10000000 {
			panic("Can't find a good starting sequence!")
			return 0
		}
	}
	return 0
}

func run13(cmd *cobra.Command, args []string) {
	input := LoadData("data/13-input.txt")
	test := LoadData("data/13-test.txt")
	Test(compute13a, test, 24)
	PrintResult(input, compute13a(input))
	Test(compute13b, test, 10)
	PrintResult(input, compute13b(input))
}
