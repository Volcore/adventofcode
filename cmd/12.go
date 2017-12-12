package cmd

import (
	"bufio"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

var day12cmd = &cobra.Command{
	Use: "day12",
	Run: run12,
}

func init() {
	RootCmd.AddCommand(day12cmd)
}

type node12 struct {
	self   int
	refs   []int
	tagged bool
}

func parse12(input string) map[int]*node12 {
	input = strings.Replace(input, ",", "", -1)
	scanner := bufio.NewScanner(strings.NewReader(input))
	m := make(map[int]*node12)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		left, _ := strconv.Atoi(parts[0])
		right := make([]int, 0)
		for _, rhs := range parts[2:] {
			val, _ := strconv.Atoi(rhs)
			right = append(right, val)
		}
		node := new(node12)
		node.self = left
		node.refs = right
		m[left] = node
	}
	return m
}

func tag12a(m map[int]*node12, idx int) {
	if m[idx].tagged {
		return
	}
	m[idx].tagged = true
	for _, child := range m[idx].refs {
		tag12a(m, child)
	}
}

func compute12a(input string) int {
	m := parse12(input)
	tag12a(m, 0)
	count := 0
	for _, node := range m {
		if node.tagged {
			count += 1
		}
	}
	return count
}

func compute12b(input string) int {
	m := parse12(input)
	count := 0
	for {
		// Find untagged subgraph
		untaggedIdx := -1
		for idx, node := range m {
			if node.tagged {
				continue
			}
			untaggedIdx = idx
			break
		}
		// Are we done?
		if untaggedIdx == -1 {
			break
		}
		// Tag all subgraph
		tag12a(m, untaggedIdx)
		count += 1
	}
	return count
}

func run12(cmd *cobra.Command, args []string) {
	input := LoadData("data/12-input.txt")
	test := LoadData("data/12-test.txt")
	Test(compute12a, test, 6)
	PrintResult(input, compute12a(input))
	Test(compute12b, test, 2)
	PrintResult(input, compute12b(input))
}
