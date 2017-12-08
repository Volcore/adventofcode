package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

var day07cmd = &cobra.Command{
	Use: "day07",
	Run: run07,
}

func init() {
	RootCmd.AddCommand(day07cmd)
}

type node07 struct {
	weight   int
	name     string
	parent   string
	children []string
}

func parseTree07(input string) map[string]*node07 {
	m := make(map[string]*node07)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		str := strings.Replace(scanner.Text(), ",", "", -1)
		parts := strings.Split(str, " ")
		name := parts[0]
		if _, ok := m[name]; !ok {
			m[name] = new(node07)
		}
		node := m[name]
		node.name = name
		node.weight, _ = strconv.Atoi(parts[1][1 : len(parts[1])-1])
		num_children := len(parts) - 3
		if num_children <= 0 {
			continue
		}
		node.children = make([]string, num_children)
		for i := 0; i < num_children; i++ {
			child := parts[i+3]
			node.children[i] = child
			if _, ok := m[child]; !ok {
				m[child] = new(node07)
			}
			m[child].parent = name
		}
	}
	return m
}

func findRoot07a(tree map[string]*node07) string {
	// Find a node without a parent
	result := ""
	for name := range tree {
		node := tree[name]
		if node.parent == "" {
			if result != "" {
				fmt.Println("More than one result:", result, name)
			}
			result = name
		}
	}
	return result
}

func compute07a(input string) string {
	m := parseTree07(input)
	return findRoot07a(m)
}

func recurse07b(tree map[string]*node07, current string) int {
	node := tree[current]
	if node.children != nil {
		subWeights := make([]int, len(node.children))
		subWeightCounts := make(map[int]int)
		subWeightIndex := make(map[int]int)
		sum := node.weight
		for idx := range node.children {
			child := node.children[idx]
			subWeights[idx] = recurse07b(tree, child)
			if subWeights[idx] < 0 {
				// Pass through if we've found the issue already
				return subWeights[idx]
			}
			subWeightCounts[subWeights[idx]]++
			subWeightIndex[subWeights[idx]] = idx
			sum += subWeights[idx]
		}
		// Check if all weights are equal
		if len(subWeightCounts) != 1 {
			if len(subWeightCounts) == 0 || len(subWeightCounts) > 2 || len(node.children) <= 2 {
				panic("Unsolvable situation, too few children or too much disagreement!")
			}
			// Find consensus
			for subweight := range subWeightCounts {
				if subWeightCounts[subweight] == 1 {
					// Found the culprit, and find an okay one
					idx := subWeightIndex[subweight]
					okidx := (idx + 1) % len(node.children)
					// Compute the delta
					delta := subweight - subWeights[okidx]
					// Compute proper weight of offender
					offender := node.children[idx]
					offenderNode := tree[offender]
					return -(offenderNode.weight - delta)
				}
			}
			// Uh-oh, we found noone
			panic("Didn't find any singular subweights that are out of line!")
		}
		return sum
	} else {
		return node.weight
	}
}

func compute07b(input string) int {
	m := parseTree07(input)
	root := findRoot07a(m)
	return recurse07b(m, root)
}

func run07(cmd *cobra.Command, args []string) {
	test := LoadData("data/07-test.txt")
	input := LoadData("data/07-input.txt")
	TestS(compute07a, test, "tknk")
	PrintResultS(input, compute07a(input))
	Test(compute07b, test, -60)
	PrintResult(input, compute07b(input))
}
