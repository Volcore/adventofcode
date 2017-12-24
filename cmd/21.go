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
	fmt.Println("")
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
		return [][]bool{
			tile[0:4],
			tile[4:8],
			tile[8:12],
			tile[12:16],
		}
	} else {
		return [][]bool{
			tile[0:3],
			tile[3:6],
			tile[6:9],
		}
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
	}
	return patterns
}

func step21(grid [][]bool, patterns []*pattern21) [][]bool {
	size := len(grid)
	tileSize := 2
	if size%2 == 0 {
		tileSize = 2
	} else {
		tileSize = 3
	}
	newgrid := [][]bool{}
	for y := 0; y < size/tileSize; y++ {
		rows := [][]bool{
			[]bool{},
			[]bool{},
			[]bool{},
			[]bool{},
		}
		for x := 0; x < size/tileSize; x++ {
			// build the tile
			tile := make([]bool, tileSize*tileSize)
			for j := 0; j < tileSize; j++ {
				for i := 0; i < tileSize; i++ {
					tile[i+j*tileSize] = grid[y*tileSize+j][x*tileSize+i]
				}
			}
			// Match the tile
			found := false
			for _, pattern := range patterns {
				for _, cmp := range pattern.patterns {
					if TestBoolSliceEqual(cmp, tile) {
						for idx, output := range pattern.output {
							rows[idx] = append(rows[idx], output...)
						}
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
		for _, row := range rows {
			if len(row) == 0 {
				break
			}
			newgrid = append(newgrid, row)
		}
	}
	return newgrid
}

func printGrid21(grid [][]bool) {
	fmt.Println("*** Grid")
	for _, tile := range grid {
		print21(tile)
	}
	fmt.Println("")
}

func compute21a(input string, depth int) int {
	patterns := parse21(input)
	grid := [][]bool{
		[]bool{false, true, false},
		[]bool{false, false, true},
		[]bool{true, true, true},
	}
	// Simulate
	for i := 0; i < depth; i++ {
		fmt.Println("Depth", i, "/", depth)
		//printGrid21(grid)
		grid = step21(grid, patterns)
	}
	//printGrid21(grid)
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
	PrintResult(input, compute21a18(input))
}
