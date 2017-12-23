package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var day21cmd = &cobra.Command{
	Use: "day21",
	Run: run21,
}

func init() {
	RootCmd.AddCommand(day21cmd)
}

type pattern21 struct {
	size   int
	tile   uint16
	output uint16
}

func parseTile21(input string) (uint16, int) {
	tile := uint16(0)
	rows := strings.Split(input, "/")
	width := len(rows[0])
	for ridx, row := range rows {
		for cidx, c := range row {
			if c == '#' {
				tile |= (1 << uint(width*ridx+cidx))
			}
		}
	}
	return tile, width
}

func parse21(input string) []*pattern21 {
	patterns := make([]*pattern21, 0)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " => ")
		p := new(pattern21)
		p.tile, p.size = parseTile21(parts[0])
		// TODO(VS): might be an array of uints for four 2x2 tile
		p.output, _ = parseTile21(parts[1])
		patterns = append(patterns, p)
	}
	return patterns
}

func step21(grid []uint16, patterns []*pattern21) []uint16 {
	return grid
}

func compute21a(input string, depth int) int {
	patterns := parse21(input)
	grid := []uint16{2 | 32 | 64 | 128 | 256}
	fmt.Println(patterns, grid)
	fmt.Printf("%b\n", grid)
	// Simulate
	for i := 0; i < depth; i++ {
		grid = step21(grid, patterns)
	}
	// Count result
	return 0
	//count := 0
	//for _, row := range grid {
	//	for _, cell := range row {
	//		if cell {
	//			count++
	//		}
	//	}
	//}
	//return count
}

func compute21a2(input string) int {
	return compute21a(input, 2)
}

func compute21a5(input string) int {
	return compute21a(input, 5)
}

func run21(cmd *cobra.Command, args []string) {
	input := LoadDataRaw("data/21-input.txt")
	test := LoadDataRaw("data/21-test.txt")
	Test(compute21a2, test, 12)
	PrintResult(input, compute21a5(input))
}
