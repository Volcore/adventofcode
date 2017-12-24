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
	patterns [][]bool
	output   [][]bool
}

func parseTile21(input string) []bool {
	rows := strings.Split(input, "/")
	tile := make([]bool, 0)
	for _, row := range rows {
		for _, c := range row {
			tile = append(tile, c == '#')
		}
	}
	return tile
}

func print21(pattern []bool) {
	for _, v := range pattern {
		if v {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
	}
	fmt.Print(" ")
}

func printTiles21(tiles [][]bool) {
	for _, tile := range tiles {
		print21(tile)
	}
	fmt.Println("")
}

func rot21(t []bool) []bool {
	if len(t) == 4 {
		// 0 1  => 1 3
		// 2 3     0 2
		return []bool{t[1], t[3], t[0], t[2]}
	} else {
		// 0 1 2 => 2 5 8
		// 3 4 5    1 4 7
		// 6 7 8    0 3 6
		return []bool{t[2], t[5], t[8], t[1], t[4], t[7], t[0], t[3], t[6]}
	}
}

func flip21x(t []bool) []bool {
	if len(t) == 4 {
		// 0 1  => 1 0
		// 2 3     3 2
		return []bool{t[1], t[0], t[3], t[2]}
	} else {
		// 0 1 2 => 2 1 0
		// 3 4 5    5 4 3
		// 6 7 8    8 7 6
		return []bool{t[2], t[1], t[0], t[5], t[4], t[3], t[8], t[7], t[6]}
	}
}

func flip21y(t []bool) []bool {
	if len(t) == 4 {
		// 0 1  => 2 3
		// 2 3     0 1
		return []bool{t[2], t[3], t[0], t[1]}
	} else {
		// 0 1 2 => 6 7 8
		// 3 4 5    3 4 5
		// 6 7 8    0 1 2
		return []bool{t[6], t[7], t[8], t[3], t[4], t[5], t[0], t[1], t[2]}
	}
}

func mutate21(tile []bool) [][]bool {
	return [][]bool{
		tile,
		rot21(tile),
		rot21(rot21(tile)),
		rot21(rot21(rot21(tile))),
		flip21x(rot21(rot21(rot21(tile)))),
		rot21(flip21x(rot21(rot21(rot21(tile))))),
		rot21(rot21(flip21x(rot21(rot21(rot21(tile)))))),
		rot21(rot21(rot21(flip21x(rot21(rot21(rot21(tile))))))),
	}
}

func split21(tile []bool) [][]bool {
	if len(tile) == 16 {
		// 0 1 2 3
		// 4 5 6 7
		// 8 9 10 11
		// 12 13 14 15
		// create 2x2 grid of 2x2 tiles
		return [][]bool{
			[]bool{tile[0], tile[1], tile[4], tile[5]},
			[]bool{tile[2], tile[3], tile[6], tile[7]},
			[]bool{tile[8], tile[9], tile[12], tile[13]},
			[]bool{tile[10], tile[11], tile[14], tile[15]},
		}
	} else {
		// return single 3x3 grid
		return [][]bool{tile}
	}
}

func parse21(input string) []*pattern21 {
	patterns := make([]*pattern21, 0)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " => ")
		p := new(pattern21)
		p.patterns = mutate21(parseTile21(parts[0]))
		p.output = split21(parseTile21(parts[1]))
		patterns = append(patterns, p)
		//fmt.Println(line)
		//printTiles21(p.patterns)
		//printTiles21(p.output)
	}
	return patterns
}

// ##./#.#/#.. => #.../###./#.##/#.##

func step21(grid [][]bool, patterns []*pattern21) [][]bool {
	newgrid := [][]bool{}
	for _, tile := range grid {
		// find match in patterns
		found := false
		for _, pattern := range patterns {
			for _, cmp := range pattern.patterns {
				if TestBoolSliceEqual(cmp, tile) {
					newgrid = append(newgrid, pattern.output...)
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		if !found {
			panic("Pattern without match!")
		}
	}
	return newgrid
}

func printGrid21(grid [][]bool) {
	fmt.Println("Grid")
	for _, tile := range grid {
		print21(tile)
	}
	fmt.Println("")
}

func compute21a(input string, depth int) int {
	patterns := parse21(input)
	grid := [][]bool{[]bool{false, true, false, false, false, true, true, true, true}}
	// Simulate
	for i := 0; i < depth; i++ {
		// Shit i misunderstood the specs... GG
		printGrid21(grid)
		grid = step21(grid, patterns)
	}
	// Count result
	count := 0
	for _, tile := range grid {
		for _, cell := range tile {
			if cell {
				count++
			}
		}
	}
	return count
}

func compute21a2(input string) int {
	return compute21a(input, 2)
}

func compute21a5(input string) int {
	return compute21a(input, 5)
}

func compute21a18(input string) int {
	return compute21a(input, 18)
}

func run21(cmd *cobra.Command, args []string) {
	input := LoadDataRaw("data/21-input.txt")
	test := LoadDataRaw("data/21-test.txt")
	Test(compute21a2, test, 12)
	PrintResult(input, compute21a5(input))
	//PrintResult(input, compute21a18(input))
}
